package repo

import (
	"fmt"
	"strings"
	"time"

	"github.com/gogs/git-module"

	"github.com/go-facegit/facegit-rpc/internal/conf"
)

var commonWikiURLSuffixes = []string{".wiki.git", ".git/wiki"}

// wikiRemoteURL returns accessible repository URL for wiki if exists.
// Otherwise, it returns an empty string.
func wikiRemoteURL(remote string) string {
	remote = strings.TrimSuffix(remote, ".git")
	for _, suffix := range commonWikiURLSuffixes {
		wikiURL := remote + suffix
		if git.IsURLAccessible(time.Minute, wikiURL) {
			return wikiURL
		}
	}
	return ""
}

func RepoMirror(RemoteAddr, UserOrOrg, ProjectName string) error {
	rootPath := conf.Repo.RootPath
	repoPath := fmt.Sprintf("%s/%s/%s.git", rootPath, UserOrOrg, ProjectName)
	wikiPath := fmt.Sprintf("%s/%s/%s.wiki.git", rootPath, UserOrOrg, ProjectName)

	if err := git.Clone(RemoteAddr, repoPath, git.CloneOptions{
		Mirror:  true,
		Quiet:   true,
		Timeout: 1000,
	}); err != nil {
		return fmt.Errorf("clone: %v", err)
	}

	wikiRemotePath := wikiRemoteURL(RemoteAddr)
	if len(wikiRemotePath) > 0 {
		if err := git.Clone(wikiRemotePath, wikiPath, git.CloneOptions{
			Mirror:  true,
			Quiet:   true,
			Timeout: 1000,
		}); err != nil {
			return fmt.Errorf("Failed to clone wiki: %v", err)
		}
	}
	return nil
}
