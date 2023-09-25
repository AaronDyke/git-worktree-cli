package gitUtils

import (
	"fmt"
	"os/exec"
	"strings"
)

func IsGitRepo() bool {
	gitCmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	out, err := gitCmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	return string(out) != "true"
}

func TopLevel() string {
	gitCmd := exec.Command("git", "rev-parse", "--show-toplevel")
	out, err := gitCmd.Output()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return string(out[:len(out)-1])
}

func BranchExists(branch string) bool {
	gitCmd := exec.Command("git", "rev-parse", "--verify", branch)
	err := gitCmd.Run()
	return err == nil
}

func Fetch() {
	gitCmd := exec.Command("git", "fetch")
	_, err := gitCmd.Output()
	if err != nil {
		fmt.Println(err)
	}
}

func WorktreeList() []string {
	gitCmd := exec.Command("git", "worktree", "list")
	out, err := gitCmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	worktrees := strings.Split(string(out), "\n")
	// remove empty string at end of slice
	worktrees = worktrees[:len(worktrees)-1]
	return worktrees
}

func WorktreeDir(branch string) string {
	topLevel := TopLevel()
	parentFolder := topLevel[:strings.LastIndex(topLevel, "/")]
	workingFolder := topLevel[strings.LastIndex(topLevel, "/")+1:]
	worktreeDir := fmt.Sprintf("%s/.worktrees/%s/%s", parentFolder, workingFolder, branch)
	// fmt.Println(worktreeDir)
	return worktreeDir
}

func CreateBranch(branch string) {
	gitCmd := exec.Command("git", "branch", branch)
	_, err := gitCmd.Output()
	if err != nil {
		fmt.Println(err)
	}
}
