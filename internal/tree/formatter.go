package tree

import (
	"fmt"
	"path/filepath"
	"sort"
	"strings"

	"github.com/fatih/color"
)

// Represents a file or directory in the tree
type TreeNode struct {
	Path string
	Type string
	Size int
}

// Represents a node in the tree hierarchy
type node struct {
	name     string
	isDir    bool
	children []*node
	path     string
}

// Color definitions
var (
	dirColor     = color.New(color.FgBlue, color.Bold)
	fileColor    = color.New(color.FgWhite)
	symlinkColor = color.New(color.FgCyan)
)

// Formats and prints the tree structure
func FormatTree(nodes []TreeNode, maxDepth int, useColor bool) {
	if !useColor {
		color.NoColor = true
	}

	// Build the tree structure
	root := buildTree(nodes)

	// Print the tree
	fileCount := 0
	dirCount := 0
	printNode(root, 0, "", true, maxDepth, &fileCount, &dirCount)

	// Print summary
	fmt.Printf("\n%d directories, %d files\n", dirCount, fileCount)
}

// Construct a tree from flat list of nodes
func buildTree(nodes []TreeNode) *node {
	root := &node{
		name:     ".",
		isDir:    true,
		children: []*node{},
	}

	// Create a map for quick lookup
	nodeMap := make(map[string]*node)
	nodeMap["."] = root

	// Sort nodes by path to ensure parents are processed before children
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].Path < nodes[j].Path
	})

	// Build the tree
	for _, tn := range nodes {
		parts := strings.Split(tn.Path, "/")
		currentPath := ""

		for i, part := range parts {
			if i > 0 {
				currentPath += "/"
			}
			currentPath += part

			if _, exists := nodeMap[currentPath]; !exists {
				newNode := &node{
					name:     part,
					isDir:    i < len(parts)-1 || tn.Type == "tree",
					children: []*node{},
					path:     currentPath,
				}

				// Find parent and add as child
				parentPath := filepath.Dir(currentPath)
				if parentPath == "." && i == 0 {
					root.children = append(root.children, newNode)
				} else if parent, ok := nodeMap[parentPath]; ok {
					parent.children = append(parent.children, newNode)
				}

				nodeMap[currentPath] = newNode
			}
		}
	}

	// Sort all children
	sortChildren(root)

	return root
}

// Recursively sorts children. directories first, then alphabetically
func sortChildren(n *node) {
	sort.Slice(n.children, func(i, j int) bool {
		if n.children[i].isDir != n.children[j].isDir {
			return n.children[i].isDir
		}
		return n.children[i].name < n.children[j].name
	})

	for _, child := range n.children {
		sortChildren(child)
	}
}

// Recursively prints a node and its children
func printNode(n *node, depth int, prefix string, isLast bool, maxDepth int, fileCount *int, dirCount *int) {
	// Skip root node
	if n.name == "." {
		for i, child := range n.children {
			printNode(child, 0, "", i == len(n.children)-1, maxDepth, fileCount, dirCount)
		}
		return
	}

	// Check depth limit
	if maxDepth > 0 && depth >= maxDepth {
		return
	}

	// Print current node
	if depth == 0 {
		// Top level, no prefix
		printNodeName(n)
		fmt.Println()
	} else {
		// Build tree characters
		var treeChar string
		if isLast {
			treeChar = "└── "
		} else {
			treeChar = "├── "
		}

		fmt.Print(prefix + treeChar)
		printNodeName(n)
		fmt.Println()
	}

	// Update counts
	if n.isDir {
		*dirCount++
	} else {
		*fileCount++
	}

	// Print children
	for i, child := range n.children {
		var childPrefix string
		if depth == 0 {
			childPrefix = ""
		} else {
			if isLast {
				childPrefix = prefix + "    "
			} else {
				childPrefix = prefix + "│   "
			}
		}
		printNode(child, depth+1, childPrefix, i == len(n.children)-1, maxDepth, fileCount, dirCount)
	}
}

// Prints a node name with the appropriate color
func printNodeName(n *node) {
	if n.isDir {
		dirColor.Print(n.name + "/")
	} else {
		fileColor.Print(n.name)
	}
}
