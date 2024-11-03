package handler

import (
	"context"
	"strconv"
	"strings"

	"github.com/gofiber/websocket/v2"
	"github.com/larek-tech/innohack/backend/internal/chat/model"
	"github.com/larek-tech/innohack/backend/pkg/jwt"
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

	h.log.Err(err).Msg("chat error")

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

	ctx := context.Background()

	// первое сообщение содержит access token
	authQuery := model.QueryDto{}
	if err := c.ReadJSON(&authQuery); err != nil {
		h.respondError(c, err)
		return
	}
	if _, err := jwt.VerifyAccessToken(authQuery.Prompt, h.jwtSecret); err != nil {
		h.respondError(c, err)
		return
	}

	sessionID, err := strconv.ParseInt(c.Params("session_id"), 10, 64)
	if err != nil {
		h.respondError(c, err)
		return
	}

	var (
		resp   model.ResponseDto
		req    model.QueryDto
		desc   = strings.Builder{}
		cancel = make(chan int64, 1)
		out    = make(chan model.ResponseDto)
	)

	for {
		if err := c.ReadJSON(&req); err != nil {
			h.respondError(c, err)
			return
		}

		queryID, err := h.service.InsertQuery(ctx, sessionID, req)
		if err != nil {
			h.respondError(c, err)
			return
		}
		req.ID = queryID

		go h.service.GetDescription(ctx, req, out, cancel)

	chunks:
		for {
			select {
			case chunk, ok := <-out:
				// если закончили читать
				if !ok {
					h.log.Error().Msg("error while processing")
					return
				}

				if chunk.IsLast {
					resp = chunk
					h.log.Debug().Int64("query id", queryID).Msg("finished processing")
					break chunks
				}

				if chunk.Charts != nil {
					// если первое сообщение с графиками, которое должно быть цельным
					copy(resp.Charts, chunk.Charts)
				} else {
					// если остальные сообщения с описанием, которые идут по токенам
					desc.WriteString(chunk.Description)
				}

				if err := c.WriteJSON(chunk); err != nil {
					h.respondError(c, err)
					continue
				}
			}
		}

		if err := h.service.InsertResponse(ctx, sessionID, resp); err != nil {
			h.respondError(c, err)
			return
		}
	}
}
