// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0
// source: ad.sql

package sqlc

import (
	"context"
	"time"
)

const createAd = `-- name: CreateAd :one
INSERT INTO ads (
  title, content,user_id,created_at,updated_at
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING id, title, content, user_id, created_at, updated_at
`

type CreateAdParams struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UserID    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (q *Queries) CreateAd(ctx context.Context, arg CreateAdParams) (Ad, error) {
	row := q.db.QueryRowContext(ctx, createAd,
		arg.Title,
		arg.Content,
		arg.UserID,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i Ad
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Content,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteAd = `-- name: DeleteAd :exec
DELETE FROM ads
WHERE id = $1
`

func (q *Queries) DeleteAd(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteAd, id)
	return err
}

const getAd = `-- name: GetAd :one
SELECT id, title, content, user_id, created_at, updated_at FROM ads
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetAd(ctx context.Context, id int64) (Ad, error) {
	row := q.db.QueryRowContext(ctx, getAd, id)
	var i Ad
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Content,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getAdsByUserID = `-- name: GetAdsByUserID :many
SELECT id, title, content, user_id, created_at, updated_at FROM ads
WHERE user_id = $1
ORDER BY created_at
`

func (q *Queries) GetAdsByUserID(ctx context.Context, userID int64) ([]Ad, error) {
	rows, err := q.db.QueryContext(ctx, getAdsByUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Ad
	for rows.Next() {
		var i Ad
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Content,
			&i.UserID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listAds = `-- name: ListAds :many
SELECT id, title, content, user_id, created_at, updated_at FROM ads
ORDER BY created_at
`

func (q *Queries) ListAds(ctx context.Context) ([]Ad, error) {
	rows, err := q.db.QueryContext(ctx, listAds)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Ad
	for rows.Next() {
		var i Ad
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Content,
			&i.UserID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAd = `-- name: UpdateAd :exec
UPDATE ads SET title = $2
WHERE id = $1
`

type UpdateAdParams struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}

func (q *Queries) UpdateAd(ctx context.Context, arg UpdateAdParams) error {
	_, err := q.db.ExecContext(ctx, updateAd, arg.ID, arg.Title)
	return err
}