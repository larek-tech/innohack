package model

import "time"

type Session struct {
	ID        int64     `db:"id"`
	UserID    int64     `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	IsDeleted bool      `db:"is_deleted"`
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
	Source      string    `db:"source"` // s3 link
	Filename    string    `db:"filename"`
	Charts      []byte    `db:"charts"`
	Description string    `db:"description"` // llm response
	Multipliers []byte    `db:"multipliers"`
	CreatedAt   time.Time `db:"created_at"`
}
