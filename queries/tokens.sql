-- name: CreateToken :one
INSERT INTO tokens
    (hash,user_id,expiry,scope)
    VALUES
    ($1,$2,$3,$4)
    RETURNING *;

-- name: DeleteAllTokensForUser :one
DELETE FROM tokens WHERE user_id = $1 RETURNING *;

-- name: DeleteScopeTokenForUser :one
DELETE FROM tokens
       WHERE user_id = $1
       AND scope = $2
       RETURNING *;

