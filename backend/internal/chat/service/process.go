package service

import (
	"context"
	"io"
	"strings"
	"time"

	"github.com/larek-tech/innohack/backend/internal/analytics/pb"
	"github.com/larek-tech/innohack/backend/internal/chat/model"
	"github.com/larek-tech/innohack/backend/pkg"
)

func (s *Service) ProcessMessage(ctx context.Context, req model.QueryDto, out chan<- model.ResponseDto, cancel <-chan int64) {
	s.getCharts(ctx, req, out)
	s.getDescription(ctx, req, out, cancel)

	// ticker := time.NewTicker(time.Second)
	// cnt := int64(0)

	// for {
	// 	select {
	// 	case sessionID := <-cancel:
	// 		s.log.Info().Int64("session id", sessionID).Msg("canceled")
	// 		return
	// 	case <-ticker.C:
	// 		// TODO: grpc stream
	// 		out <- model.ResponseDto{Description: strconv.FormatInt(cnt, 64)}
	// 		cnt++
	// 		if cnt > 10 {
	// 			close(out)
	// 			return
	// 		}
	// 	}
	// }
}

func (s *Service) getCharts(ctx context.Context, req model.QueryDto, out chan<- model.ResponseDto) {
	// TODO: поменять на выбор дат
	startDate := time.Now().AddDate(-2, 0, 0).String()
	endDate := time.Now().AddDate(-1, 0, 0).String()

	charts, err := s.analytics.GetCharts(ctx, &pb.Params{
		QueryId:   req.ID,
		StartDate: startDate,
		EndDate:   endDate,
		Prompt:    "",
	})
	if err != nil {
		s.log.Err(pkg.WrapErr(err)).Msg("get charts")
		close(out)
		return
	}

	chartsResponse := model.ResponseDto{
		QueryID:     req.ID,
		Sources:     charts.GetSources(),
		Filenames:   charts.GetFilenames(),
		Charts:      model.ChartsFromPb(charts.GetCharts()),
		Description: charts.Description,
		Multipliers: model.MultipliersFromPb(charts.GetMultipliers()),
	}

	out <- chartsResponse
}

func (s *Service) getDescription(ctx context.Context, req model.QueryDto, out chan<- model.ResponseDto, cancel <-chan int64) {
	// TODO: поменять на выбор дат
	startDate := time.Now().AddDate(-2, 0, 0).String()
	endDate := time.Now().AddDate(-1, 0, 0).String()

	stream, err := s.analytics.GetDescriptionStream(ctx, &pb.Params{
		QueryId:   req.ID,
		StartDate: startDate,
		EndDate:   endDate,
		Prompt:    req.Prompt,
	})
	if err != nil {
		close(out)
		s.log.Err(pkg.WrapErr(err)).Msg("receive stream description")
		return
	}

	var (
		sources     = make([]string, 0)
		filenames   = make([]string, 0)
		charts      = make([]model.Chart, 0)
		multipliers = make([]model.Multiplier, 0)
		buff        = strings.Builder{}
		idx         uint
	)

	for {
		select {
		case queryID := <-cancel:
			s.log.Debug().Int64("query id", queryID).Msg("canceled")
			return
		default:
			resp, err := stream.Recv()
			if err == io.EOF {
				out <- model.ResponseDto{
					QueryID:     req.ID,
					Sources:     sources,
					Filenames:   filenames,
					Charts:      charts,
					Description: buff.String(),
					Multipliers: multipliers,
					IsLast:      true,
				}
				return
			}
			if err != nil {
				s.log.Err(pkg.WrapErr(err)).Msg("reading stream")
				close(out)
				return
			}

			curSources := resp.GetSources()
			sources = append(sources, curSources...)

			curFilenames := resp.GetFilenames()
			filenames = append(filenames, curFilenames...)

			curDescription := resp.GetDescription()

			if idx == 0 {
				curCharts := model.ChartsFromPb(resp.GetCharts())
				charts = append(charts, curCharts...)

				curMultipliers := model.MultipliersFromPb(resp.GetMultipliers())
				multipliers = append(multipliers, curMultipliers...)

				out <- model.ResponseDto{
					QueryID:     req.ID,
					Sources:     curSources,
					Filenames:   curFilenames,
					Charts:      curCharts,
					Description: curDescription,
					Multipliers: curMultipliers,
				}
			} else {
				out <- model.ResponseDto{
					QueryID:     req.ID,
					Sources:     curSources,
					Filenames:   curFilenames,
					Description: curDescription,
				}
			}

			buff.WriteString(curDescription)
		}
	}
}
