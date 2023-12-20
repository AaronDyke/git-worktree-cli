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

// openCmd represents the open command
var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open worktree in VSCode",
	Long:  `Open worktree in VSCode`,
	Run: func(cmd *cobra.Command, args []string) {
		if !gitUtils.IsGitRepo() {
			return
		}

		branchName := ""
		if len(args) == 0 {
			// User has not specified a branch, prompt for one
			branch, err := gitUtils.PromptForWorktree()
			if err != nil {
				fmt.Println("Error getting worktree", err)
				return
			}
			branchName = branch
		} else if len(args) == 1 {
			// User has specified a branch
			if gitUtils.BranchExists(args[0]) {
				branchName = args[0]
			} else {
				fmt.Println("Branch does not exist")
				return
			}
		}

		WorktreeDir := gitUtils.WorktreeDir(branchName)
		if !utils.PathExists(WorktreeDir) {
			fmt.Println("Worktree does not exist on your computer")
			if utils.AskForConfirmation(fmt.Sprintf("Branch %s does not exist. Do you want to create it?", args[0])) {
				gitUtils.AddWorkTree(branchName)
			} else {
				return
			}
		}
		fmt.Println("Opening worktree", WorktreeDir)
		gitUtils.SwitchWorkTree(branchName)
	},
}

func init() {
	rootCmd.AddCommand(openCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// openCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// openCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
