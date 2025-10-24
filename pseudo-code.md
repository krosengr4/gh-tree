### main.go
- Purpose: Entry point that launches the CLI application
- Steps:
    1. Import cmd package
    2. Call the Execute() function that is in cmd/root.go
    3. Handle any errors that occur during execution
    4. Exit with the appropriate status code

### cmd/root.go
- Purpose: Define the root cobra command and all CLI flags
- Steps:
    1. Import necessary packages
    2. Define global variables for CLI flags
    3. Create the root command with Cobra (name, short description, long description)
    4. Create Run() function that:
        - Get the repository argument from command line
        - Parse the repo string (owner/repo or URL)
        - Create GitHub client with token
        - Fetch the repo tree from GitHub API
        - Format and print the tree (feature: colorize the tree)
        - *** Make sure to handle errors at each step
    5. Create init() function that handles all the flags:
        - --branch or -b for branch name
        - --depth or -d for max depth of the tree
        - --token or -t for github token
        - --no-color to disable color from the output
    6. Create the Execute() function that runs the root command

### internal/parser/repo.go
- Purpose: Parse different repository input formats into owner and repo name
- Steps:
    1. Define RepoInfo struct containing Owner and Repo fields
    2. Create ParseRepo(input string) function that:
        - Take the input "owner/repoName" and split on the "/"
        - Validate that both the owner and the repo are not empty
        - Return RepoInfo struct and any errors
    3. Create helper func for Validation:
        - Check for valid owner and repo
        - Handle edge cases (ex: trailing slashes, git suffix, etc)

### internal/github/client.go
- Purpose: Interact with GitHub API to fetch the repository tree
- Steps:
    1. Import GitHub API client library and oauth2
    2. Define a Client struct that wraps the GitHub API client
    3. Create NewClient(token string) function that:
        - Create oauth2 token source if token provided
        - Initialize GitHub client with authentification
        - Return the Client struct
    4. Create GetTree(owner, repo, branch string, recursive bool) function that:
        - Get the reference (branch / commit SHA) from GitHub
        - Extract the Commit SHA from the reference
        - Set recursive flag to get the full tree
        - Return the tree structure
        - Handle API errors (rate limits, 404, authentication, etc.)
    5. Create helper struct TreeNode to represent files/directories:
        - Fields: Path, Type (file or dir), Size
    6. Create function to convert the GitHub tree to the TreeNode struct
    7. Add rate limit checking and error handling

### internal/tree/formatter.go
- Purpose: Format the tree structure and add colors
- Steps:
    1. Import the "faith/color" color package
    2. Define colors for directory, file, and symlink
    3. Create FormatTree(nodes []TreeNode, macDepth int, useColor bool) function that:
        - Build a tree structure from flat list of nodes
        - Create a map of paths to organize hierarchy
    4. Create recursive printNode() helper function that:
        - Takes node, depth, prefix, and isLast parameters
        - Print tree characters (├──, └──, │)
        - Apply colors based on node type
        - Handle depth limiting
        - Recursively print children with updated prefix
    5. Create helper function for tree characters that:
        - Use box-drawing characters (├──, └──, │)
        - Track whether node is last child to adjust prefix
    6. Add sorting logic
        - Sort directories before files
        - Sort alphabetically within each category
    7. Create function to disable colors when needed
    8. Add file count and directory count summary at the end


### ***ORDER***
1. internal/parser/repo.go
2. internal/github/client.go
3. internal/tree/formatter.go
4. cmd/root.go
5. main.go
