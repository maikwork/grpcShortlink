package main

import (
	"context"
	"log"

	"github.com/maikwork/grpcShortlink/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("%v", err)
	}

	client := pb.NewShortLinkServiceClient(conn)

	req := &pb.RequestCreateLink{Long: "asadfdsa"}
	client.CreateLink(context.Background(), req)
}
