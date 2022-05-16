package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/maikwork/grpcGeteway/internal/model"
)

func Connect(d model.DB) (*sql.DB, error) {
	db, err := sql.Open(d.Name, d.URL)
	if err != nil {
		log.Printf("Don't open a connection to the sql")
		return nil, err
	}

	return db, nil
}
