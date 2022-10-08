-- name: GetAllUser :many
SELECT * FROM users
WHERE deleted_at IS NULL
ORDER BY
  CASE WHEN sqlc.arg('firstname_asc')::bool THEN firstname END asc,
  CASE WHEN sqlc.arg('firstname_desc')::bool THEN firstname END desc,
  CASE WHEN sqlc.arg('lastname_asc')::bool THEN lastname END asc,
  CASE WHEN sqlc.arg('lastname_desc')::bool THEN lastname END desc,
  CASE WHEN sqlc.arg('email_asc')::bool THEN email END asc,
  CASE WHEN sqlc.arg('email_desc')::bool THEN email END desc,
  CASE WHEN sqlc.arg('created_at_asc')::bool THEN created_at END asc,
  CASE WHEN sqlc.arg('created_at_desc')::bool THEN created_at END desc,
  CASE WHEN sqlc.arg('role_asc')::bool THEN role END asc,
  CASE WHEN sqlc.arg('role_desc')::bool THEN role END desc
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1
AND deleted_at IS NULL
LIMIT 1;

-- name: DeleteUserByID :exec
UPDATE
    users
SET
    deleted_at = NOW()
WHERE 
    id = $1;

-- name: UpdateUser :exec
UPDATE 
    users
SET
    firstname = $2,
    lastname = $3,
    email = $4,
    updated_at = NOW()
WHERE
    id = $1;

-- name: UpdateRole :exec
UPDATE 
    users
SET
    role = $2,
    updated_at = NOW()
WHERE
    id = $1;

-- name: CheckUserByID :one
SELECT EXISTS(
    SELECT * FROM users
    WHERE id = $1
    AND deleted_at IS NULL
);