# gh-tree

A fast CLI tool to visualize the file structure of any GitHub repository directly from your terminal, with beautiful colored output.

![gh-tree demo](screenshot.png)

## Features

- 🌳 **Beautiful tree visualization** - Display repository structure with box-drawing characters
- 🎨 **Colored output** - Directories and files are color-coded for easy reading
- 🔍 **Depth limiting** - Control how deep to traverse the directory tree
- 🔐 **Authentication support** - Use GitHub tokens for private repos and higher rate limits
- ⚡ **Fast** - Fetches entire repository structure in one API call
- 🎯 **Branch selection** - View any branch, not just the default

## Installation

### Option 1: Build from source

```bash
git clone https://github.com/kevinrosengren/gh-tree.git
cd gh-tree
go build -o gh-tree
```

### Option 2: Install with Go

```bash
go install github.com/kevinrosengren/gh-tree@latest
```

## Usage

### Basic Usage

```bash
gh-tree owner/repository
```

**Example:**
```bash
gh-tree anthropics/anthropic-sdk-go
```

### With Options

```bash
# View a specific branch
gh-tree anthropics/anthropic-sdk-go --branch main

# Limit tree depth
gh-tree anthropics/anthropic-sdk-go --depth 2

# Use authentication token
gh-tree myusername/private-repo --token ghp_yourtoken123

# Disable colors
gh-tree anthropics/anthropic-sdk-go --no-color

# Combine multiple options
gh-tree anthropics/anthropic-sdk-go -b main -d 3 -t ghp_token
```

## Command Line Flags

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--branch` | `-b` | `main` | Branch name to fetch |
| `--depth` | `-d` | `0` (unlimited) | Maximum depth of tree to display |
| `--token` | `-t` | `""` | GitHub personal access token |
| `--no-color` | - | `false` | Disable colored output |
| `--help` | `-h` | - | Display help information |

## GitHub Token (Optional)

While not required for public repositories, using a GitHub token provides several benefits:

- Access to private repositories
- Higher rate limits (5,000 requests/hour vs 60 requests/hour)
- Faster response times

### Creating a GitHub Token

1. Go to [GitHub Settings > Tokens](https://github.com/settings/tokens)
2. Click **Generate new token** → **Classic**
3. Give it a descriptive name (e.g., "gh-tree CLI")
4. Select scopes:
   - `repo` - for private repositories
   - `public_repo` - for public repositories only
5. Click **Generate token**
6. Copy the token (starts with `ghp_`)

### Using the Token

```bash
gh-tree owner/repo --token ghp_yourtoken123
```

Or set it as an environment variable:
```bash
export GITHUB_TOKEN=ghp_yourtoken123
gh-tree owner/repo -t $GITHUB_TOKEN
```

## Examples

### Display a public repository
```bash
./gh-tree krosengr4/byteboard
```

Output:
```
public/
├── file.svg
├── globe.svg
├── next.svg
├── vercel.svg
└── window.svg
src/
├── app/
│   ├── login/
│   │   └── page.tsx
│   ├── register/
│   │   └── page.tsx
│   ├── favicon.ico
│   ├── globals.css
│   ├── layout.tsx
│   └── page.tsx
├── contexts/
│   └── AuthContext.tsx
├── lib/
│   └── api.ts
└── store/
    └── useStore.ts

8 directories, 25 files
```

### Limit depth for large repositories
```bash
./gh-tree torvalds/linux --depth 1
```

### View a specific branch
```bash
./gh-tree facebook/react --branch canary
```

## How It Works

1. **Parse Input** - Extracts owner and repository name from input
2. **GitHub API** - Fetches repository tree structure using GitHub's Git Data API
3. **Build Tree** - Converts flat file list into hierarchical tree structure
4. **Format Output** - Renders tree with colors and box-drawing characters
5. **Display** - Prints formatted tree to terminal

## Project Structure

```
gh-tree/
├── cmd/
│   └── root.go          # Cobra CLI commands and flags
├── internal/
│   ├── github/
│   │   └── client.go    # GitHub API client
│   ├── parser/
│   │   └── repo.go      # Repository string parser
│   └── tree/
│       └── formatter.go # Tree formatting and display
├── main.go              # Application entry point
├── go.mod
└── README.md
```

## Requirements

- Go 1.25 or higher
- Internet connection
- GitHub personal access token (optional, for private repos)

## Dependencies

- [cobra](https://github.com/spf13/cobra) - CLI framework
- [go-github](https://github.com/google/go-github) - GitHub API client
- [color](https://github.com/fatih/color) - Terminal colors
- [zerolog](https://github.com/rs/zerolog) - Structured logging
- [oauth2](https://golang.org/x/oauth2) - OAuth2 authentication

## Troubleshooting

### "branch not found" error
Some repositories use `master` instead of `main` as the default branch:
```bash
gh-tree owner/repo --branch master
```

### Rate limit errors
GitHub limits unauthenticated requests to 60 per hour. Use a token:
```bash
gh-tree owner/repo --token ghp_yourtoken
```

### Colors not displaying
If colors aren't rendering properly in your terminal:
```bash
gh-tree owner/repo --no-color
```

## Contributing

Contributions are welcome! Feel free to:
- Report bugs
- Suggest new features
- Submit pull requests

## License

MIT License - see LICENSE file for details

## Author

Kevin Rosengren - [GitHub](https://github.com/kevinrosengren)

---

**Note:** This tool uses the GitHub API and is subject to GitHub's rate limiting policies.
