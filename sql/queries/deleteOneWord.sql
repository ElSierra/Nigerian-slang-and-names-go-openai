-- name: DeleteOneWord :one
SELECT * FROM dictionary WHERE word = $1;