-- name: GetUser :one
SELECT id, username, email FROM users WHERE id = $1;