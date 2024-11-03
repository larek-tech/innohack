package repo

import (
	"context"

	"github.com/larek-tech/innohack/backend/internal/chat/model"
	"github.com/larek-tech/innohack/backend/pkg/storage/postgres"
)

type QueryRepo struct {
	pg *postgres.Postgres
}

func NewQueryRepo(pg *postgres.Postgres) *QueryRepo {
	return &QueryRepo{pg: pg}
}

const insertQuery = `
	insert into query (session_id, prompt)
	values ($1, $2)
	returning id;
`

func (r *QueryRepo) InsertQuery(ctx context.Context, query model.Query) (int64, error) {
	var queryID int64
	err := r.pg.Query(ctx, &queryID, insertQuery, query.SessionID, query.Prompt)
	if err != nil {
		return queryID, err
	}
	return queryID, nil
}
