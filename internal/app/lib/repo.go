package lib

import (
	"fmt"

	"github.com/go-facegit/facegit-rpc/internal/conf"
	"github.com/go-facegit/facegit-rpc/internal/tools"
	"github.com/gogs/git-module"
)

type Repo struct {
	s string
}

type RepoCreateArgs struct {
	UserName    string
	ProjectName string
}

func (rp *Repo) Create(args *RepoCreateArgs, reply *int) error {
	fmt.Println("repo name", args.UserName, args.ProjectName)

	rootPath := conf.Repo.RootPath

	if tools.IsExist(rootPath) {

		repoPath := fmt.Sprintf("%s/%s/%s", rootPath, args.UserName, args.ProjectName)

		if err := git.Init(repoPath, git.InitOptions{Bare: true}); err != nil {
			return fmt.Errorf("init repository: %v", err)
		}
	}

	*reply = 1
	return nil
}
