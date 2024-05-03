package database

import "database/sql"

type SQLStore struct {
	conn *sql.DB
	*Queries
}

func NewStore(conn *sql.DB) *SQLStore {
	return &SQLStore{
		conn:    conn,
		Queries: New(conn),
	}
}
