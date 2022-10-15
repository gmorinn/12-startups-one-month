-- name: GetAllView :many
SELECT * FROM viewers
WHERE deleted_at IS NULL
ORDER BY
CASE WHEN sqlc.arg('date_viewed_asc')::bool THEN date_viewed END asc,
CASE WHEN sqlc.arg('date_viewed_desc')::bool THEN date_viewed END desc
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: GetViewByID :one
SELECT * FROM viewers
WHERE id = $1
AND deleted_at IS NULL
LIMIT 1;

-- name: DeleteViewByID :exec
UPDATE
    viewers
SET
    deleted_at = NOW()
WHERE 
    id = $1;

-- name: CreateView :one
INSERT INTO viewers (user_id_viewer, profil_id_viewed)
VALUES ($1, $2)
RETURNING *;

-- name: CheckViewByID :one
SELECT EXISTS (
    SELECT 1
    FROM viewers
    WHERE id = $1
    AND deleted_at IS NULL
);

-- name: GetViewsByUserID :many
SELECT * FROM viewers
WHERE profil_id_viewed = $1
AND deleted_at IS NULL
AND date_viewed > NOW() - INTERVAL '14 days'
ORDER BY date_viewed DESC;