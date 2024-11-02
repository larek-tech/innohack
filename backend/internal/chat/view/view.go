package view

import (
	"bufio"
	"bytes"
	"context"
	"fmt"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/websocket/v2"
	"github.com/larek-tech/innohack/backend/internal/chat/service"

	"github.com/rs/zerolog"
)

type View struct {
	s   *service.Service
	log *zerolog.Logger
}

func New(log *zerolog.Logger, s *service.Service) *View {
	return &View{
		s:   s,
		log: log,
	}
}

func (v *View) ChatPage(c *fiber.Ctx) error {

	if websocket.IsWebSocketUpgrade(c) {
		c.Locals("allowed", true)
		fmt.Println("a")
		return c.Next()
	} else {
		v.log.Info().Msg("can't upgrade")
	}

	return adaptor.HTTPHandler(
		templ.Handler(ChatPage()),
	)(c)
}

func (v *View) closeHandler(code int, text string) error {
	v.log.Info().Int("code", code).Str("text", text).Msg("close handler")
	return nil
}

// processes websocket connection by starting new goroutine
func (v *View) Message(c *websocket.Conn) {
	ctx := context.TODO()
	v.log.Info().Str("addr", c.LocalAddr().String()).Msg("new conn")
	c.SetCloseHandler(v.closeHandler)
	defer func() {
		if err := c.Close(); err != nil {
			v.log.Warn().Err(err).Msg("failed to close websocket conn")
		}
		v.log.Info().Msg("conn closed")
	}()

	for {
		var (
			req map[string]interface{}
			b   bytes.Buffer
		)
		err := c.ReadJSON(&req)
		if err != nil {
			v.log.Info().Msg("unable to map json")
			break
		}
		response, err := v.s.ProcessMessage(ctx, req["content"].(string))
		if err != nil {
			v.log.Info().Msg("unable to map json")
			break
		}
		msgComponent := Message(response)
		wb := bufio.NewWriter(&b)
		msgComponent.Render(context.Background(), wb)
		err = wb.Flush()
		if err != nil {
			v.log.Info().Msg("unable to map json")
			break
		}

		err = c.WriteMessage(1, b.Bytes())
		if err != nil {
			v.log.Info().Msg("unable to send data")
			break
		}

	}
}
