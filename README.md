# Git Worktree CLI

A command-line tool to easily manage Git worktrees with an intuitive interface and VS Code integration.

## Overview

Git Worktree CLI (alias: `wt`) simplifies working with Git worktrees by providing an easy-to-use interface for creating, managing, and switching between worktrees. Instead of manually managing worktree directories and paths, this tool handles the organization automatically and integrates seamlessly with VS Code.

## What are Git Worktrees?

Git worktrees allow you to have multiple working directories attached to the same repository. This is useful when you need to:
- Work on multiple branches simultaneously without stashing changes
- Quickly switch between features without losing your current work-in-progress
- Review pull requests while keeping your current development environment intact
- Run tests on one branch while developing on another

## Features

- **Easy Worktree Management**: Create, list, open, and remove worktrees with simple commands
- **Interactive Prompts**: Use interactive branch/worktree selection when commands are run without arguments
- **Automatic Organization**: Worktrees are organized in a `.worktrees` directory structure
- **Branch Creation**: Create new branches and their corresponding worktrees in one command
- **VS Code Integration**: Automatically open worktrees in VS Code
- **Smart Workflows**: Automatic fetching, confirmation prompts, and existence checks
- **Cross-platform**: Works on Linux, macOS, and Windows

## Installation

### From Source

```bash
# Clone the repository
git clone https://github.com/AaronDyke/git-worktree-cli.git
cd git-worktree-cli

# Build and install
make build
```

This will install the binary as `wt` in your `$GOPATH/bin` directory.

### Using Go Install

```bash
go install github.com/AaronDyke/git-worktree-cli@latest
```

### Using GoReleaser (Pre-built Binaries)

Download pre-built binaries from the [releases page](https://github.com/AaronDyke/git-worktree-cli/releases).

## Requirements

- Git 2.5+ (for worktree support)
- Go 1.20+ (for building from source)
- VS Code (optional, for `open` command and automatic opening)

## Worktree Organization

This tool organizes worktrees in a consistent directory structure:

```
parent-directory/
├── your-repo/              # Main repository
└── .worktrees/
    └── your-repo/
        ├── feature-branch/
        ├── bugfix-branch/
        └── another-branch/
```

For example, if your main repository is at `/home/user/projects/myapp`, worktrees will be created in `/home/user/projects/.worktrees/myapp/`.

## Usage

### Basic Commands

#### List Worktrees

```bash
wt list
```

Displays all existing worktrees for the current repository.

#### Add a Worktree

```bash
# Interactive mode - prompts you to select a branch
wt add

# Specify a branch name
wt add feature-branch

# Create a new branch and add it as a worktree
wt add -b new-feature

# Add worktree and open in VS Code immediately
wt add -o feature-branch
```

**Flags:**
- `-b, --branch`: Create a new branch if it doesn't exist
- `-o, --open`: Automatically open the worktree in VS Code after creation

**Behavior:**
- If the branch doesn't exist, you'll be prompted to create it
- If a worktree already exists for the branch, you'll be asked if you want to switch to it
- After creation, you'll be prompted to open the worktree in VS Code
- Automatically runs `git fetch` before processing

#### Open a Worktree

```bash
# Interactive mode - prompts you to select a worktree
wt open

# Open a specific worktree
wt open feature-branch
```

Opens the specified worktree in VS Code. If the worktree doesn't exist locally, you'll be prompted to create it.

#### Remove a Worktree

```bash
# Interactive mode - prompts you to select a worktree
wt remove

# Remove a specific worktree
wt remove feature-branch
```

Removes the specified worktree from your system and Git's worktree list.

## Examples

### Typical Workflow

```bash
# Start working on a new feature
wt add -b feature/new-login

# This will:
# 1. Create a new branch called "feature/new-login"
# 2. Add it as a worktree at .worktrees/your-repo/feature/new-login
# 3. Prompt you to open it in VS Code

# Later, switch to work on a bug fix
wt add -o hotfix/critical-bug

# List all your worktrees
wt list

# When done with a feature
wt remove feature/new-login
```

### Working on Multiple Features

```bash
# Create worktrees for different tasks
wt add -b feature/authentication
wt add -b feature/dashboard
wt add -b bugfix/login-error

# List to see all worktrees
wt list

# Switch between them as needed
wt open feature/authentication
wt open feature/dashboard

# Clean up when done
wt remove feature/authentication
```

## Command Reference

| Command | Description | Aliases |
|---------|-------------|---------|
| `wt add [branch]` | Create a new worktree for a branch | |
| `wt list` | List all worktrees | |
| `wt open [branch]` | Open a worktree in VS Code | |
| `wt remove [branch]` | Remove a worktree | |
| `wt help` | Display help information | |

### Global Behavior

- All commands verify you're inside a Git repository before executing
- Interactive prompts are provided when branch/worktree arguments are omitted
- Confirmation prompts prevent accidental operations

## Development

### Building

```bash
# Build and install to $GOPATH/bin
make build

# Watch mode (requires entr)
make watch
```

### Project Structure

```
.
├── cmd/                    # Command implementations
│   ├── root.go            # Root command setup
│   ├── add.go             # Add worktree command
│   ├── list.go            # List worktrees command
│   ├── open.go            # Open worktree command
│   └── remove.go          # Remove worktree command
├── pkg/
│   ├── git/
│   │   └── gitUtils.go    # Git operations utilities
│   └── util.go            # General utilities
├── main.go                # Application entry point
├── go.mod                 # Go module definition
├── Makefile              # Build configuration
└── .goreleaser.yaml      # Release configuration
```

### Dependencies

- [cobra](https://github.com/spf13/cobra) - CLI framework
- [promptui](https://github.com/manifoldco/promptui) - Interactive prompts

## Troubleshooting

### "Not inside a git repo"

Make sure you're running the command from within a Git repository.

### VS Code doesn't open

Ensure the `code` command is available in your PATH. You can verify this by running:
```bash
which code  # On macOS/Linux
where code  # On Windows
```

To install the `code` command in VS Code:
1. Open VS Code
2. Open the Command Palette (Cmd+Shift+P on macOS, Ctrl+Shift+P on Windows/Linux)
3. Type "Shell Command: Install 'code' command in PATH"

### Worktree already exists

If you try to create a worktree that already exists, the tool will offer to switch to it instead.

## Contributing

Contributions are welcome! Please feel free to submit issues or pull requests.

## License

This project is open source. See the LICENSE file for details.

## Acknowledgments

Built with:
- [Cobra](https://github.com/spf13/cobra) - A Commander for modern Go CLI interactions
- [promptui](https://github.com/manifoldco/promptui) - Interactive prompt library for Go
