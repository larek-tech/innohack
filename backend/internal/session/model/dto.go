package model

import (
	"time"

	chatmodel "github.com/larek-tech/innohack/backend/internal/chat/model"
)

type SessionDto struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"createdAt"`
}

type SessionContentDto struct {
	Query    chatmodel.QueryDto    `json:"query"`
	Response chatmodel.ResponseDto `json:"response"`
}
