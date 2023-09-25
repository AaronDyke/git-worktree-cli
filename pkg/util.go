package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func PathExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func OpenDir(dir string) {
	openCmd := exec.Command("code", dir)
	_, err := openCmd.Output()
	if err != nil {
		fmt.Println("Error opening worktree using VS Code", err)
		return
	}
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
