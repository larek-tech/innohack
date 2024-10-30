package storage

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/larek-tech/innohack/backend/internal/auth/model"
	"github.com/larek-tech/innohack/backend/internal/shared"
	"github.com/larek-tech/innohack/backend/pkg/pg"
	"github.com/rs/zerolog/log"
)

type PG struct {
	pool    *pgxpool.Pool
	queries *pg.Queries
}

func NewPG(p *pgxpool.Pool, q *pg.Queries) *PG {
	return &PG{
		queries: q,
		pool:    p,
	}
}

func (p *PG) CreateWithEmail(ctx context.Context, payload *model.EmailRegisterData) (int64, error) {
	tx, err := p.pool.BeginTx(ctx, pgx.TxOptions{})
	defer func() {
		if txErr := tx.Rollback(ctx); txErr != nil {
			log.Err(err).Msg("unable to rollback tx in Create withEmail")
		}
	}()
	if err != nil {
		return 0, err
	}
	q := p.queries.WithTx(tx)

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
	if err = tx.Commit(ctx); err != nil {
		log.Err(err)
		return 0, err
	}

	return userID, nil
}

func (p *PG) Save(ctx context.Context, token, userAgent string, userID int64) (int64, error) {
	return p.queries.CreateUserSession(ctx, pg.CreateUserSessionParams{
		SessionID: token,
		UserID:    userID,
		UserAgent: userAgent,
	})
}

func (p *PG) GetByEmail(ctx context.Context, email string) (*model.EmailUserData, error) {
	result, err := p.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return &model.EmailUserData{
		UserID:   result.ID,
		Email:    result.Email,
		Password: result.Password,
	}, nil
}
