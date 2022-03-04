package lib

import (
	"fmt"
	"os"
	"time"

	"github.com/go-facegit/facegit-rpc/internal/conf"
	"github.com/go-facegit/facegit-rpc/internal/tools"
	"github.com/gogs/git-module"
)

type Repo struct {
	Commit  *git.Commit
	Tag     *git.Tag
	GitRepo *git.Repository
}

type RepoCreateArgs struct {
	UserOrOrg   string
	ProjectName string
	TreePath    string
}

type RepoBaseInfo struct {
	UserOrOrg   string
	ProjectName string
	TreePath    string
}

type RepoBaseList struct {
	UserOrOrg   string
	ProjectName string
	TreePath    string
}

type RepoList struct {
	Info RepoBaseInfo
	List RepoBaseList
}

//Init Git Repo Project
func (repo *Repo) Create(args *RepoCreateArgs, reply *bool) error {
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

// Delete Repo
func (repo *Repo) Delete(args *RepoCreateArgs, reply *bool) error {
	*reply = false
	rootPath := conf.Repo.RootPath
	repoPath := fmt.Sprintf("%s/%s/%s.git", rootPath, args.UserOrOrg, args.ProjectName)
	if err := os.RemoveAll(repoPath); err != nil {
		desc := fmt.Sprintf("[%s]: %v", repoPath, err)
		return fmt.Errorf("delete repository error: %s", desc)
	}
	*reply = true
	return nil
}

// Repo List Rpc
func (repo *Repo) List(args *RepoCreateArgs, reply *bool) error {
	*reply = false
	var err error

	rootPath := conf.Repo.RootPath
	projPath := fmt.Sprintf("%s/%s", args.UserOrOrg, args.ProjectName)

	repoPath := fmt.Sprintf("%s/%s.git", rootPath, projPath)

	repo.GitRepo, err = git.Open(repoPath)
	if err != nil {
		return fmt.Errorf("repository[%s] error: %s", projPath, err)
	}

	repo.Commit, err = repo.GitRepo.BranchCommit("master")
	if err != nil {
		return fmt.Errorf("repository[%s] commit error: %s", projPath, err)
	}

	entry, err := repo.Commit.TreeEntry(args.TreePath)
	if err != nil {
		return fmt.Errorf("repository[%s] get tree entry error: %s", projPath, err)
	}

	tree, err := repo.Commit.Subtree(args.TreePath)
	if err != nil {
		return fmt.Errorf("repository[%s] subtree error: %s", projPath, err)
	}

	entries, err := tree.Entries()
	if err != nil {
		return fmt.Errorf("%s: list entries", err)
	}
	entries.Sort()

	f, err := entries.CommitsInfo(repo.Commit, git.CommitsInfoOptions{
		Path:           args.TreePath,
		MaxConcurrency: 5,
		Timeout:        5 * time.Minute,
	})

	latestCommit := repo.Commit
	if len(args.TreePath) > 0 {
		latestCommit, err = repo.Commit.CommitByPath(git.CommitByRevisionOptions{Path: args.TreePath})
		if err != nil {
			return fmt.Errorf("get commit by path: %s", err)
		}
	}

	fmt.Println(latestCommit.ID.String())

	for k, v := range f {
		fmt.Println(k, v.Entry.Name(), v.Commit.ID, v.Entry.Type())
	}

	fmt.Println("entry tree:", entry.Name(), entry.ID(), entry.Type())
	fmt.Println("subtree error:", f, err)

	return nil
}
