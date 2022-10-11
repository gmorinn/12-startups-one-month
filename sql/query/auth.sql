-- name: LoginUser :one
SELECT id, firstname lastname, email, role, age, sexe, goals, ideal_partners,
    profile_picture, is_premium, city, ask, badge, formule FROM users
WHERE email = $1
AND password = crypt($2, password)
AND deleted_at IS NULL;

-- name: Signup :one
INSERT INTO users (email, password) 
VALUES ($1, crypt($2, gen_salt('bf')))
RETURNING *;

-- name: CheckEmailExist :one
SELECT EXISTS(
    SELECT * FROM users
    WHERE email = $1
    AND deleted_at IS NULL
);