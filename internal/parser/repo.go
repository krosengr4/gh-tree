package parser

import (
	"fmt"
	"strings"
)

type RepoInfo struct {
	Owner    string
	RepoName string
}

func ParseRepo(input string) (*RepoInfo, error) {

	parts := strings.Split(input, "/")

	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid format. command must be ghtree owner/repoName")
	}

	owner := strings.TrimSpace(parts[0])
	repoName := strings.TrimSpace(parts[1])

	if owner == "" {
		return nil, fmt.Errorf("empty owner in owner/repoName")
	}
	if repoName == "" {
		return nil, fmt.Errorf("empty repo name in owner/repoName")
	}

	return &RepoInfo{
		Owner:    owner,
		RepoName: repoName,
	}, nil
}
