// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: delete_user.sql

package userMethods

import (
	"context"
)

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}
