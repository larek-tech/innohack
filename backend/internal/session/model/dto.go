package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/larek-tech/innohack/backend/internal/chat/model"
)

type SessionDto struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"createdAt"`
}

type SessionContentDto struct {
	Query    model.QueryDto    `json:"query"`
	Response model.ResponseDto `json:"response"`
}
