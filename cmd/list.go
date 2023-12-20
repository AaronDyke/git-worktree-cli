/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	gitUtils "github.com/AaronDyke/git-worktree-cli/pkg/git"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List current worktrees",
	Long:  `List current worktrees`,
	Run: func(cmd *cobra.Command, args []string) {
		if !gitUtils.IsGitRepo() {
			return
		}

		worktrees := gitUtils.WorktreeList()
		fmt.Println("")
		for _, worktree := range worktrees {
			fmt.Println(worktree)
		}
		fmt.Println("")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
