// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: getOneWord.sql

package database

import (
	"context"
)

const getOneWord = `-- name: GetOneWord :one
SELECT id, word, origin, fullword, definition, etymology, type, sentence, created_at, updated_at FROM dictionary WHERE word = $1
`

func (q *Queries) GetOneWord(ctx context.Context, word string) (Dictionary, error) {
	row := q.db.QueryRowContext(ctx, getOneWord, word)
	var i Dictionary
	err := row.Scan(
		&i.ID,
		&i.Word,
		&i.Origin,
		&i.Fullword,
		&i.Definition,
		&i.Etymology,
		&i.Type,
		&i.Sentence,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}