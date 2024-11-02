package repo

import (
	"context"

	"github.com/larek-tech/innohack/backend/pkg"
	"github.com/larek-tech/innohack/backend/pkg/storage/postgres"
)

type SessionRepo struct {
	pg *postgres.Postgres
}

func NewSessionRepo(pg *postgres.Postgres) *SessionRepo {
	return &SessionRepo{pg: pg}
}

const insertSession = `
	insert into session(user_id)
	values ($1)
	returning id;
`

func (r *SessionRepo) InsertSession(ctx context.Context, userID int64) (int64, error) {
	var sessionID int64
	err := r.pg.Query(ctx, &sessionID, insertSession, userID)
	if err != nil {
		return sessionID, pkg.WrapErr(err)
	}
	return sessionID, nil
}
