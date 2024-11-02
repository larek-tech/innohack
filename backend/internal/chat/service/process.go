package service

import (
	"context"
	"time"

	"github.com/larek-tech/innohack/backend/internal/analytics/pb"
	"github.com/larek-tech/innohack/backend/internal/chat/model"
	"github.com/larek-tech/innohack/backend/pkg"
)

func (s *Service) ProcessMessage(ctx context.Context, req model.QueryDto, out chan<- model.ResponseDto, cancel <-chan int64) {
	charts, err := s.analytics.GetCharts(ctx, &pb.Params{
		QueryId:   req.ID,
		StartDate: time.Now().AddDate(-2, 0, 0).String(),
		EndDate:   time.Now().AddDate(-1, 0, 0).String(),
		Prompt:    "",
	})
	if err != nil {
		s.log.Err(pkg.WrapErr(err)).Msg("get charts")
		close(out)
		return
	}

	chartsResponse := model.ResponseDto{
		QueryID:     req.ID,
		Source:      charts.GetSource(),
		Filename:    charts.GetFilename(),
		Charts:      []model.Chart{},
		Description: "",
		Multipliers: []model.Multiplier{},
		CreatedAt:   time.Time{},
	}

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
