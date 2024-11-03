package controller

import (
	"context"
	"io"
	"strings"

	"github.com/larek-tech/innohack/backend/internal/analytics/pb"
	"github.com/larek-tech/innohack/backend/internal/chat/model"
	"github.com/larek-tech/innohack/backend/pkg"
	"github.com/rs/zerolog/log"
)

func (ctrl *Controller) GetDescription(ctx context.Context, req model.QueryDto, out chan<- model.ResponseDto, cancel <-chan int64) {
	stream, err := ctrl.analytics.GetDescriptionStream(ctx, &pb.Params{
		QueryId: req.ID,
		Prompt:  req.Prompt,
	})
	if err != nil {
		close(out)
		log.Err(pkg.WrapErr(err)).Msg("receive stream description")
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
			log.Debug().Int64("query id", queryID).Msg("canceled")
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
				log.Err(pkg.WrapErr(err)).Msg("reading stream")
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
