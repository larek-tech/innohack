package repo

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/larek-tech/innohack/backend/internal/session/model"
	"github.com/larek-tech/innohack/backend/pkg/storage/postgres"
)

type Repo struct {
	pg *postgres.Postgres
}

func New(pg *postgres.Postgres) *Repo {
	return &Repo{pg: pg}
}

const insertSession = `
	insert into session(id, user_id)
	values ($1, $2)
`

func (r *Repo) InsertSession(ctx context.Context, sessionID uuid.UUID, userID int64) error {
	_, err := r.pg.Exec(ctx, insertSession, sessionID, userID)
	return err
}

const getSessionByID = `
	select id, user_id, created_at, updated_at, is_deleted from session
	where id = $1;
`

func (r *Repo) GetSessionByID(ctx context.Context, sessionID uuid.UUID) (model.Session, error) {
	var session model.Session
	if err := r.pg.Query(ctx, &session, getSessionByID, sessionID); err != nil {
		return session, err
	}
	return session, nil
}

const getSessionContent = `
	select 
	    s.user_id as user_id,
		(q.id, q.session_id, q.prompt, q.created_at) as query,
		(r.id, r.session_id, r.query_id, r.sources, r.filenames, r.description, r.created_at) as response
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

func (r *Repo) GetSessionContent(ctx context.Context, sessionID uuid.UUID) ([]model.SessionContent, error) {
	var content []model.SessionContent
	if err := r.pg.QuerySlice(ctx, &content, getSessionContent, sessionID); err != nil {
		return nil, err
	}
	return content, nil
}

const listSessions = `
	select id, user_id, created_at, updated_at from session
	where is_deleted = false and user_id = $1
	order by created_at;
`

func (r *Repo) ListSessions(ctx context.Context, userID int64) ([]model.Session, error) {
	var sessions []model.Session
	if err := r.pg.QuerySlice(ctx, &sessions, listSessions, userID); err != nil {
		return nil, err
	}
	return sessions, nil
}

const updateSessionTitle = `
	update session set
		title = $3
	where id = $1 and user_id = $2 and is_deleted = false;
`

func (r *Repo) UpdateSessionTitle(ctx context.Context, sessionID uuid.UUID, userID int64, title string) error {
	tag, err := r.pg.Exec(ctx, updateSessionTitle, sessionID, userID, title)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return errors.New("no rows updated")
	}
	return nil
}

const deleteSession = `
	delete from session
	where id = $1;
`

func (r *Repo) DeleteSession(ctx context.Context, sessionID uuid.UUID) error {
	_, err := r.pg.Exec(ctx, deleteSession, sessionID)
	return err
}
