package service

import (
	"context"
	"strconv"
	"time"

	"github.com/larek-tech/innohack/backend/internal/chat/model"
)

func (s *Service) ProcessMessage(ctx context.Context, req model.QueryDto, out chan<- model.ResponseDto, cancel <-chan int64) {
	ticker := time.NewTicker(time.Second)
	cnt := int64(0)
	desc := ""

	for {
		select {
		case sessionID := <-cancel:
			s.log.Info().Int64("session id", sessionID).Msg("canceled")
			return
		case <-ticker.C:
			// TODO: grpc stream
			desc += strconv.FormatInt(cnt, 10)
			out <- model.ResponseDto{Description: desc}
			cnt++
			if cnt > 10 {
				close(out)
				return
			}
		}
	}
}
