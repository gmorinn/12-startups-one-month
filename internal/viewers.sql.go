// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: viewers.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const checkViewByID = `-- name: CheckViewByID :one
SELECT EXISTS (
    SELECT 1
    FROM viewers
    WHERE id = $1
    AND deleted_at IS NULL
)
`

func (q *Queries) CheckViewByID(ctx context.Context, id uuid.UUID) (bool, error) {
	row := q.db.QueryRowContext(ctx, checkViewByID, id)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const createView = `-- name: CreateView :one
INSERT INTO viewers (user_id_viewer, profil_id_viewed)
VALUES ($1, $2)
RETURNING id, created_at, updated_at, deleted_at, user_id_viewer, profil_id_viewed, date_viewed
`

type CreateViewParams struct {
	UserIDViewer   uuid.UUID `json:"user_id_viewer"`
	ProfilIDViewed uuid.UUID `json:"profil_id_viewed"`
}

func (q *Queries) CreateView(ctx context.Context, arg CreateViewParams) (Viewer, error) {
	row := q.db.QueryRowContext(ctx, createView, arg.UserIDViewer, arg.ProfilIDViewed)
	var i Viewer
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.UserIDViewer,
		&i.ProfilIDViewed,
		&i.DateViewed,
	)
	return i, err
}

const deleteViewByID = `-- name: DeleteViewByID :exec
UPDATE
    viewers
SET
    deleted_at = NOW()
WHERE 
    id = $1
`

func (q *Queries) DeleteViewByID(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteViewByID, id)
	return err
}

const getAllView = `-- name: GetAllView :many
SELECT id, created_at, updated_at, deleted_at, user_id_viewer, profil_id_viewed, date_viewed FROM viewers
WHERE deleted_at IS NULL
ORDER BY
CASE WHEN $1::bool THEN date_viewed END asc,
CASE WHEN $2::bool THEN date_viewed END desc
LIMIT $4 OFFSET $3
`

type GetAllViewParams struct {
	DateViewedAsc  bool  `json:"date_viewed_asc"`
	DateViewedDesc bool  `json:"date_viewed_desc"`
	Offset         int32 `json:"offset"`
	Limit          int32 `json:"limit"`
}

func (q *Queries) GetAllView(ctx context.Context, arg GetAllViewParams) ([]Viewer, error) {
	rows, err := q.db.QueryContext(ctx, getAllView,
		arg.DateViewedAsc,
		arg.DateViewedDesc,
		arg.Offset,
		arg.Limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Viewer{}
	for rows.Next() {
		var i Viewer
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.UserIDViewer,
			&i.ProfilIDViewed,
			&i.DateViewed,
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

const getViewByID = `-- name: GetViewByID :one
SELECT id, created_at, updated_at, deleted_at, user_id_viewer, profil_id_viewed, date_viewed FROM viewers
WHERE id = $1
AND deleted_at IS NULL
LIMIT 1
`

func (q *Queries) GetViewByID(ctx context.Context, id uuid.UUID) (Viewer, error) {
	row := q.db.QueryRowContext(ctx, getViewByID, id)
	var i Viewer
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.UserIDViewer,
		&i.ProfilIDViewed,
		&i.DateViewed,
	)
	return i, err
}

const getViewsByUserID = `-- name: GetViewsByUserID :many
SELECT id, created_at, updated_at, deleted_at, user_id_viewer, profil_id_viewed, date_viewed FROM viewers
WHERE profil_id_viewed = $1
AND deleted_at IS NULL
AND date_viewed > NOW() - INTERVAL '14 days'
ORDER BY date_viewed DESC
`

func (q *Queries) GetViewsByUserID(ctx context.Context, profilIDViewed uuid.UUID) ([]Viewer, error) {
	rows, err := q.db.QueryContext(ctx, getViewsByUserID, profilIDViewed)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Viewer{}
	for rows.Next() {
		var i Viewer
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.UserIDViewer,
			&i.ProfilIDViewed,
			&i.DateViewed,
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
