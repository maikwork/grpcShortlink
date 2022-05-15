package repository

import "database/sql"

type SQLQuery interface {
	Get(shortLink string) string
	Create(link string)
}

type SQL struct {
	conn *sql.DB
}

func NewSQL(c *sql.DB) *SQL {
	return &SQL{
		conn: c,
	}
}
