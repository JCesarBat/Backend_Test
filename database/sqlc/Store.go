package database

import "database/sql"

type SQLStore struct {
	conn *sql.DB
	*Queries
}

// NewStore creates a new store
func NewStore(conn *sql.DB) *SQLStore {
	return &SQLStore{
		conn:    conn,
		Queries: New(conn),
	}
}
