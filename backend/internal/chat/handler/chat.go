package handler

import (
	"github.com/gofiber/websocket/v2"
	"github.com/larek-tech/innohack/backend/internal/chat/model"
)

func (h *Handler) closeHandler(code int, text string) error {
	h.log.Info().Int("code", code).Str("text", text).Msg("close handler")
	return nil
}

func (h *Handler) respondError(c *websocket.Conn, err error) {
	resp := model.ResponseDto{
		Err:    err,
		IsLast: true,
	}

	if err := c.WriteJSON(resp); err != nil {
		h.log.Warn().Err(err).Msg("failed to respond with error")
		return
	}
}

func (h *Handler) ProcessConn(c *websocket.Conn) {
	h.log.Info().Str("addr", c.LocalAddr().String()).Msg("new conn")
	c.SetCloseHandler(h.closeHandler)

	defer func() {
		if err := c.Close(); err != nil {
			h.log.Warn().Err(err).Msg("failed to close websocket conn")
			return
		}
		h.log.Info().Msg("conn closed")
	}()

	// ctx := context.Background()

	// sessionID, err := h.service.InsertSession(ctx, authCookie)
	// if err != nil {
	// 	log.Err(err).Msg("chat error")
	// 	h.respondError(c, ctx, err)
	// 	return
	// }

	// var (
	// 	resp   model.ResponseDto
	// 	req    model.QueryDto
	// 	desc   = strings.Builder{}
	// 	cancel = make(chan int64, 1)
	// 	out    = make(chan model.ResponseDto)
	// )

	// for {
	// 	if err := c.ReadJSON(&req); err != nil {
	// 		log.Err(err).Msg("chat error")
	// 		h.respondError(c, ctx, err)
	// 		return
	// 	}

	// 	queryID, err := h.service.InsertQuery(ctx, sessionID, req)
	// 	if err != nil {
	// 		log.Err(err).Msg("chat error")
	// 		h.respondError(c, ctx, err)
	// 		return
	// 	}
	// 	req.ID = queryID

	// 	go h.service.GetDescription(ctx, req, out, cancel)

	// chunks:
	// 	for {
	// 		select {
	// 		case chunk, ok := <-out:
	// 			// если закончили читать
	// 			if !ok {
	// 				h.log.Error().Msg("error while processing")
	// 				return
	// 			}

	// 			if chunk.IsLast {
	// 				resp = chunk
	// 				h.log.Debug().Int64("query id", queryID).Msg("finished processing")
	// 				break chunks
	// 			}

	// 			if chunk.Charts != nil {
	// 				// если первое сообщение с графиками, которое должно быть цельным
	// 				copy(resp.Charts, chunk.Charts)
	// 			} else {
	// 				// если остальные сообщения с описанием, которые идут по токенам
	// 				desc.WriteString(chunk.Description)
	// 			}

	// 			msg, err := h.newHTMLMessage(ctx, chunk, nil)
	// 			if err != nil {
	// 				log.Err(err).Msg("chat error")
	// 				h.respondError(c, ctx, err)
	// 				continue
	// 			}

	// 			if err := c.WriteMessage(websocket.TextMessage, msg); err != nil {
	// 				h.log.Err(err).Msg("write message chunk")
	// 				h.respondError(c, ctx, err)
	// 				continue
	// 			}
	// 		}
	// 	}

	// 	if err := h.service.InsertResponse(ctx, sessionID, resp); err != nil {
	// 		h.log.Err(err).Msg("save response")
	// 		h.respondError(c, ctx, err)
	// 		return
	// 	}
	// }
}
