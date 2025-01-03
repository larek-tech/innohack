package model

import (
	"fmt"
	"time"

	"github.com/larek-tech/innohack/backend/internal/analytics/pb"
)

type Filter struct {
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
}

type ChartReport struct {
	Summary     string                  `json:"summary"`
	Info        map[string]ChartsLegend `json:"info"`
	Multipliers []Multiplier            `json:"multipliers"`
	StartDate   int                     `json:"startDate"`
	EndDate     int                     `json:"endDate"`
}

type ChartsLegend struct {
	Charts []Chart           `json:"charts"`
	Legend map[string]string `json:"legend"`
}

type Chart struct {
	Color   string   `json:"color"`
	Records []Record `json:"records"` // для отрисовки графа
	Type    string   `json:"type"`    // пока что bar chart
}

func ChartFromPb(in *pb.Chart) Chart {
	inRecords := in.GetRecords()
	records := make([]Record, len(inRecords))
	for idx := range len(inRecords) {
		records[idx] = Record{inRecords[idx].X, inRecords[idx].Y}
	}

	return Chart{
		Color:   in.GetColor(),
		Records: records,
		Type:    ChartType(in.GetType()).ToString(),
	}
}

type ChartType uint8

const (
	UndefinedChart ChartType = iota
	BarChart
	PieChart
	LineChart
)

func (t ChartType) ToString() string {
	switch t {
	case BarChart:
		return "bar"
	case PieChart:
		return "pie"
	case LineChart:
		return "line"
	default:
		panic(fmt.Sprintf("unexpected model.ChartType: %#v", t))
	}
}

type Record struct {
	X string  `json:"x"` // формат: квартал - год
	Y float64 `json:"y"`
}

type Multiplier struct {
	Key   string  `json:"key"`
	Value float64 `json:"value"`
}
