package model

import (
	"time"
)

type QueryDto struct {
	ID        int64     `json:"id"`
	Prompt    string    `json:"prompt"`
	CreatedAt time.Time `json:"createdAt"`
}

type ResponseDto struct {
	QueryID     int64     `json:"queryId"`
	Sources     []string  `json:"sources"` // s3 link
	Filenames   []string  `json:"filenames"`
	Description string    `json:"description"` // llm response
	CreatedAt   time.Time `json:"createdAt"`
	Err         error     `json:"error,omitempty"`
	IsLast      bool      `json:"isLast"`
}
