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

func (s *Service) GetCharts(ctx context.Context, req model.QueryDto, out chan<- model.ResponseDto) {
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
		Charts:      model.ChartsFromPb(charts.GetCharts()),
		Multipliers: model.MultipliersFromPb(charts.GetMultipliers()),
	}

	out <- chartsResponse
}

func (s *Service) GetDescription(ctx context.Context, req model.QueryDto, out chan<- model.ResponseDto, cancel <-chan int64) {
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
		sources   = make([]string, 0)
		filenames = make([]string, 0)
		buff      = strings.Builder{}
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
					Description: buff.String(),
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

			out <- model.ResponseDto{
				QueryID:     req.ID,
				Sources:     curSources,
				Filenames:   curFilenames,
				Description: curDescription,
			}

			buff.WriteString(curDescription)
		}
	}
}
