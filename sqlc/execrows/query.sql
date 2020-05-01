-- name: DeleteAuthorN :execrows
DELETE FROM authors
WHERE id = $1;