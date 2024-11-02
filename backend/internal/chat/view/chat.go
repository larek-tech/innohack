package view

import (
	"bufio"
	"bytes"
	"context"
	"strings"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/websocket/v2"
	"github.com/larek-tech/innohack/backend/internal/chat/model"
	"github.com/larek-tech/innohack/backend/internal/shared"
	"github.com/rs/zerolog/log"
)

const (
	unauthorizedCookie = "unauthorized"
)

func (v *View) ChatPage(c *fiber.Ctx) error {
	return adaptor.HTTPHandler(
		templ.Handler(ChatPage()),
	)(c)
}

func (v *View) closeHandler(code int, text string) error {
	v.log.Info().Int("code", code).Str("text", text).Msg("close handler")
	return nil
}

func (v *View) respondError(c *websocket.Conn, err error) {
	var resp model.ResponseDto
	if err != nil {
		resp.Err = err
	}
	if err := c.WriteJSON(resp); err != nil {
		v.log.Err(err).Msg("failed to send error message on ws")
	}
}

func (v *View) newHTMLMessage(ctx context.Context, resp model.ResponseDto) ([]byte, error) {
	b := bytes.NewBuffer(nil)
	msgComponent := Message(resp)
	wb := bufio.NewWriter(b)
	if err := msgComponent.Render(ctx, wb); err != nil {
		return nil, err
	}
	wb.Flush()
	return b.Bytes(), nil
}

func (v *View) ProcessConn(c *websocket.Conn) {
	v.log.Info().Str("addr", c.LocalAddr().String()).Msg("new conn")
	c.SetCloseHandler(v.closeHandler)

	defer func() {
		if err := c.Close(); err != nil {
			v.log.Warn().Err(err).Msg("failed to close websocket conn")
			return
		}
		v.log.Info().Msg("conn closed")
	}()

	authCookie := c.Cookies(shared.AuthCookieName, unauthorizedCookie)
	if authCookie == unauthorizedCookie {
		// TODO: add check for cookie on front
		log.Error().Msg("chat error")
		v.respondError(c, shared.ErrInvalidCredentials)
		return
	}

	ctx := context.Background()
	sessionID, err := v.service.InsertSession(ctx, authCookie)
	if err != nil {
		log.Err(err).Msg("chat error")
		v.respondError(c, err)
		return
	}

	var (
		resp   model.ResponseDto
		req    = model.QueryDto{SessionID: sessionID}
		desc   = strings.Builder{}
		cancel = make(chan int64, 1)
		out    = make(chan model.ResponseDto)
	)

	for {
		if err := c.ReadJSON(&req); err != nil {
			log.Err(err).Msg("chat error")
			v.respondError(c, err)
			return
		}

		queryID, err := v.service.InsertQuery(ctx, req)
		if err != nil {
			log.Err(err).Msg("chat error")
			v.respondError(c, err)
			return
		}
		req.ID = queryID

		go v.service.ProcessMessage(ctx, req, out, cancel)

	chunks:
		for {
			select {
			case chunk, ok := <-out:
				// если закончили читать
				if !ok {
					resp.Description = desc.String()
					desc.Reset()
					break chunks
				}

				if chunk.Charts != nil {
					// если первое сообщение с графиками, которое должно быть цельным
					copy(resp.Charts, chunk.Charts)
				} else {
					// если остальные сообщения с описанием, которые идут по токенам
					desc.WriteString(chunk.Description)
				}

				msg, err := v.newHTMLMessage(ctx, chunk)
				if err != nil {
					log.Err(err).Msg("chat error")
					v.respondError(c, err)
					continue
				}

				if err := c.WriteMessage(websocket.TextMessage, msg); err != nil {
					v.log.Err(err).Msg("write message chunk")
					v.respondError(c, err)
					continue
				}
			}
		}

		// TODO: save resp to db
	}
}
