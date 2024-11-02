package model

type Response struct {
	SessionID   int64
	Records     []Record // для отрисовки графа
	Description string   // llm response
}

type Record struct {
	Source   string // s3 link
	Filename string
	X        string // формат: квартал - год
	Y        float64
	Type     string // пока что bar chart
}
