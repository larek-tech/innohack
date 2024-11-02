package model

import "time"

type QueryDto struct {
	ID        int64     `json:"id"`
	Prompt    string    `json:"prompt"`
	CreatedAt time.Time `json:"created_at"`
}

type ResponseDto struct {
	QueryID     int64
	Source      string // s3 link
	Filename    string
	Charts      []Chart
	Description string // llm response
	Multipliers []Multiplier
	CreatedAt   time.Time
}

type Chart struct {
	Title       string
	Records     []Record // для отрисовки графа
	Type        string   // пока что bar chart
	Description string   // llm response TODO: возможно, не получится
}

type Record struct {
	X string // формат: квартал - год
	Y float64
}

type Multiplier struct {
	Key   string
	Value float64
}

type SessionContentDto struct {
	Query    QueryDto    `json:"query"`
	Response ResponseDto `json:"response"`
}
