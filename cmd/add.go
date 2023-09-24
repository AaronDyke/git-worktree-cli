/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os/exec"

	utils "github.com/AaronDyke/git-worktree-cli/pkg"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Create a new worktree for the given branch",
	Long:  `Create a new worktree for the given branch.`,
	Run: func(cmd *cobra.Command, args []string) {
		if !utils.IsGitRepo() {
			fmt.Println("Not inside a git repo")
			return
		}

		utils.GitFetch()

		if !utils.GitBranchExists(args[0]) {
			if branch, _ := cmd.Flags().GetBool("branch"); branch {
				utils.CreateGitBranch(args[0])
			} else {
				if utils.AskForConfirmation(fmt.Sprintf("Branch %s does not exist. Do you wnat to create it?", args[0])) {
					utils.CreateGitBranch(args[0])
				} else {
					return
				}
			}
		}

		WorktreeDir := utils.GitWorktreeDir(args[0])

		gitCmd := exec.Command("git", "worktree", "add", WorktreeDir, args[0])
		out, err := gitCmd.Output()
		if err != nil {
			fmt.Println("Error adding worktree", err)
		}
		fmt.Println(string(out))
		if open, _ := cmd.Flags().GetBool("open"); open {
			utils.OpenDir(WorktreeDir)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().BoolP("branch", "b", false, "Create a new branch and add it as a worktree")
	addCmd.Flags().BoolP("open", "o", false, "Open the worktree, after creation, in VS Code")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
