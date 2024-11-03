package model

import (
	"time"

	"github.com/google/uuid"
	chatmodel "github.com/larek-tech/innohack/backend/internal/chat/model"
)

type SessionDto struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"createdAt"`
}

type SessionContentDto struct {
	Query    chatmodel.QueryDto    `json:"query"`
	Response chatmodel.ResponseDto `json:"response"`
}
