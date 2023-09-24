package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func GitTopLevel() string {
	gitCmd := exec.Command("git", "rev-parse", "--show-toplevel")
	out, err := gitCmd.Output()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return string(out[:len(out)-1])
}

func GitWorktreeDir(branch string) string {
	topLevel := GitTopLevel()
	parentFolder := topLevel[:strings.LastIndex(topLevel, "/")]
	workingFolder := topLevel[strings.LastIndex(topLevel, "/")+1:]
	worktreeDir := fmt.Sprintf("%s/.worktrees/%s/%s", parentFolder, workingFolder, branch)
	// fmt.Println(worktreeDir)
	return worktreeDir
}

func GitBranchExists(branch string) bool {
	gitCmd := exec.Command("git", "rev-parse", "--verify", branch)
	err := gitCmd.Run()
	return err == nil
}

func GitFetch() {
	gitCmd := exec.Command("git", "fetch")
	_, err := gitCmd.Output()
	if err != nil {
		fmt.Println(err)
	}
}

func IsGitRepo() bool {
	gitCmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	out, err := gitCmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	return string(out) != "true"
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func OpenDir(dir string) {
	openCmd := exec.Command("code", dir)
	_, err := openCmd.Output()
	if err != nil {
		fmt.Println("Error opening worktree using VS Code", err)
		return
	}
}

