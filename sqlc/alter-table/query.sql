-- name: GetWriter :one
SELECT * FROM writers
WHERE id = $1 LIMIT 1;