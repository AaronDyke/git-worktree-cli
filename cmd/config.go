/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/AaronDyke/git-worktree-cli/pkg/config"
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage git-worktree-cli configuration",
	Long:  `Manage git-worktree-cli configuration including editor preferences.`,
}

// configSetCmd represents the config set command
var configSetCmd = &cobra.Command{
	Use:   "set <key> <value>",
	Short: "Set a configuration value",
	Long:  `Set a configuration value. Currently supports: editor`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]
		value := args[1]

		switch key {
		case "editor":
			err := config.SetEditor(value)
			if err != nil {
				fmt.Printf("Error setting editor: %v\n", err)
				return
			}
			fmt.Printf("Editor set to: %s\n", value)
		default:
			fmt.Printf("Unknown configuration key: %s\n", key)
			fmt.Println("Supported keys: editor")
		}
	},
}

// configGetCmd represents the config get command
var configGetCmd = &cobra.Command{
	Use:   "get <key>",
	Short: "Get a configuration value",
	Long:  `Get a configuration value. Currently supports: editor`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]

		switch key {
		case "editor":
			editor, err := config.GetConfiguredEditor()
			if err != nil {
				fmt.Printf("Error getting editor: %v\n", err)
				return
			}

			cfg, _ := config.LoadConfig()
			if cfg.Editor == "" {
				fmt.Printf("Editor: %s (auto-detected)\n", editor.Name)
			} else {
				fmt.Printf("Editor: %s\n", editor.Name)
			}
			fmt.Printf("Command: %s\n", editor.Command)
		default:
			fmt.Printf("Unknown configuration key: %s\n", key)
			fmt.Println("Supported keys: editor")
		}
	},
}

// configListCmd represents the config list command
var configListCmd = &cobra.Command{
	Use:   "list",
	Short: "List available editors",
	Long:  `List all supported editors and their availability on your system.`,
	Run: func(cmd *cobra.Command, args []string) {
		all, _ := cmd.Flags().GetBool("all")

		fmt.Println("Available Editors:")
		fmt.Println()

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
		fmt.Fprintln(w, "NAME\tCOMMAND\tAVAILABLE\tDESCRIPTION")
		fmt.Fprintln(w, "----\t-------\t---------\t-----------")

		currentEditor, _ := config.GetConfiguredEditor()

		for _, editor := range config.SupportedEditors {
			available := config.IsCommandAvailable(editor.Command)

			// Skip unavailable editors unless --all flag is used
			if !available && !all {
				continue
			}

			availableStr := "✗"
			if available {
				availableStr = "✓"
			}

			name := editor.Name
			if currentEditor != nil && currentEditor.Name == editor.Name {
				name = fmt.Sprintf("%s (current)", name)
			}

			fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", name, editor.Command, availableStr, editor.Description)
		}

		w.Flush()
		fmt.Println()
		fmt.Println("Use 'wt config set editor <name>' to set your preferred editor")
	},
}

// configDetectCmd represents the config detect command
var configDetectCmd = &cobra.Command{
	Use:   "detect",
	Short: "Detect available editors",
	Long:  `Detect and display editors that are currently available on your system.`,
	Run: func(cmd *cobra.Command, args []string) {
		available := config.DetectAvailableEditors()

		if len(available) == 0 {
			fmt.Println("No supported editors found on your system.")
			fmt.Println()
			fmt.Println("Supported editors:")
			for _, editor := range config.SupportedEditors {
				fmt.Printf("  - %s (%s)\n", editor.Name, editor.Command)
			}
			return
		}

		fmt.Printf("Found %d available editor(s):\n\n", len(available))

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
		fmt.Fprintln(w, "NAME\tCOMMAND\tDESCRIPTION")
		fmt.Fprintln(w, "----\t-------\t-----------")

		for _, editor := range available {
			fmt.Fprintf(w, "%s\t%s\t%s\n", editor.Name, editor.Command, editor.Description)
		}

		w.Flush()

		// Show which one would be used by default
		defaultEditor, err := config.GetConfiguredEditor()
		if err == nil {
			fmt.Println()
			cfg, _ := config.LoadConfig()
			if cfg.Editor == "" {
				fmt.Printf("Default editor (auto-detected): %s\n", defaultEditor.Name)
			} else {
				fmt.Printf("Configured editor: %s\n", defaultEditor.Name)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(configSetCmd)
	configCmd.AddCommand(configGetCmd)
	configCmd.AddCommand(configListCmd)
	configCmd.AddCommand(configDetectCmd)

	configListCmd.Flags().BoolP("all", "a", false, "Show all editors including unavailable ones")
}
