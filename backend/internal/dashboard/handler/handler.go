package handler

import (
	"context"

	"github.com/larek-tech/innohack/backend/internal/dashboard/model"
)

type dashboardController interface {
	GetCharts(ctx context.Context, filter model.Filter) (model.ChartReport, error)
}

type Handler struct {
	ctrl dashboardController
}

func New(ctrl dashboardController) *Handler {
	return &Handler{ctrl: ctrl}
}
