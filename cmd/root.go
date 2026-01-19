package cmd

import (
	"fmt"
	"gh-tree/internal/github"
	"gh-tree/internal/parser"
	"gh-tree/internal/tree"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// Global variables for flags
var (
	branch  string
	depth   int
	token   string
	noColor bool
)

// Root commands
var rootCmd = &cobra.Command{
	Use:   "gh-tree [owner/repo]",
	Short: "Display GitHub repository tree structure",
	Long:  `gh-tree fetches and displays the file tree of a GitHub repository in a tree format with colors`,
	Args:  cobra.ExactArgs(1),
	Run:   runTree,
}

// Run function. Main logic
func runTree(cmd *cobra.Command, args []string) {
	// Get repository argument
	repoInput := args[0]

	// Parse the repo string
	repoInfo, err := parser.ParseRepo(repoInput)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse repository")
	}

	// Create GitHub client with token
	client := github.NewClient(token)

	// Set default branch if not provided
	if branch == "" {
		branch = "main"
	}

	// Fetch the repo tree from github
	nodes, err := client.GetTree(repoInfo.Owner, repoInfo.RepoName, branch, true)
	if err != nil {
		log.Error().Err(err).Msg("Failed to fetch the repository tree")
		os.Exit(1)
	}

	// Format and print the tree
	useColor := !noColor
	tree.FormatTree(nodes, depth, useColor)
}

// Init func to set up all the flags
func init() {
	// Branch flag
	rootCmd.Flags().StringVarP(&branch, "branch", "b", "", "Branch name (default: main)")

	// Depth flag
	rootCmd.Flags().IntVarP(&depth, "depth", "d", 0, "Maximum depth of tree (0 = unlimited)")

	// Token flag
	rootCmd.Flags().StringVarP(&token, "token", "t", "", "GitHub personal access token")

	// No-color flag
	rootCmd.Flags().BoolVar(&noColor, "no-color", false, "Disable color input")
}

// Start the CLI
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
