package controller

import (
	"context"

	"github.com/larek-tech/innohack/backend/internal/analytics/pb"
	"github.com/larek-tech/innohack/backend/internal/dashboard/model"
)

func (ctrl *Controller) GetCharts(ctx context.Context, filter model.Filter) (model.ChartReport, error) {
	report, err := ctrl.analytics.GetCharts(ctx, &pb.Filter{
		StartDate: int64(filter.StartDate.Year()),
		EndDate:   int64(filter.EndDate.Year()),
	})
	if err != nil {
		return model.ChartReport{}, err
	}

	info := report.GetInfo()
	multipliers := report.GetMultipliers()

	res := model.ChartReport{
		Summary:     report.GetSummary(),
		Info:        map[string]model.ChartsLegend{},
		Multipliers: make([]model.Multiplier, len(multipliers)),
		StartDate:   filter.StartDate.Year(),
		EndDate:     filter.EndDate.Year(),
	}
	for title, chartLegend := range info {
		res.Info[title] = model.ChartsLegend{
			Charts: make([]model.Chart, len(chartLegend.Charts)),
			Legend: map[string]string{},
		}

		for idx := range len(chartLegend.Charts) {
			res.Info[title].Charts[idx] = model.ChartFromPb(chartLegend.Charts[idx])
		}
		for color, desc := range chartLegend.Legend {
			res.Info[title].Legend[chartLegend.Legend[color]] = desc
		}
	}

	for idx := range len(multipliers) {
		res.Multipliers[idx] = model.Multiplier{
			Key:   multipliers[idx].Key,
			Value: multipliers[idx].Value,
		}
	}

	return res, nil
}
