package github

import (
	"context"

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
