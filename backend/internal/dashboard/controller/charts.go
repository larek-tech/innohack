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

	titleCharts := report.GetCharts()
	multipliers := report.GetMultipliers()

	res := model.ChartReport{
		Summary:     report.GetSummary(),
		Charts:      map[string][]model.Chart{},
		Multipliers: make([]model.Multiplier, len(multipliers)),
		Legend:      map[string]string{},
		StartDate:   filter.StartDate.Year(),
		EndDate:     filter.EndDate.Year(),
	}
	for title, chartList := range titleCharts {
		res.Charts[title] = make([]model.Chart, len(chartList.Charts))
		for idx, chart := range chartList.Charts {
			res.Charts[title][idx] = model.ChartFromPb(chart)
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
