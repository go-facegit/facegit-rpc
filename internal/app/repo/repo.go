package repo

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/go-facegit/facegit-rpc/internal/conf"
	"github.com/go-facegit/facegit-rpc/internal/cryptoutil"
	"github.com/go-facegit/facegit-rpc/internal/tools"
	"github.com/go-facegit/facegit-rpc/pb"
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

type RepoNewestInfo struct {
	AuthorName string
	Name       string
	Id         string
	Message    string
	When       time.Time
}

type RepoRetList struct {
	Newest RepoNewestInfo
	List   []*git.EntryCommitInfo
}

const (
	ENV_AUTH_USER_ID           = "FG_AUTH_USER_ID"
	ENV_AUTH_USER_NAME         = "FG_AUTH_USER_NAME"
	ENV_AUTH_USER_EMAIL        = "FG_AUTH_USER_EMAIL"
	ENV_REPO_OWNER_NAME        = "FG_REPO_OWNER_NAME"
	ENV_REPO_OWNER_SALT_MD5    = "FG_REPO_OWNER_SALT_MD5"
	ENV_REPO_ID                = "FG_REPO_ID"
	ENV_REPO_NAME              = "FG_REPO_NAME"
	ENV_REPO_CUSTOM_HOOKS_PATH = "FG_REPO_CUSTOM_HOOKS_PATH"
)

type ComposeHookEnvsOptions struct {
	AuthUser  string
	OwnerName string
	OwnerSalt string
	RepoID    int64
	RepoName  string
	RepoPath  string
}

func ComposeHookEnvs(opts ComposeHookEnvsOptions) []string {
	envs := []string{
		"SSH_ORIGINAL_COMMAND=1",
		ENV_AUTH_USER_ID + "=" + opts.AuthUser,
		ENV_AUTH_USER_NAME + "=" + opts.AuthUser,
		ENV_AUTH_USER_EMAIL + "=midoks@163.com",
		ENV_REPO_OWNER_NAME + "=" + opts.OwnerName,
		ENV_REPO_OWNER_SALT_MD5 + "=" + cryptoutil.MD5("TkNOFygiTd"),
		ENV_REPO_ID + "=" + tools.ToStr(opts.RepoID),
		ENV_REPO_NAME + "=" + opts.RepoName,
		ENV_REPO_CUSTOM_HOOKS_PATH + "=" + filepath.Join(opts.RepoPath, "custom_hooks"),
	}
	return envs
}

// discardLocalRepoBranchChanges discards local commits/changes of
// given branch to make sure it is even to remote branch.
func discardLocalRepoBranchChanges(localPath, branch string) error {
	if !tools.IsExist(localPath) {
		return nil
	}

	// No need to check if nothing in the repository.
	if !git.RepoHasBranch(localPath, branch) {
		return nil
	}

	rev := "origin/" + branch
	if err := git.RepoReset(localPath, rev, git.ResetOptions{Hard: true}); err != nil {
		return fmt.Errorf("reset [revision: %s]: %v", rev, err)
	}
	return nil
}

// UpdateLocalCopy fetches latest changes of given branch from repoPath to localPath.
// It creates a new clone if local copy does not exist, but does not checks out to a
// specific branch if the local copy belongs to a wiki.
// For existing local copy, it checks out to target branch by default, and safe to
// assume subsequent operations are against target branch when caller has confidence
// about no race condition.
func UpdateLocalCopyBranch(repoPath, localPath, branch string, isWiki bool) (err error) {
	if !tools.IsExist(localPath) {
		// Checkout to a specific branch fails when wiki is an empty repository.
		if isWiki {
			branch = ""
		}
		if err = git.Clone(repoPath, localPath, git.CloneOptions{
			Branch:  branch,
			Timeout: time.Duration(300) * time.Second,
		}); err != nil {
			return fmt.Errorf("git clone [branch: %s]: %v", branch, err)
		}
		return nil
	}
	fmt.Println("gitRepo:", localPath)
	gitRepo, err := git.Open(localPath)
	if err != nil {
		return fmt.Errorf("open repository: %v", err)
	}

	if err = gitRepo.Fetch(git.FetchOptions{
		Prune: true,
	}); err != nil {
		return fmt.Errorf("fetch: %v", err)
	}

	if err = gitRepo.Checkout(branch); err != nil {
		return fmt.Errorf("checkout [branch: %s]: %v", branch, err)
	}

	// Reset to align with remote in case of force push.
	rev := "origin/" + branch
	if err = gitRepo.Reset(rev, git.ResetOptions{
		Hard: true,
	}); err != nil {
		return fmt.Errorf("reset [revision: %s]: %v", rev, err)
	}
	return nil
}

/////////////////////
/// REPO OP --------]
/////////////////////

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

