package gitUtils

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/manifoldco/promptui"
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

func WorktreeExists(branch string) bool {
	worktrees := WorktreeList()
	for _, worktree := range worktrees {
		checkingBranch := strings.Fields(worktree)[2]
		checkingBranch = checkingBranch[1 : len(checkingBranch)-1]
		if checkingBranch == branch {
			return true
		}
	}
	return false
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

func PromptForBranch() (string, error) {
	gitCmd := exec.Command("git", "branch")
	out, err := gitCmd.Output()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	branches := strings.Split(string(out), "\n")

	// remove leading and trailing whitespace
	for i, branch := range branches {
		branches[i] = strings.TrimSpace(branch)
	}

	prompt := promptui.Select{
		Label: "Select Git Branch",
		Items: branches,
	}

	_, result, err := prompt.Run()
	//  remove star or plus from selected branch
	if strings.HasPrefix(result, "*") {
		result = strings.TrimSpace(result[1:])
	}

	return result, err
}

func PromptForWorktree() (string, error) {
	worktrees := WorktreeList()

	prompt := promptui.Select{
		Label: "Select Git Worktree",
		Items: worktrees,
	}

	_, result, err := prompt.Run()
	result = strings.Fields(result)[2]
	result = result[1 : len(result)-1]
	return result, err
}

func CreateBranch(branch string) {
	gitCmd := exec.Command("git", "branch", branch)
	_, err := gitCmd.Output()
	if err != nil {
		fmt.Println(err)
	}
}

func AddWorkTree(branch string) {
	WorktreeDir := WorktreeDir(branch)
	gitCmd := exec.Command("git", "worktree", "add", WorktreeDir, branch)
	_, err := gitCmd.Output()
	if err != nil {
		fmt.Println(err)
	}
}
