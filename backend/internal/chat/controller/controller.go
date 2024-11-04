package controller

import (
	"context"

	"github.com/larek-tech/innohack/backend/internal/analytics/pb"
	"github.com/larek-tech/innohack/backend/internal/chat/model"
	"google.golang.org/grpc"
)

type queryRepo interface {
	InsertQuery(ctx context.Context, data model.Query) (int64, error)
}

type responseRepo interface {
	InsertResponse(ctx context.Context, data model.Response) error
}

type Controller struct {
	jwtSecret string
	qr        queryRepo
	rr        responseRepo
	analytics pb.AnalyticsClient
}

func New(jwtSecret string, grpcConn *grpc.ClientConn, qr queryRepo, rr responseRepo) *Controller {
	return &Controller{
		jwtSecret: jwtSecret,
		qr:        qr,
		rr:        rr,
		analytics: pb.NewAnalyticsClient(grpcConn),
	}
}