func RepoList(UserOrOrg, ProjectName, TreePath string) (error, RepoRetList) {

	var err error
	repo := Repo{}
	repoList := RepoRetList{}
	rootPath := conf.Repo.RootPath
	projPath := fmt.Sprintf("%s/%s", UserOrOrg, ProjectName)
	repoPath := fmt.Sprintf("%s/%s.git", rootPath, projPath)

	repo.GitRepo, err = git.Open(repoPath)
	if err != nil {
		return fmt.Errorf("repository[%s] error: %s", projPath, err), repoList
	}

	repo.Commit, err = repo.GitRepo.BranchCommit("master")
	if err != nil {
		return fmt.Errorf("repository[%s] commit error: %s", projPath, err), repoList
	}

	entry, err := repo.Commit.TreeEntry(TreePath)
	if err != nil {
		return fmt.Errorf("repository[%s] get tree entry error: %s", projPath, err), repoList
	}

	tree, err := repo.Commit.Subtree(TreePath)
	if err != nil {
		return fmt.Errorf("repository[%s] subtree error: %s", projPath, err), repoList
	}

	entries, err := tree.Entries()
	if err != nil {
		return fmt.Errorf("%s: list entries", err), repoList
	}
	entries.Sort()

	fList, err := entries.CommitsInfo(repo.Commit, git.CommitsInfoOptions{
		Path:           TreePath,
		MaxConcurrency: 5,
		Timeout:        5 * time.Minute,
	})

	latestCommit := repo.Commit
	if len(TreePath) > 0 {
		latestCommit, err = repo.Commit.CommitByPath(git.CommitByRevisionOptions{Path: TreePath})
		if err != nil {
			return fmt.Errorf("get commit by path: %s", err), repoList
		}
	}

	info := RepoNewestInfo{
		AuthorName: latestCommit.Author.Name,
		Id:         fmt.Sprintf("%s", latestCommit.ID),
		Message:    latestCommit.Message,
		Name:       entry.Name(),
		When:       latestCommit.Author.When,
	}

	repoList = RepoRetList{
		Newest: info,
		List:   fList,
	}

	// for k, v := range fList {
	// 	fmt.Println(k, v.Entry.Name(), v.Commit.ID, v.Entry.Type())
	// }

	return nil, repoList
}

func RepoEditorFile(UserOrOrg, ProjectName string, opts *pb.ReqUpdateOptions) error {
	var err error

	rootPath := conf.Repo.RootPath
	projPath := fmt.Sprintf("%s/%s", UserOrOrg, ProjectName)
	repoPath := fmt.Sprintf("%s/%s.git", rootPath, projPath)

	tmpPath := path.Join(path.Dir(rootPath), "tmp")
	uPath := path.Join(tmpPath, UserOrOrg)

	oldFilePath := path.Join(uPath, opts.OldTreeName)
	filePath := path.Join(uPath, opts.NewTreeName)

	if err = discardLocalRepoBranchChanges(uPath, opts.OldBranch); err != nil {
		return fmt.Errorf("discard local repo branch[%s] changes: %v", opts.OldBranch, err)
	} else if err = UpdateLocalCopyBranch(repoPath, uPath, opts.OldBranch, false); err != nil {
		return fmt.Errorf("update local copy branch[%s]: %v", opts.OldBranch, err)
	}

	if err := os.MkdirAll(path.Dir(filePath), os.ModePerm); err != nil {
		return err
	}

	// If it's meant to be a new file, make sure it doesn't exist.
	if opts.IsNewFile {
		if tools.IsExist(filePath) {
			return fmt.Errorf("File Already Exist[%s]!", filePath)
		}
	}

	if tools.IsFile(oldFilePath) && opts.OldTreeName != opts.NewTreeName {
		if err = git.RepoMove(uPath, opts.OldTreeName, opts.NewTreeName); err != nil {
			return fmt.Errorf("git mv %q %q: %v", opts.OldTreeName, opts.NewTreeName, err)
		}
	}

	if err := ioutil.WriteFile(filePath, []byte(opts.Content), 0666); err != nil {
		return fmt.Errorf("write file: %v", err)
	}

	gitSig := &git.Signature{
		Name:  UserOrOrg,
		Email: "midoks@163.com",
		When:  time.Now(),
	}

	if err := git.RepoAdd(uPath, git.AddOptions{All: true}); err != nil {
		return fmt.Errorf("git add --all: %v", err)
	} else if err := git.RepoCommit(uPath, gitSig, opts.Message); err != nil {
		return fmt.Errorf("commit changes on %q: %v", uPath, err)
	}

	envs := ComposeHookEnvs(ComposeHookEnvsOptions{
		AuthUser:  UserOrOrg,
		OwnerName: UserOrOrg,
		// OwnerSalt: repo.MustOwner().Salt,
		RepoID:   1,
		RepoName: UserOrOrg,
		RepoPath: repoPath,
	})

	// fmt.Println("envs:", envs)

	if err := git.RepoPush(uPath, "origin", opts.NewBranch, git.PushOptions{Envs: envs}); err != nil {
		return fmt.Errorf("git push origin %s: %v", opts.NewBranch, err)
	}

	// fmt.Println("RepoEditor:", err)

	return nil
}
