// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: query.sql

package querytest

import (
	"context"
)

const listBaz = `-- name: ListBaz :iter
SELECT one, two FROM baz
`

func (q *Queries) ListBaz(ctx context.Context, iter func(i Baz) error) error {
	rows, err := q.db.QueryContext(ctx, listBaz)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var i Baz
		if err := rows.Scan(&i.One, &i.Two); err != nil {
			return err
		}
		err = iter(i)
		if err != nil {
			return err
		}
	}
	if err := rows.Close(); err != nil {
		return err
	}
	if err := rows.Err(); err != nil {
		return err
	}
	return nil
}

const listFoo = `-- name: ListFoo :iter
SELECT bar FROM foo
`

func (q *Queries) ListFoo(ctx context.Context, iter func(bar bool) error) error {
	rows, err := q.db.QueryContext(ctx, listFoo)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var bar bool
		if err := rows.Scan(&bar); err != nil {
			return err
		}
		err = iter(bar)
		if err != nil {
			return err
		}
	}
	if err := rows.Close(); err != nil {
		return err
	}
	if err := rows.Err(); err != nil {
		return err
	}
	return nil
}
