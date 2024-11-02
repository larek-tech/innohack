package model

import "time"

type QueryDto struct {
	ID        int64
	SessionID int64
	Prompt    string
}

type ResponseDto struct {
	QueryID     int64
	Source      string // s3 link
	Filename    string
	Charts      []Chart
	Description string // llm response
	Multipliers []Multiplier
	Err         error
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
