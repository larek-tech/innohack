package model

import (
	"time"

	"github.com/google/uuid"
	chatmodel "github.com/larek-tech/innohack/backend/internal/chat/model"
)

type Session struct {
	ID        uuid.UUID `db:"id"`
	UserID    int64     `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	IsDeleted bool      `db:"is_deleted"`
}

func (s *Session) ToDto() *SessionDto {
	return &SessionDto{
		ID:        s.ID,
		Title:     "Новый чат", // TODO: поменять
		CreatedAt: s.CreatedAt,
	}
}

type SessionContent struct {
	UserID   int64              `db:"user_id"`
	Query    chatmodel.Query    `db:"query"`
	Response chatmodel.Response `db:"response"`
}

func (c *SessionContent) ToDto() *SessionContentDto {
	return &SessionContentDto{
		Query: chatmodel.QueryDto{
			ID:        c.Query.ID,
			Prompt:    c.Query.Prompt,
			CreatedAt: c.Query.CreatedAt,
		},
		Response: chatmodel.ResponseDto{
			QueryID:     c.Query.ID,
			Sources:     c.Response.Sources,
			Filenames:   c.Response.Filenames,
			Description: c.Response.Description,
			CreatedAt:   c.Response.CreatedAt,
		},
	}
}
