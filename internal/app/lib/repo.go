package lib

import (
	"fmt"
	"os"
	// "time"

	"github.com/go-facegit/facegit-rpc/internal/conf"
	"github.com/go-facegit/facegit-rpc/internal/tools"
	"github.com/gogs/git-module"
)

type Repo struct {
}

type RepoCreateArgs struct {
	UserOrOrg   string
	ProjectName string
}

//Init Git Repo Project
func (rp *Repo) Create(args *RepoCreateArgs, reply *bool) error {
	*reply = false
	rootPath := conf.Repo.RootPath
	repoPath := fmt.Sprintf("%s/%s/%s.git", rootPath, args.UserOrOrg, args.ProjectName)
	if !tools.IsExist(repoPath) {
		if err := git.Init(repoPath, git.InitOptions{Bare: true}); err != nil {
			return fmt.Errorf("init repository: %v", err)
		}
		*reply = true
	}
	return nil
}

func (rp *Repo) Delete(args *RepoCreateArgs, reply *bool) error {
	*reply = false
	rootPath := conf.Repo.RootPath
	repoPath := fmt.Sprintf("%s/%s/%s.git", rootPath, args.UserOrOrg, args.ProjectName)
	if err := os.RemoveAll(repoPath); err != nil {
		desc := fmt.Sprintf("[%s]: %v", repoPath, err)
		return fmt.Errorf("delete repository: %s", desc)
	}
	*reply = true
	return nil
}
