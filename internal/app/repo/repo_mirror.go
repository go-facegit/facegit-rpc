package repo

import (
	// "fmt"
	"strings"
	"time"

	"github.com/gogs/git-module"
	// "github.com/go-facegit/facegit-rpc/internal/conf"
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
