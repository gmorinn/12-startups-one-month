// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: auth.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

const checkEmailExist = `-- name: CheckEmailExist :one
SELECT EXISTS(
    SELECT id, created_at, updated_at, deleted_at, email, password, firstname, lastname, role, age, sexe, goals, ideal_partners, profile_picture, is_premium, city, ask, badge, formule FROM users
    WHERE email = $1
    AND deleted_at IS NULL
)
`

func (q *Queries) CheckEmailExist(ctx context.Context, email string) (bool, error) {
	row := q.db.QueryRowContext(ctx, checkEmailExist, email)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const loginUser = `-- name: LoginUser :one
SELECT id, firstname lastname, email, role, age, sexe, goals, ideal_partners,
    profile_picture, is_premium, city, ask, badge, formule FROM users
WHERE email = $1
AND password = crypt($2, password)
AND deleted_at IS NULL
`

type LoginUserParams struct {
	Email string `json:"email"`
	Crypt string `json:"crypt"`
}

type LoginUserRow struct {
	ID             uuid.UUID      `json:"id"`
	Lastname       sql.NullString `json:"lastname"`
	Email          string         `json:"email"`
	Role           []Role         `json:"role"`
	Age            sql.NullInt32  `json:"age"`
	Sexe           Sexe           `json:"sexe"`
	Goals          []Goals        `json:"goals"`
	IdealPartners  sql.NullString `json:"ideal_partners"`
	ProfilePicture sql.NullString `json:"profile_picture"`
	IsPremium      sql.NullBool   `json:"is_premium"`
	City           sql.NullString `json:"city"`
	Ask            int32          `json:"ask"`
	Badge          bool           `json:"badge"`
	Formule        Formule        `json:"formule"`
}

func (q *Queries) LoginUser(ctx context.Context, arg LoginUserParams) (LoginUserRow, error) {
	row := q.db.QueryRowContext(ctx, loginUser, arg.Email, arg.Crypt)
	var i LoginUserRow
	err := row.Scan(
		&i.ID,
		&i.Lastname,
		&i.Email,
		pq.Array(&i.Role),
		&i.Age,
		&i.Sexe,
		pq.Array(&i.Goals),
		&i.IdealPartners,
		&i.ProfilePicture,
		&i.IsPremium,
		&i.City,
		&i.Ask,
		&i.Badge,
		&i.Formule,
	)
	return i, err
}

const signup = `-- name: Signup :one
INSERT INTO users (email, password) 
VALUES ($1, crypt($2, gen_salt('bf')))
RETURNING id, created_at, updated_at, deleted_at, email, password, firstname, lastname, role, age, sexe, goals, ideal_partners, profile_picture, is_premium, city, ask, badge, formule
`

type SignupParams struct {
	Email string `json:"email"`
	Crypt string `json:"crypt"`
}

func (q *Queries) Signup(ctx context.Context, arg SignupParams) (User, error) {
	row := q.db.QueryRowContext(ctx, signup, arg.Email, arg.Crypt)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.Email,
		&i.Password,
		&i.Firstname,
		&i.Lastname,
		pq.Array(&i.Role),
		&i.Age,
		&i.Sexe,
		pq.Array(&i.Goals),
		&i.IdealPartners,
		&i.ProfilePicture,
		&i.IsPremium,
		&i.City,
		&i.Ask,
		&i.Badge,
		&i.Formule,
	)
	return i, err
}
