package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Repository interface {
	Querier
}

type PSQLRepository struct {
	*Queries
	db *sql.DB
}

func NewPSQLRepository(DNS string) (*PSQLRepository, error) {
	db, err := sql.Open("postgres", DNS)
	if err != nil {
		return nil, fmt.Errorf("connection to db error: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("check conn to db error : %w", err)
	}

	return &PSQLRepository{db: db, Queries: New(db)}, nil
}
