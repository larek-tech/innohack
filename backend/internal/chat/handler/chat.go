package handler

import (
	"context"
	"strconv"
	"strings"

	"github.com/gofiber/websocket/v2"
	"github.com/google/uuid"
	"github.com/larek-tech/innohack/backend/internal/chat/model"
	"github.com/larek-tech/innohack/backend/pkg/jwt"
	"github.com/rs/zerolog/log"
)

func (h *Handler) closeHandler(code int, text string) error {
	log.Info().Int("code", code).Str("text", text).Msg("close handler")
	return nil
}

func (h *Handler) respondError(c *websocket.Conn, err error) {
	resp := model.ResponseDto{
		Err:    err,
		IsLast: true,
	}

	log.Err(err).Msg("chat error")

	if err := c.WriteJSON(resp); err != nil {
		log.Warn().Err(err).Msg("failed to respond with error")
		return
	}
}

func (h *Handler) ProcessConn(c *websocket.Conn) {
	ctx := context.Background()

	ctx, span := h.tracer.Start(ctx, "chat.handler.process_conn")
	defer span.End()

	log.Info().Str("addr", c.LocalAddr().String()).Msg("new conn")
	c.SetCloseHandler(h.closeHandler)

	defer func() {
		if err := c.Close(); err != nil {
			log.Warn().Err(err).Msg("failed to close websocket conn")
			return
		}
		log.Info().Msg("conn closed")
	}()

	// первое сообщение содержит access token
	authQuery := model.QueryDto{}
	if err := c.ReadJSON(&authQuery); err != nil {
		h.respondError(c, err)
		return
	}

	token, err := jwt.VerifyAccessToken(authQuery.Prompt, h.jwtSecret)
	if err != nil {
		h.respondError(c, err)
		return
	}

	subject, err := token.Claims.GetSubject()
	if err != nil {
		h.respondError(c, err)
		return
	}
	userID, err := strconv.ParseInt(subject, 10, 64)
	if err != nil {
		h.respondError(c, err)
		return
	}

	sessionID, err := uuid.Parse(c.Params("session_id"))
	if err != nil {
		h.respondError(c, err)
		return
	}

	defer func() {
		if err := h.sc.Cleanup(ctx, sessionID, userID); err != nil {
			log.Warn().Err(err).Msg("cleanup session")
		}
	}()

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

		queryID, err := h.cc.InsertQuery(ctx, sessionID, req)
		if err != nil {
			h.respondError(c, err)
			return
		}
		req.ID = queryID

		go h.cc.GetDescription(ctx, req, out, cancel)

	chunks:
		for {
			select {
			case chunk, ok := <-out:
				if !ok {
					log.Error().Msg("error while processing")
					return
				}

				desc.WriteString(chunk.Description)

				if err := c.WriteJSON(chunk); err != nil {
					h.respondError(c, err)
					continue
				}

				if chunk.IsLast {
					resp = chunk
					log.Debug().Int64("query id", queryID).Msg("finished processing")
					break chunks
				}
			}
		}

		if err := h.cc.InsertResponse(ctx, sessionID, resp); err != nil {
			h.respondError(c, err)
			return
		}
	}
}
