// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: get_user.sql

package userMethods

import (
	"context"
)

const getUser = `-- name: GetUser :one
SELECT id, username, email FROM users WHERE id = $1
`

type GetUserRow struct {
	ID       int32
	Username string
	Email    string
}

func (q *Queries) GetUser(ctx context.Context, id int32) (GetUserRow, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i GetUserRow
	err := row.Scan(&i.ID, &i.Username, &i.Email)
	return i, err
}
