-- name: RegisterUser :one
INSERT INTO users
    (name, email, role, phone,password)
    VALUES
    ($1,$2,$3,$4,$5)
    RETURNING id;

-- name: GetUserById :one
SELECT * FROM users WHERE id = $1;

-- name: GetUserByToken :one
SELECT * FROM users INNER JOIN tokens on
    users.id = tokens.user_id
    WHERE tokens.hash = $1
    AND tokens.scope = $2
    AND tokens.expiry > $3;

-- name: GetUserByCredentials :one
SELECT * FROM users
    WHERE email = $1
    AND password = $2;

-- name: DeleteUser :one
DELETE FROM users WHERE id = $1 RETURNING *;
