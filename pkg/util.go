package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/AaronDyke/git-worktree-cli/pkg/config"
)

func PathExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func OpenDir(dir string) {
	// Get configured editor
	editor, err := config.GetConfiguredEditor()
	if err != nil {
		fmt.Printf("Error getting editor: %v\n", err)
		fmt.Println("Tip: Configure an editor with 'wt config set editor <name>'")
		fmt.Println("     Or run 'wt config list' to see available editors")
		return
	}

	openCmd := exec.Command(editor.Command, dir)
	_, err = openCmd.Output()
	if err != nil {
		fmt.Printf("Error opening directory using %s: %v\n", editor.Name, err)
		return
	}
	fmt.Printf("Opened directory in %s\n", editor.Name)
}

func AskForConfirmation(s string) bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s [y/n]: ", s)

		response, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}

		response = strings.ToLower(strings.TrimSpace(response))

		if response == "y" || response == "yes" {
			return true
		} else if response == "n" || response == "no" {
			return false
		}
	}
}
