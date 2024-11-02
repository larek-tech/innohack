package repo

import (
	"context"

	"github.com/larek-tech/innohack/backend/internal/auth/model"
	"github.com/larek-tech/innohack/backend/pkg"
	"github.com/larek-tech/innohack/backend/pkg/storage/postgres"
)

type Repo struct {
	pg *postgres.Postgres
}

func New(pg *postgres.Postgres) *Repo {
	return &Repo{pg: pg}
}

const insertUser = `
	insert into user(email, password)
	values ($1, $2);
`

func (r *Repo) InsertUser(ctx context.Context, user model.User) (int64, error) {
	var userID int64

	err := r.pg.Query(ctx, &userID, insertUser, user.Email, user.Password)
	if err != nil {
		return userID, pkg.WrapErr(err)
	}
	return userID, nil
}

const findUserByEmail = `
	select (id, password) from user
	where email = $1;
`

func (r *Repo) FindUserByEmail(ctx context.Context, email string) (model.User, error) {
	var user model.User

	err := r.pg.Query(ctx, &user, findUserByEmail, email)
	if err != nil {
		return user, pkg.WrapErr(err)
	}

	user.Email = email
	return user, nil
}
