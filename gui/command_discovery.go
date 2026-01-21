package gui
package gui

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
	"sync"
)

// CommandDiscovery handles automatic command discovery
type CommandDiscovery struct {
	mu              sync.RWMutex
	discoveredCmds  map[string]*DiscoveredCommand
	suggestions     map[string][]string
	completions     map[string][]string
}

// DiscoveredCommand represents a discovered command
type DiscoveredCommand struct {
	Name        string
	Description string
	Category    string
	Module      string
	Usage       string
	Aliases     []string
	Options     []CommandOption
	AutoComplete bool
}

// CommandOption represents a command option
type CommandOption struct {
	Name        string
	Short       string
	Description string
	Type        string
	Required    bool
	Default     string
}

// NewCommandDiscovery creates a new command discovery service
func NewCommandDiscovery() *CommandDiscovery {
	return &CommandDiscovery{
		discoveredCmds: make(map[string]*DiscoveredCommand),
		suggestions:    make(map[string][]string),
		completions:    make(map[string][]string),
	}
}

// DiscoverCommands discovers available commands from modules
func (cd *CommandDiscovery) DiscoverCommands(ctx context.Context, modulePaths []string) error {
	cd.mu.Lock()
	defer cd.mu.Unlock()

	for _, modulePath := range modulePaths {
		// Discover commands from module
		commands, err := cd.discoverModuleCommands(ctx, modulePath)
		if err != nil {
			continue
		}

		for _, cmd := range commands {
			cd.discoveredCmds[cmd.Name] = cmd
		}
	}

	return nil
}

// discoverModuleCommands discovers commands from a single module
func (cd *CommandDiscovery) discoverModuleCommands(ctx context.Context, modulePath string) ([]*DiscoveredCommand, error) {
	var commands []*DiscoveredCommand

	// Simulate command discovery from module
	commands = append(commands, &DiscoveredCommand{
		Name:        "start",
		Description: "Start the module",
		Category:    "lifecycle",
		Module:      modulePath,
		Usage:       "start [options]",
		AutoComplete: true,
		Options: []CommandOption{
			{Name: "daemon", Short: "d", Description: "Run as daemon", Type: "bool"},
			{Name: "config", Short: "c", Description: "Config file", Type: "string"},
		},
	})

	commands = append(commands, &DiscoveredCommand{
		Name:        "stop",
		Description: "Stop the module",
		Category:    "lifecycle",
		Module:      modulePath,
		Usage:       "stop [options]",
		AutoComplete: true,
	})

	commands = append(commands, &DiscoveredCommand{
		Name:        "restart",
		Description: "Restart the module",
		Category:    "lifecycle",
		Module:      modulePath,
		Usage:       "restart [options]",
		AutoComplete: true,
	})

	return commands, nil
}

// GetSuggestions returns command suggestions for a prefix
func (cd *CommandDiscovery) GetSuggestions(prefix string) []string {
	cd.mu.RLock()
	defer cd.mu.RUnlock()

	var suggestions []string
	for name, cmd := range cd.discoveredCmds {
		if strings.HasPrefix(name, prefix) {
			suggestions = append(suggestions, fmt.Sprintf("%s - %s", name, cmd.Description))
		}
	}
	return suggestions
}

// GetCompletions returns command completions
func (cd *CommandDiscovery) GetCompletions(cmdName string) []string {
	cd.mu.RLock()
	defer cd.mu.RUnlock()

	cmd, exists := cd.discoveredCmds[cmdName]
	if !exists {
		return []string{}
	}

	var completions []string
	for _, opt := range cmd.Options {
		if opt.Short != "" {
			completions = append(completions, fmt.Sprintf("-%s", opt.Short))
		}
		completions = append(completions, fmt.Sprintf("--%s", opt.Name))
	}
	return completions
}

// GetCommand retrieves a discovered command
func (cd *CommandDiscovery) GetCommand(name string) (*DiscoveredCommand, error) {
	cd.mu.RLock()
	defer cd.mu.RUnlock()

	cmd, exists := cd.discoveredCmds[name]
	if !exists {
		return nil, fmt.Errorf("command %s not found", name)
	}
	return cmd, nil
}

// ListCommands returns all discovered commands
func (cd *CommandDiscovery) ListCommands() map[string]*DiscoveredCommand {
	cd.mu.RLock()
	defer cd.mu.RUnlock()

	commands := make(map[string]*DiscoveredCommand)
	for k, v := range cd.discoveredCmds {
		commands[k] = v
	}
	return commands
}

// ListCommandsByCategory returns commands by category
func (cd *CommandDiscovery) ListCommandsByCategory(category string) []*DiscoveredCommand {
	cd.mu.RLock()
	defer cd.mu.RUnlock()

	var commands []*DiscoveredCommand
	for _, cmd := range cd.discoveredCmds {
		if cmd.Category == category {
			commands = append(commands, cmd)
		}
	}
	return commands
}

// ExecuteWithAutoComplete executes a command with auto-completion
func (cd *CommandDiscovery) ExecuteWithAutoComplete(ctx context.Context, cmdLine string) (string, error) {
	parts := strings.Fields(cmdLine)
	if len(parts) == 0 {
		return "", fmt.Errorf("empty command")
	}

	cmdName := parts[0]
	args := parts[1:]

	// Execute the command
	cmd := exec.CommandContext(ctx, cmdName, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return string(output), err
	}

	return string(output), nil
}

// GetManPageForCommand returns the man page for a command
func (cd *CommandDiscovery) GetManPageForCommand(cmdName string) (string, error) {
	cmd, err := cd.GetCommand(cmdName)
	if err != nil {
		return "", err
	}

	manPage := fmt.Sprintf(`
NAME
    %s - %s

USAGE
    %s

CATEGORY
    %s

MODULE
    %s

OPTIONS
`, cmd.Name, cmd.Description, cmd.Usage, cmd.Category, cmd.Module)

	for _, opt := range cmd.Options {
		manPage += fmt.Sprintf("    -%s, --%s\n        %s\n", opt.Short, opt.Name, opt.Description)
	}

	return manPage, nil
}
