package handler

import (
	"context"

	"github.com/larek-tech/innohack/backend/internal/dashboard/model"
	"go.opentelemetry.io/otel/trace"
)

type dashboardController interface {
	GetCharts(ctx context.Context, filter model.Filter) (model.ChartReport, error)
}

type Handler struct {
	ctrl   dashboardController
	tracer trace.Tracer
}

func New(tracer trace.Tracer, ctrl dashboardController) *Handler {
	return &Handler{
		tracer: tracer,
		ctrl:   ctrl,
	}
}
