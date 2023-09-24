/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

type Worktree struct {
	Branch string
	Path   string
	Hash   string
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List current worktrees",
	Long:  `List current worktrees`,
	Run: func(cmd *cobra.Command, args []string) {
		gitCmd := exec.Command("git", "worktree", "list")
		out, err := gitCmd.Output()
		if err != nil {
			fmt.Println(err)
		}

		gitCmdOutput := string(out)
		gitCmdOutputLines := strings.Split(gitCmdOutput, "\n")
		// remove the empty string at the end of the slice
		gitCmdOutputLines = gitCmdOutputLines[:len(gitCmdOutputLines)-1]

		for _, line := range gitCmdOutputLines {
			lineSplit := strings.Split(line, " ")
			clean_lineSplit := []string{}
			for _, str := range lineSplit {
				if str != "" {
					clean_lineSplit = append(clean_lineSplit, str)
				}
			}

			worktree := Worktree{
				Path:   clean_lineSplit[0],
				Hash:   clean_lineSplit[1],
				Branch: clean_lineSplit[2],
			}

			if verbose, _ := cmd.Flags().GetBool("verbose"); verbose {
				fmt.Println(worktree.Path, worktree.Branch, worktree.Hash)
			} else {
				fmt.Println(worktree.Branch)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolP("verbose", "v", false, "Verbose output")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
