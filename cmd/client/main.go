package main

import (
	"context"
	"log"

	"github.com/maikwork/grpcShortlink/internal/api"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("%v", err)
	}

	client := api.NewShortLinkServiceClient(conn)

	req := &api.RequestCreateLink{Link: "asadfdsa"}
	client.CreateLink(context.Background(), req)
}
