-- name: GetAllAvis :many
SELECT * FROM avis
WHERE deleted_at IS NULL
ORDER BY
CASE WHEN sqlc.arg('note_asc')::bool THEN note END asc,
CASE WHEN sqlc.arg('note_desc')::bool THEN note END desc
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: GetAvisByID :one
SELECT * FROM avis
WHERE id = $1
AND deleted_at IS NULL
LIMIT 1;

-- name: GetAvisByUserID :many
SELECT * FROM avis
WHERE user_id_target = $1
AND deleted_at IS NULL;

-- name: DeleteAvisByID :exec
UPDATE
    avis
SET
    deleted_at = NOW()
WHERE 
    id = $1;

-- name: UpdateAvis :exec
UPDATE 
    avis
SET
    comment = $2,
    note = $3,
    updated_at = NOW()
WHERE
    id = $1;

-- name: CreateAvis :one
INSERT INTO avis (user_id_target, user_id_writer, comment, note)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: CheckAvisByID :one
SELECT EXISTS (
    SELECT 1
    FROM avis
    WHERE id = $1
    AND deleted_at IS NULL
);

-- name: CheckIfAvisExist :one
SELECT EXISTS (
    SELECT 1
    FROM avis
    WHERE user_id_target = $1
    AND user_id_writer = $2
    AND deleted_at IS NULL
);