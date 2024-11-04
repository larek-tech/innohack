package model

import (
	"time"

	"github.com/google/uuid"
)

type Query struct {
	ID        int64     `db:"id"`
	SessionID uuid.UUID `db:"session_id"`
	Prompt    string    `db:"prompt"`
	CreatedAt time.Time `db:"created_at"`
}

type Response struct {
	ID          int64     `db:"id"`
	SessionID   uuid.UUID `db:"session_id"`
	QueryID     int64     `db:"query_id"`
	Sources     []string  `db:"sources"` // s3 link
	Filenames   []string  `db:"filenames"`
	Description string    `db:"description"` // llm response
	CreatedAt   time.Time `db:"created_at"`
}
