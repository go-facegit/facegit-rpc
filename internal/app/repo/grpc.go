package repo

import (
	// "log"

	"github.com/go-facegit/facegit-rpc/pb"
	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) Create(ctx context.Context, message *pb.ReqBase) (*pb.RespBool, error) {
	tof, err := RepoCreate(message.UserOrOrg, message.ProjectName)
	return &pb.RespBool{TrueOrFalse: tof}, err
}

func (s *Server) Delete(ctx context.Context, message *pb.ReqBase) (*pb.RespBool, error) {
	tof, err := RepoDelete(message.UserOrOrg, message.ProjectName)
	return &pb.RespBool{TrueOrFalse: tof}, err
}

func (s *Server) List(ctx context.Context, message *pb.ReqList) (*pb.RespList, error) {
	err, data := RepoList(message.UserOrOrg, message.ProjectName, message.TreePath)

	if err != nil {
		return &pb.RespList{}, err
	}
	return data, nil
}

func (s *Server) Editor(ctx context.Context, opts *pb.ReqUpdateOptions) (*pb.RespBool, error) {
	err := RepoEditorFile(opts.UserOrOrg, opts.ProjectName, opts)
	// fmt.Println("Editor:", err)
	if err != nil {
		return &pb.RespBool{TrueOrFalse: true}, nil
	}
	return &pb.RespBool{TrueOrFalse: false}, err
}

// ------ mirror ---------

func (s *Server) CreateMirror(ctx context.Context, opts *pb.ReqMirror) (*pb.RespBool, error) {
	err := RepoCreateMirror(opts.RemoteAddr, opts.UserOrOrg, opts.ProjectName)
	if err == nil {
		return &pb.RespBool{TrueOrFalse: true}, nil
	}
	return &pb.RespBool{TrueOrFalse: false}, nil
}
