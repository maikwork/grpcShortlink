package main

import (
	"flag"
	"log"
	"net"

	"github.com/maikwork/grpcShortlink/internal/api"
	"github.com/maikwork/grpcShortlink/internal/database"
	"github.com/maikwork/grpcShortlink/internal/helper"
	"github.com/maikwork/grpcShortlink/internal/repository"
	"github.com/maikwork/grpcShortlink/internal/server"
	"google.golang.org/grpc"
)

func main() {
	path := flag.String("d", "./cmd/run/config.yaml", "path to config")
	flag.Parse()

	cfg, err := helper.ReadConfig(*path)
	if err != nil {
		log.Fatal("Don't read config")
	}

	log.Printf("config: %v", cfg)

	db, err := database.Connect(cfg.DB)
	if err != nil {
		log.Fatalf("Don't work sql: %v", err)
	}
	defer db.Close()

	l, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("Don't connect: %v", err)
	}
	defer l.Close()

	rep := repository.NewSQL(db)
	s := server.NewServer(rep)

	grpcServer := grpc.NewServer()
	api.RegisterShortLinkServiceServer(grpcServer, s)
	grpcServer.Serve(l)
}
