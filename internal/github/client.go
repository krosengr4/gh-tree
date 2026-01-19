package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v57/github"

	"golang.org/x/oauth2"
)

type Client struct {
	client *github.Client
	ctx    context.Context
}

type TreeNode struct {
	Path string
	Type string
	Size int64
}

func NewClient(token string) *Client {
	ctx := context.Background()
	var client *github.Client

	if token != "" {
		// Create authenticated client
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)

		tc := oauth2.NewClient(ctx, ts)
		client = github.NewClient(tc)
	} else {
		client = github.NewClient(nil)
	}

	return &Client{
		client: client,
		ctx:    ctx,
	}
}

func (c *Client) GetTree(owner, repo, branch string, recursive bool) ([]TreeNode, error) {
	// Get reference for the branch
	ref, _, err := c.client.Git.GetRef(c.ctx, owner, repo, "refs/heads/"+branch)
	if err != nil {
		return nil, fmt.Errorf("failed to get branch reference: %w", err)
	}

	// Extract the commit SHA
	sha := ref.GetObject().GetSHA()

	// Get tree using commit SHA
	tree, _, err := c.client.Git.GetTree(c.ctx, owner, repo, sha, recursive)
	if err != nil {
		return nil, fmt.Errorf("failed to get tree: %w", err)
	}

	// Convert tree entries into tree nodes
	var nodes []TreeNode
	for _, entry := range tree.Entries {
		node := TreeNode{
			Path: entry.GetPath(),
			Type: entry.GetType(),
			Size: int64(entry.GetSize()),
		}
		nodes = append(nodes, node)
	}

	return nodes, nil
}
