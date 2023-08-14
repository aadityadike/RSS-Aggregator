// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: deleteUser.sql

package database

import (
	"context"
)

const deleteUser = `-- name: DeleteUser :exec
DELETE from users WHERE name = $1
`

func (q *Queries) DeleteUser(ctx context.Context, name string) error {
	_, err := q.db.ExecContext(ctx, deleteUser, name)
	return err
}
