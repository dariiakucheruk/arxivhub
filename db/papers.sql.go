// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: papers.sql

package db

import (
	"context"
	"database/sql"
)

const createPaper = `-- name: CreatePaper :one
INSERT INTO papers (
    arxiv_id,
    title,
    abstract,
    authors,
    short_authors
) VALUES ($1, $2, $3, $4, $5
) RETURNING arxiv_id, title, abstract, authors, short_authors, date
`

type CreatePaperParams struct {
	ArxivID      string `json:"arxiv_id"`
	Title        string `json:"title"`
	Abstract     string `json:"abstract"`
	Authors      string `json:"authors"`
	ShortAuthors string `json:"short_authors"`
}

func (q *Queries) CreatePaper(ctx context.Context, arg CreatePaperParams) (Paper, error) {
	row := q.db.QueryRowContext(ctx, createPaper,
		arg.ArxivID,
		arg.Title,
		arg.Abstract,
		arg.Authors,
		arg.ShortAuthors,
	)
	var i Paper
	err := row.Scan(
		&i.ArxivID,
		&i.Title,
		&i.Abstract,
		&i.Authors,
		&i.ShortAuthors,
		&i.Date,
	)
	return i, err
}

const deletePaper = `-- name: DeletePaper :exec
DELETE FROM papers
WHERE arxiv_id = $1
`

func (q *Queries) DeletePaper(ctx context.Context, arxivID string) error {
	_, err := q.db.ExecContext(ctx, deletePaper, arxivID)
	return err
}

const deleteSavedPaper = `-- name: DeleteSavedPaper :exec
DELETE from saved_papers
WHERE user_id = $1 AND paper_id = $2
`

type DeleteSavedPaperParams struct {
	UserID  sql.NullInt64  `json:"user_id"`
	PaperID sql.NullString `json:"paper_id"`
}

func (q *Queries) DeleteSavedPaper(ctx context.Context, arg DeleteSavedPaperParams) error {
	_, err := q.db.ExecContext(ctx, deleteSavedPaper, arg.UserID, arg.PaperID)
	return err
}

const getPaper = `-- name: GetPaper :one
SELECT arxiv_id, title, abstract, authors, short_authors, date FROM papers
WHERE arxiv_id = $1
LIMIT 1
`

func (q *Queries) GetPaper(ctx context.Context, arxivID string) (Paper, error) {
	row := q.db.QueryRowContext(ctx, getPaper, arxivID)
	var i Paper
	err := row.Scan(
		&i.ArxivID,
		&i.Title,
		&i.Abstract,
		&i.Authors,
		&i.ShortAuthors,
		&i.Date,
	)
	return i, err
}

const getPapers = `-- name: GetPapers :many
SELECT arxiv_id, title, abstract, authors, short_authors, date FROM papers
ORDER BY id
LIMIT $1
OFFSET $2
`

type GetPapersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetPapers(ctx context.Context, arg GetPapersParams) ([]Paper, error) {
	rows, err := q.db.QueryContext(ctx, getPapers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Paper{}
	for rows.Next() {
		var i Paper
		if err := rows.Scan(
			&i.ArxivID,
			&i.Title,
			&i.Abstract,
			&i.Authors,
			&i.ShortAuthors,
			&i.Date,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getSavedPapersForUser = `-- name: GetSavedPapersForUser :many
SELECT id, user_id, paper_id FROM saved_papers
WHERE user_id = $1
ORDER BY id
LIMIT $2
OFFSET $3
`

type GetSavedPapersForUserParams struct {
	UserID sql.NullInt64 `json:"user_id"`
	Limit  int32         `json:"limit"`
	Offset int32         `json:"offset"`
}

func (q *Queries) GetSavedPapersForUser(ctx context.Context, arg GetSavedPapersForUserParams) ([]SavedPaper, error) {
	rows, err := q.db.QueryContext(ctx, getSavedPapersForUser, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []SavedPaper{}
	for rows.Next() {
		var i SavedPaper
		if err := rows.Scan(&i.ID, &i.UserID, &i.PaperID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const savePaperForUser = `-- name: SavePaperForUser :one
INSERT INTO saved_papers(
     user_id,
     paper_id
) VALUES ($1, $2
) RETURNING id, user_id, paper_id
`

type SavePaperForUserParams struct {
	UserID  sql.NullInt64  `json:"user_id"`
	PaperID sql.NullString `json:"paper_id"`
}

func (q *Queries) SavePaperForUser(ctx context.Context, arg SavePaperForUserParams) (SavedPaper, error) {
	row := q.db.QueryRowContext(ctx, savePaperForUser, arg.UserID, arg.PaperID)
	var i SavedPaper
	err := row.Scan(&i.ID, &i.UserID, &i.PaperID)
	return i, err
}
