/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	utils "github.com/AaronDyke/git-worktree-cli/pkg"
	gitUtils "github.com/AaronDyke/git-worktree-cli/pkg/git"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Create a new worktree for the given branch",
	Long:  `Create a new worktree for the given branch.`,
	Run: func(cmd *cobra.Command, args []string) {
		if !gitUtils.IsGitRepo() {
			return
		}

		gitUtils.Fetch()

		branchName := ""
		if len(args) == 0 {
			// User has not specified a branch, prompt for one
			branch, er := gitUtils.PromptForBranch()
			if er != nil {
				fmt.Println("Error getting branch", er)
				return
			}
			branchName = branch
		} else if len(args) == 1 {
			if gitUtils.WorktreeExists(args[0]) {
				fmt.Println("Worktree already exists")
				if utils.AskForConfirmation("Do you want to switch to the worktree?") {
					gitUtils.SwitchWorkTree(args[0])
				}
				return
			}
			if gitUtils.BranchExists(args[0]) {
				branchName = args[0]
			} else {
				// Branch does not exist
				if branch, _ := cmd.Flags().GetBool("branch"); branch {
					// User has specified to create a new branch
					gitUtils.CreateBranch(args[0])
				} else {
					// User has not specified to create a new branch
					if utils.AskForConfirmation(fmt.Sprintf("Branch %s does not exist. Do you want to create it?", args[0])) {
						gitUtils.CreateBranch(args[0])
					} else {
						return
					}
				}
				branchName = args[0]
			}
		}

		if gitUtils.WorktreeExists(branchName) {
			fmt.Println("Worktree already exists")
			if utils.AskForConfirmation("Do you want to switch to the worktree?") {
				gitUtils.SwitchWorkTree(branchName)
			}
			return
		} else {
			gitUtils.AddWorkTree(branchName)
		}

		if open, _ := cmd.Flags().GetBool("open"); open {
			gitUtils.SwitchWorkTree(branchName)
		} else if utils.AskForConfirmation("Do you want to switch to the new worktree?") {
			gitUtils.SwitchWorkTree(branchName)
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
