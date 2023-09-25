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
			fmt.Println("Not inside a git repo")
			return
		}

		if !gitUtils.BranchExists(args[0]) {
			fmt.Println("Branch does not exist")
			return
		}

		WorktreeDir := gitUtils.WorktreeDir(args[0])
		if !utils.PathExists(WorktreeDir) {
			fmt.Println("Worktree does not exist on your computer")
			return
		}

		utils.OpenDir(WorktreeDir)
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
