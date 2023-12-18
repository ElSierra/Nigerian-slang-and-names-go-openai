-- name: CreateWord :one
INSERT INTO dictionary (word, origin, fullWord, definition, etymology, type, sentence) 
VALUES ($1, $2, $3, $4, $5, $6, $7) 
RETURNING *;