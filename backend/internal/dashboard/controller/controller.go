package controller

import (
	"github.com/larek-tech/innohack/backend/internal/analytics/pb"
	"google.golang.org/grpc"
)

type Controller struct {
	analytics pb.AnalyticsClient
}

func New(grpcConn *grpc.ClientConn) *Controller {
	return &Controller{
		analytics: pb.NewAnalyticsClient(grpcConn),
	}
}
