package model

import (
	"fmt"
	"time"

	"github.com/larek-tech/innohack/backend/internal/analytics/pb"
)

type QueryDto struct {
	ID        int64     `json:"id"`
	Prompt    string    `json:"prompt"`
	CreatedAt time.Time `json:"created_at"`
}

type ResponseDto struct {
	QueryID     int64        `json:"query_id"`
	Sources     []string     `json:"sources"` // s3 link
	Filenames   []string     `json:"filenames"`
	Charts      []Chart      `json:"charts"`
	Description string       `json:"description"` // llm response
	Multipliers []Multiplier `json:"multipliers"`
	CreatedAt   time.Time    `json:"created_at"`
	IsLast      bool
}

type Chart struct {
	GID         string    `json:"gid"`
	DataGID     string    `json:"data-gid"`
	Title       string    `json:"title"`
	Records     []Record  `json:"records"`     // для отрисовки графа
	Type        ChartType `json:"type"`        // пока что bar chart
	Description string    `json:"description"` // llm response TODO: возможно, не получится
}

func (c Chart) GetType() string {
	switch c.Type {
	case BarChart:
		return "bar"
	case PieChart:
		return "pie"
	case UndefinedChart:
		panic(fmt.Sprintf("unexpected model.ChartType: %#v", c.Type))
	default:
		panic(fmt.Sprintf("unexpected model.ChartType: %#v", c.Type))
	}
}

func ChartsFromPb(charts []*pb.Chart) []Chart {
	res := make([]Chart, len(charts))
	for i, chart := range charts {
		res[i] = Chart{
			Title:       chart.GetTitle(),
			Records:     RecordsFromPb(chart.GetRecords()),
			Type:        ChartType(chart.GetType()),
			Description: chart.GetDescription(),
		}
	}
	return res
}

type ChartType uint8

const (
	UndefinedChart ChartType = iota
	BarChart
	PieChart
)

type Record struct {
	X string  `json:"x"` // формат: квартал - год
	Y float64 `json:"y"`
}

func RecordsFromPb(records []*pb.Record) []Record {
	res := make([]Record, len(records))
	for j, record := range records {
		res[j] = Record{
			X: record.GetX(),
			Y: record.GetY(),
		}
	}
	return res
}

type Multiplier struct {
	Key   string  `json:"key"`
	Value float64 `json:"value"`
}

func MultipliersFromPb(multipliers []*pb.Multiplier) []Multiplier {
	res := make([]Multiplier, len(multipliers))
	for idx, multiplier := range multipliers {
		res[idx] = Multiplier{
			Key:   multiplier.GetKey(),
			Value: multiplier.GetValue(),
		}
	}
	return res
}

type SessionContentDto struct {
	Query    QueryDto    `json:"query"`
	Response ResponseDto `json:"response"`
}
