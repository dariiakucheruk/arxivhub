-- name: CreatePaper :one
INSERT INTO papers (
    arxiv_id,
    title,
    abstract,
    authors,
    short_authors
) VALUES ($1, $2, $3, $4, $5
) RETURNING *;

-- name: GetPaper :one
SELECT * FROM papers
WHERE arxiv_id = $1
LIMIT 1;

-- name: GetPapers :many
SELECT * FROM papers
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeletePaper :exec
DELETE FROM papers
WHERE arxiv_id = $1;

-- name: SavePaperForUser :one
INSERT INTO saved_papers(
     user_id,
     paper_id
) VALUES ($1, $2
) RETURNING *;

-- name: GetSavedPapersForUser :many
SELECT * FROM saved_papers
WHERE user_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: DeleteSavedPaper :exec
DELETE from saved_papers
WHERE user_id = $1 AND paper_id = $2;

