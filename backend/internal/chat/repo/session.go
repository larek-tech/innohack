package repo

import (
	"context"
	"errors"

	"github.com/larek-tech/innohack/backend/internal/chat/model"
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

const getSessionContent = `
	select 
		(q.id, q.session_id, q.prompt, q.created_at) as query,
		(r.id, r.session_id, r.query_id, r.source, r.filename, r.charts, r.description, r.multipliers r.created_at) as response
	from query q
	join
	    response r
		on q.id = r.query_id
	join 
	    session s
		on q.session_id = s.id
	where
	    q.session_id = $1 
	  	and s.is_deleted = false
	order by q.id;
`

func (r *SessionRepo) GetSessionContent(ctx context.Context, sessionID int64) ([]model.SessionContent, error) {
	var content []model.SessionContent
	if err := r.pg.QuerySlice(ctx, &content, getSessionContent, sessionID); err != nil {
		return nil, pkg.WrapErr(err)
	}
	return content, nil
}

const listSessions = `
	select id, user_id, created_at, updated_at from session
	where is_deleted = false and user_id = $1
	order by created_at;
`

func (r *SessionRepo) ListSessions(ctx context.Context, userID int64) ([]model.Session, error) {
	var sessions []model.Session
	if err := r.pg.QuerySlice(ctx, &sessions, listSessions, userID); err != nil {
		return nil, pkg.WrapErr(err)
	}
	return sessions, nil
}

const updateSessionTitle = `
	update session set
		title = $3
	where id = $1 and user_id = $2 and is_deleted = false;
`

func (r *SessionRepo) UpdateSessionTitle(ctx context.Context, sessionID, userID int64, title string) error {
	tag, err := r.pg.Exec(ctx, updateSessionTitle, sessionID, userID, title)
	if err != nil {
		return pkg.WrapErr(err)
	}
	if tag.RowsAffected() == 0 {
		return pkg.WrapErr(errors.New("no rows updated"))
	}
	return nil
}
