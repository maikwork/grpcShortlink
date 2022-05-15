package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/maikwork/grpcShortlink/internal/model"
	"github.com/maikwork/grpcShortlink/internal/repository"
	"github.com/maikwork/grpcShortlink/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	rep  repository.SQLQuery
	conf *model.Config
}

func NewServer(r repository.SQLQuery, cfg *model.Config) *Server {
	return &Server{rep: r, conf: cfg}
}

func (s *Server) CreateLink(ctx context.Context, req *pb.RequestCreateLink) (*pb.ResponseCreateLink, error) {
	s.rep.Create(req.Long)
	return nil, nil
}

func (s *Server) GetLink(ctx context.Context, req *pb.RequestGetLink) (*pb.ResponseGetLink, error) {
	link := s.rep.Get(req.Short)
	return &pb.ResponseGetLink{Long: link}, nil
}

func runRest(srv *Server) {
	pg := srv.conf.SettingServ.GRPC.Port
	pr := srv.conf.SettingServ.REST.Port
	port := fmt.Sprintf(":%v", pr)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	host := fmt.Sprintf("localhost:%v", pg)
	err := pb.RegisterShortLinkServiceHandlerFromEndpoint(ctx, mux, host, opts)
	if err != nil {
		panic(err)
	}

	log.Printf("Server listening at %v...", pr)

	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatal(err)
	}
}

func runGrpc(srv *Server) {
	pg := srv.conf.SettingServ.GRPC.Port
	port := fmt.Sprintf(":%v", pg)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterShortLinkServiceServer(s, srv)

	log.Printf("Server listening at %v...", pg)

	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

func (s *Server) Run() {
	go runRest(s)
	runGrpc(s)
}
