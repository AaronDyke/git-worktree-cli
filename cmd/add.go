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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if !utils.IsGitRepo() {
			fmt.Println("Not inside a git repo")
			return
		}

		utils.GitFetch()

		if !utils.GitBranchExists(args[0]) {
			fmt.Println("Branch does not exist")
			return
		}

		WorktreeDir := utils.GitWorktreeDir(args[0])

		gitCmd := exec.Command("git", "worktree", "add", WorktreeDir, args[0])
		out, err := gitCmd.Output()
		if err != nil {
			fmt.Println("Error adding worktree", err)
		}
		fmt.Println(string(out))
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
