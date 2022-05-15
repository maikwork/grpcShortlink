package main

import (
	"flag"
	"log"

	"github.com/maikwork/grpcShortlink/internal/database"
	"github.com/maikwork/grpcShortlink/internal/helper"
	"github.com/maikwork/grpcShortlink/internal/repository"
	"github.com/maikwork/grpcShortlink/internal/server"
)

func main() {
	path := flag.String("d", "./cmd/server/config.yaml", "path to config")
	flag.Parse()

	cfg, err := helper.ReadConfig(*path)
	if err != nil {
		log.Fatalf("%v", err)
	}

	log.Printf("config: %v", cfg)

	db, err := database.Connect(cfg.DB)
	if err != nil {
		log.Fatalf("Don't work sql: %v", err)
	}
	defer db.Close()

	rep := repository.NewSQL(db)
	s := server.NewServer(rep, cfg)

	s.Run()
}
