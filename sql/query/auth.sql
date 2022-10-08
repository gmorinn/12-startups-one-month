-- name: LoginUser :one
SELECT id, firstname lastname, email, role FROM users
WHERE email = $1
AND password = crypt($2, password)
AND deleted_at IS NULL;

-- name: Signup :one
INSERT INTO users (email, password, firstname, lastname) 
VALUES ($1, crypt($2, gen_salt('bf')), $3, $4)
RETURNING *;

-- name: CheckEmailExist :one
SELECT EXISTS(
    SELECT * FROM users
    WHERE email = $1
    AND deleted_at IS NULL
);