package repo

import (
	"context"

	"github.com/larek-tech/innohack/backend/internal/chat/model"
	"github.com/larek-tech/innohack/backend/pkg"
	"github.com/larek-tech/innohack/backend/pkg/storage/postgres"
)

type ResponseRepo struct {
	pg *postgres.Postgres
}

func NewResponseRepo(pg *postgres.Postgres) *ResponseRepo {
	return &ResponseRepo{pg: pg}
}

const insertResponse = `
	insert into response(
		session_id, 
		query_id, 
		sources, 
		filenames, 
		charts, 
		description,
		multipliers
	)
	values ($1, $2, $3, $4, $5, $6, $7);
`

func (r *ResponseRepo) InsertResponse(ctx context.Context, data model.Response) error {
	_, err := r.pg.Exec(
		ctx,
		insertResponse,
		data.SessionID,
		data.QueryID,
		data.Sources,
		data.Filenames,
		data.Charts,
		data.Description,
		data.Multipliers,
	)
	return pkg.WrapErr(err)
}
