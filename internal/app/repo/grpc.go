package repo

import (
	// "log"

	"github.com/go-facegit/facegit-rpc/pb"
	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) S(ctx context.Context, message *pb.Message) (*pb.Message, error) {
	return &pb.Message{Body: "hh"}, nil
}

func (s *Server) List(ctx context.Context, message *pb.ReqList) (*pb.Message, error) {
	return &pb.Message{Body: "hh"}, nil
}
