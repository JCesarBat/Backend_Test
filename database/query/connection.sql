
-- name: GetConnectionID :one
SELECT * FROM connections
where user_id = $1 LIMIT 1;

-- name: UpdateConnection :one
UPDATE connections
SET
    active =$1
where
        user_id = $2
RETURNING *;