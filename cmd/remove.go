/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	gitUtils "github.com/AaronDyke/git-worktree-cli/pkg/git"
	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if !gitUtils.IsGitRepo() {
			return
		}

		branchName := ""
		if len(args) == 0 {
			// User has not specified a worktree branch, prompt for one
			branch, err := gitUtils.PromptForWorktree()
			if err != nil {
				fmt.Println("Error getting worktree", err)
				return
			}
			branchName = branch
		} else if len(args) == 1 {
			// User has specified a worktree branch
			if gitUtils.WorktreeExists(args[0]) {
				branchName = args[0]
			} else {
				fmt.Println("Worktree does not exist")
				return
			}
		} else {
			fmt.Println("Too many arguments")
			return
		}

		gitUtils.RemoveWorkTree(branchName)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
