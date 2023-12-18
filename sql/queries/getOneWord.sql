-- name: GetOneWord :one
SELECT * FROM dictionary WHERE word = $1;