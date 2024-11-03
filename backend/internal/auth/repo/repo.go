package repo

import (
	"context"

	"github.com/larek-tech/innohack/backend/internal/auth/model"
	"github.com/larek-tech/innohack/backend/pkg/storage/postgres"
)

type Repo struct {
	pg *postgres.Postgres
}

func New(pg *postgres.Postgres) *Repo {
	return &Repo{pg: pg}
}

const insertUser = `
	insert into users(email, password)
	values ($1, $2)
	returning id;
`

func (r *Repo) InsertUser(ctx context.Context, user model.User) (int64, error) {
	var userID int64

	err := r.pg.Query(ctx, &userID, insertUser, user.Email, user.Password)
	if err != nil {
		return userID, err
	}
	return userID, nil
}

const findUserByEmail = `
	select id, email, password, created_at from users
	where email = $1;
`

func (r *Repo) FindUserByEmail(ctx context.Context, email string) (model.User, error) {
	var user model.User

	err := r.pg.Query(ctx, &user, findUserByEmail, email)
	if err != nil {
		return user, err
	}
	return user, nil
}
