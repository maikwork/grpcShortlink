package repository

import (
	"fmt"
	"log"

	"github.com/maikwork/grpcGeteway/internal/domain"
)

func (s *SQL) Create(link string) {
	fmt.Printf("Create link")
	db := s.conn

	d, err := domain.Decode(link)
	if err != nil {
		log.Printf("%v", err)
	}

	query := fmt.Sprintf("INSERT INTO link(link, short_link) VALUES('%v', '%v')", link, d)

	_, err = db.Query(query)
	if err != nil {
		log.Printf("%v", err)
	}
}

func (s *SQL) Get(short string) string {
	fmt.Printf("Get link")
	var result string
	db := s.conn

	query := fmt.Sprintf("SELECT link FROM links WHERE short_link='%v'", short)

	d, err := db.Query(query)
	if err != nil {
		fmt.Printf("%v", err)
	}
	defer d.Close()

	d.Scan(&result)

	return result
}
