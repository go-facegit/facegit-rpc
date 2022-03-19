package repo

import (
	// "log"
	"fmt"

	"github.com/go-facegit/facegit-rpc/pb"
	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) Delete(ctx context.Context, message *pb.ReqBase) (*pb.RespBool, error) {
	tof, err := RepoDelete(message.UserOrOrg, message.ProjectName)
	return &pb.RespBool{TrueOrFalse: tof}, err
}

func (s *Server) Create(ctx context.Context, message *pb.ReqBase) (*pb.RespBool, error) {
	tof, err := RepoCreate(message.UserOrOrg, message.ProjectName)
	return &pb.RespBool{TrueOrFalse: tof}, err
}

func (s *Server) List(ctx context.Context, message *pb.ReqList) (*pb.Message, error) {

	b, err := RepoList(message.UserOrOrg, message.ProjectName, message.TreePath)
	fmt.Println(b, err)
	return &pb.Message{Body: "hh"}, nil
}
