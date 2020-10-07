
-- name: ListUserAuthn :many
SELECT * FROM authn;

-- name: CreateOrUpdateUser :one
INSERT INTO authn (id) VALUES ($1)
ON CONFLICT (id)
DO UPDATE SET
    updated_at = NOW()
RETURNING *;
