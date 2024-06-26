// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: create_user.sql

package userMethods

import (
	"context"
)

const createUser = `-- name: CreateUser :exec
INSERT INTO users (username, email, password) VALUES ($1, $2, $3)
`

type CreateUserParams struct {
	Username string
	Email    string
	Password string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.ExecContext(ctx, createUser, arg.Username, arg.Email, arg.Password)
	return err
}
