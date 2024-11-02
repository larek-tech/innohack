package model

import (
	"encoding/json"
	"time"

	"github.com/larek-tech/innohack/backend/pkg"
)

type Session struct {
	ID        int64     `db:"id"`
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

type Query struct {
	ID        int64     `db:"id"`
	SessionID int64     `db:"session_id"`
	Prompt    string    `db:"prompt"`
	CreatedAt time.Time `db:"created_at"`
}

type Response struct {
	ID          int64     `db:"id"`
	SessionID   int64     `db:"session_id"`
	QueryID     int64     `db:"query_id"`
	Sources     []string  `db:"sources"` // s3 link
	Filenames   []string  `db:"filenames"`
	Charts      []byte    `db:"charts"`
	Description string    `db:"description"` // llm response
	Multipliers []byte    `db:"multipliers"`
	CreatedAt   time.Time `db:"created_at"`
}

type SessionContent struct {
	Query    Query    `db:"query"`
	Response Response `db:"response"`
}

func (c *SessionContent) ToDto() (*SessionContentDto, error) {
	var (
		charts      []Chart
		multipliers []Multiplier
	)

	if err := json.Unmarshal(c.Response.Charts, &charts); err != nil {
		return nil, pkg.WrapErr(err, "dto charts")
	}

	if err := json.Unmarshal(c.Response.Multipliers, &multipliers); err != nil {
		return nil, pkg.WrapErr(err, "dto multipliers")
	}

	return &SessionContentDto{
		Query: QueryDto{
			ID:        c.Query.ID,
			Prompt:    c.Query.Prompt,
			CreatedAt: c.Query.CreatedAt,
		},
		Response: ResponseDto{
			QueryID:     c.Query.ID,
			Sources:     c.Response.Sources,
			Filenames:   c.Response.Filenames,
			Charts:      charts,
			Description: c.Response.Description,
			Multipliers: multipliers,
			CreatedAt:   c.Response.CreatedAt,
		},
	}, nil
}
