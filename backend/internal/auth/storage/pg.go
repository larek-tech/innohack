package storage

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/larek-tech/innohack/backend/internal/auth/service"
	"github.com/larek-tech/innohack/backend/internal/shared"
	"github.com/larek-tech/innohack/backend/pkg/pg"
)

type PG struct {
	pool    *pgxpool.Pool
	queries *pg.Queries
}

// TODO: move from pgx to sql interface

func NewPG(p *pgxpool.Pool, q *pg.Queries) *PG {
	return &PG{
		queries: q,
		pool:    p,
	}
}

func (p *PG) CreateUserWithEmail(ctx context.Context, payload *service.EmailRegisterData) (int64, error) {
	tx, err := p.pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		// TODO: define storage level errors
		return 0, err
	}
	q := p.queries.WithTx(tx)

	defer func() {
		_ = tx.Rollback(ctx)
	}()

	userID, err := q.CreateUserRecord(ctx, pg.CreateUserRecordParams{
		Email:     payload.Email,
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		OtpSecret: uuid.NewString(),
	})
	if err != nil {
		return 0, shared.ErrEmailAlreadyTaken
	}
	err = q.CreateEmailRecord(ctx, pg.CreateEmailRecordParams{
		UserID:   userID,
		Password: payload.Password,
	})
	if err != nil {
		return 0, shared.ErrEmailAlreadyTaken
	}
	tx.Commit(ctx)
	return userID, nil
}

func (p *PG) SaveSessionToken(ctx context.Context, token, userAgent string, userID int64) (int64, error) {
	return p.queries.CreateUserSession(ctx, pg.CreateUserSessionParams{
		SessionID: token,
		UserID:    userID,
		UserAgent: userAgent,
	})
}

func (p *PG) GetUserByEmail(ctx context.Context, email string) (*service.EmailUserData, error) {
	result, err := p.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return &service.EmailUserData{
		UserID:   result.ID,
		Email:    result.Email,
		Password: result.Password,
	}, nil
}
