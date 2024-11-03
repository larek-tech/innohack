package controller

import (
	"context"

	"github.com/larek-tech/innohack/backend/internal/analytics/pb"
	"github.com/larek-tech/innohack/backend/internal/dashboard/model"
)

func (ctrl *Controller) GetCharts(ctx context.Context, filter model.Filter) (model.ChartReport, error) {
	report, err := ctrl.analytics.GetCharts(ctx, &pb.Filter{
		StartDate: filter.StartDate.String(),
		EndDate:   filter.EndDate.String(),
	})
	if err != nil {
		return model.ChartReport{}, err
	}

	charts := report.GetCharts()
	multipliers := report.GetMultipliers()

	res := model.ChartReport{
		Description: report.GetDescription(),
		StartDate:   filter.StartDate,
		EndDate:     filter.EndDate,
		Charts:      make([]model.Chart, len(charts)),
		Multipliers: make([]model.Multiplier, len(multipliers)),
	}
	for idx := range len(charts) {
		res.Charts[idx] = model.ChartFromPb(charts[idx])
	}

	for idx := range len(multipliers) {
		res.Multipliers[idx] = model.Multiplier{
			Key:   multipliers[idx].Key,
			Value: multipliers[idx].Value,
		}
	}

	return res, nil
}
