package config

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Editor represents a code editor configuration
type Editor struct {
	Name        string
	Command     string
	Description string
}

// Config represents the application configuration
type Config struct {
	Editor string `yaml:"editor"`
}

// SupportedEditors lists all supported editors
var SupportedEditors = []Editor{
	{Name: "vscode", Command: "code", Description: "Visual Studio Code"},
	{Name: "vscodium", Command: "codium", Description: "VSCodium (FOSS VS Code)"},
	{Name: "cursor", Command: "cursor", Description: "Cursor AI Editor"},
	{Name: "intellij", Command: "idea", Description: "IntelliJ IDEA"},
	{Name: "goland", Command: "goland", Description: "GoLand"},
	{Name: "sublime", Command: "subl", Description: "Sublime Text"},
	{Name: "vim", Command: "vim", Description: "Vim"},
	{Name: "neovim", Command: "nvim", Description: "Neovim"},
	{Name: "emacs", Command: "emacs", Description: "Emacs"},
	{Name: "atom", Command: "atom", Description: "Atom"},
	{Name: "webstorm", Command: "webstorm", Description: "WebStorm"},
	{Name: "pycharm", Command: "pycharm", Description: "PyCharm"},
	{Name: "phpstorm", Command: "phpstorm", Description: "PhpStorm"},
	{Name: "clion", Command: "clion", Description: "CLion"},
	{Name: "rider", Command: "rider", Description: "Rider"},
}

// GetConfigPath returns the path to the config file
func GetConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	configDir := filepath.Join(homeDir, ".config", "git-worktree-cli")
	return filepath.Join(configDir, "config.yaml"), nil
}

// LoadConfig loads the configuration from disk
func LoadConfig() (*Config, error) {
	configPath, err := GetConfigPath()
	if err != nil {
		return nil, err
	}

	// Return default config if file doesn't exist
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return &Config{Editor: ""}, nil
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

// SaveConfig saves the configuration to disk
func SaveConfig(config *Config) error {
	configPath, err := GetConfigPath()
	if err != nil {
		return err
	}

	// Create config directory if it doesn't exist
	configDir := filepath.Dir(configPath)
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return err
	}

	data, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, data, 0644)
}

// IsCommandAvailable checks if a command is available in PATH
func IsCommandAvailable(command string) bool {
	_, err := exec.LookPath(command)
	return err == nil
}

// DetectAvailableEditors returns a list of editors that are available on the system
func DetectAvailableEditors() []Editor {
	var available []Editor
	for _, editor := range SupportedEditors {
		if IsCommandAvailable(editor.Command) {
			available = append(available, editor)
		}
	}
	return available
}

// GetEditorByName returns an editor by its name
func GetEditorByName(name string) (*Editor, error) {
	for _, editor := range SupportedEditors {
		if editor.Name == name {
			return &editor, nil
		}
	}
	return nil, fmt.Errorf("editor '%s' not found", name)
}

// GetEditorByCommand returns an editor by its command
func GetEditorByCommand(command string) (*Editor, error) {
	for _, editor := range SupportedEditors {
		if editor.Command == command {
			return &editor, nil
		}
	}
	return nil, fmt.Errorf("editor with command '%s' not found", command)
}

// GetConfiguredEditor returns the configured editor or detects one if not configured
func GetConfiguredEditor() (*Editor, error) {
	config, err := LoadConfig()
	if err != nil {
		return nil, err
	}

	// If editor is configured, use it
	if config.Editor != "" {
		editor, err := GetEditorByName(config.Editor)
		if err != nil {
			// Try to find by command
			editor, err = GetEditorByCommand(config.Editor)
			if err != nil {
				return nil, fmt.Errorf("configured editor '%s' not found", config.Editor)
			}
		}

		// Check if the editor is actually available
		if !IsCommandAvailable(editor.Command) {
			return nil, fmt.Errorf("configured editor '%s' (command: %s) is not available in PATH", editor.Name, editor.Command)
		}

		return editor, nil
	}

	// Auto-detect available editors
	available := DetectAvailableEditors()
	if len(available) == 0 {
		return nil, fmt.Errorf("no supported editors found on your system")
	}

	// Return the first available editor (VS Code preferred if available)
	for _, editor := range available {
		if editor.Name == "vscode" {
			return &editor, nil
		}
	}

	return &available[0], nil
}

// SetEditor sets the editor in the configuration
func SetEditor(editorName string) error {
	// Validate the editor exists
	editor, err := GetEditorByName(editorName)
	if err != nil {
		// Try to find by command
		editor, err = GetEditorByCommand(editorName)
		if err != nil {
			return fmt.Errorf("editor '%s' is not supported", editorName)
		}
	}

	// Check if the editor is available
	if !IsCommandAvailable(editor.Command) {
		return fmt.Errorf("editor '%s' (command: %s) is not available in PATH", editor.Name, editor.Command)
	}

	config := &Config{Editor: editor.Name}
	return SaveConfig(config)
}
