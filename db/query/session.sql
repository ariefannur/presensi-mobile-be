-- name: CreateSession :one
INSERT INTO "Sessions" (
  user_id, token, device
) VALUES (
  $1, $2, $3
)
RETURNING id, user_id, token, time, device, status;