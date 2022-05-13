package server

import (
	"context"

	"github.com/maikwork/grpcShortlink/internal/api"
	"github.com/maikwork/grpcShortlink/internal/repository"
)

type Server struct {
	rep repository.SQLQuery
}

func NewServer(r repository.SQLQuery) *Server {
	return &Server{rep: r}
}

func (s *Server) CreateLink(ctx context.Context, req *api.RequestCreateLink) (*api.ResponseCreateLink, error) {
	s.rep.Create(req.Link)
	return nil, nil
}

func (s *Server) GetLink(ctx context.Context, req *api.RequestGetLink) (*api.ResponseGetLink, error) {
	link := s.rep.Get(req.ShortLink)
	return &api.ResponseGetLink{Link: link}, nil
}
