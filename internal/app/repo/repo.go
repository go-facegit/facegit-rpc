package repo

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
	Id     string
	UpTime string
	Info   int
}

type RepoRetList struct {
	Info RepoBaseInfo
	// List []git.EntryCommitInfo
}

//Init Git Repo Project
func RepoCreate(UserOrOrg, ProjectName string) (bool, error) {
	rootPath := conf.Repo.RootPath
	repoPath := fmt.Sprintf("%s/%s/%s.git", rootPath, UserOrOrg, ProjectName)
	if !tools.IsExist(repoPath) {
		if err := git.Init(repoPath, git.InitOptions{Bare: true}); err != nil {
			return false, fmt.Errorf("init repository: %v", err)
		}
	}
	return true, nil
}

// Delete Repo
func RepoDelete(UserOrOrg, ProjectName string) (bool, error) {
	rootPath := conf.Repo.RootPath
	repoPath := fmt.Sprintf("%s/%s/%s.git", rootPath, UserOrOrg, ProjectName)
	if err := os.RemoveAll(repoPath); err != nil {
		desc := fmt.Sprintf("[%s]: %v", repoPath, err)
		return false, fmt.Errorf("delete repository error: %s", desc)
	}
	return true, nil
}

func RepoList(UserOrOrg, ProjectName, TreePath string) (bool, error) {

	var err error
	repo := Repo{}
	rootPath := conf.Repo.RootPath
	projPath := fmt.Sprintf("%s/%s", UserOrOrg, ProjectName)
	repoPath := fmt.Sprintf("%s/%s.git", rootPath, projPath)

	repo.GitRepo, err = git.Open(repoPath)
	if err != nil {
		return false, fmt.Errorf("repository[%s] error: %s", projPath, err)
	}

	repo.Commit, err = repo.GitRepo.BranchCommit("master")
	if err != nil {
		return false, fmt.Errorf("repository[%s] commit error: %s", projPath, err)
	}

	entry, err := repo.Commit.TreeEntry(TreePath)
	if err != nil {
		return false, fmt.Errorf("repository[%s] get tree entry error: %s", projPath, err)
	}

	tree, err := repo.Commit.Subtree(TreePath)
	if err != nil {
		return false, fmt.Errorf("repository[%s] subtree error: %s", projPath, err)
	}

	entries, err := tree.Entries()
	if err != nil {
		return false, fmt.Errorf("%s: list entries", err)
	}
	entries.Sort()

	f, err := entries.CommitsInfo(repo.Commit, git.CommitsInfoOptions{
		Path:           TreePath,
		MaxConcurrency: 5,
		Timeout:        5 * time.Minute,
	})

	latestCommit := repo.Commit
	if len(TreePath) > 0 {
		latestCommit, err = repo.Commit.CommitByPath(git.CommitByRevisionOptions{Path: TreePath})
		if err != nil {
			return false, fmt.Errorf("get commit by path: %s", err)
		}
	}

	id := fmt.Sprintf("%s", latestCommit.ID)

	fmt.Println("id:", id)
	info := RepoBaseInfo{
		Id:   "mmmm",
		Info: 2,
	}

	reply := &RepoRetList{
		Info: info,
		// List: f,
	}
	fmt.Println(reply)

	for k, v := range f {
		fmt.Println(k, v.Entry.Name(), v.Commit.ID, v.Entry.Type())
	}

	fmt.Println(latestCommit.ID.String)
	fmt.Println("entry tree:", entry.Name(), entry.ID(), entry.Type())
	fmt.Println("subtree error:", f, err)
	return true, nil
}
