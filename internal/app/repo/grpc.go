package repo

import (
	// "log"
	"fmt"

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

	list := data.List
	var retList []*pb.RespStructList
	for _, v := range list {

		t := &pb.RespStructList{
			Name:     v.Entry.Name(),
			CommitId: fmt.Sprintf("%s", v.Commit.ID),
			Type:     fmt.Sprintf("%s", v.Entry.Type()),
			When:     v.Commit.Committer.When.String(),
		}

		retList = append(retList, t)
	}

	return &pb.RespList{
		Newest: &pb.RespStructNewest{
			Message:    data.Newest.Message,
			CommitId:   data.Newest.CommitId,
			AuthorName: data.Newest.AuthorName,
			When:       data.Newest.When.String(),
			BranchName: data.Newest.BranchName,
		},
		List: retList,
	}, nil
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
