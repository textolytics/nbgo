package cli

import (
	"context"
	"fmt"
	"sync"
)

// Command represents a CLI command
type Command struct {
	Name        string
	Description string
	Usage       string
	Handler     CommandHandler
	Flags       map[string]*Flag
}

// CommandHandler is the callback for command execution
type CommandHandler func(ctx context.Context, args []string) error

// Flag represents a command flag
type Flag struct {
	Name        string
	Short       string
	Description string
	Required    bool
	Default     string
	Value       string
}

// CommandRegistry manages all available commands
type CommandRegistry struct {
	mu       sync.RWMutex
	commands map[string]*Command
}

// NewCommandRegistry creates a new command registry
func NewCommandRegistry() *CommandRegistry {
	return &CommandRegistry{
		commands: make(map[string]*Command),
	}
}

// Register registers a command
func (cr *CommandRegistry) Register(cmd *Command) error {
	cr.mu.Lock()
	defer cr.mu.Unlock()
	if _, exists := cr.commands[cmd.Name]; exists {
		return fmt.Errorf("command %s already registered", cmd.Name)
	}
	cr.commands[cmd.Name] = cmd
	return nil
}

// Get retrieves a command by name
func (cr *CommandRegistry) Get(name string) (*Command, bool) {
	cr.mu.RLock()
	defer cr.mu.RUnlock()
	cmd, exists := cr.commands[name]
	return cmd, exists
}

// List returns all registered commands
func (cr *CommandRegistry) List() []*Command {
	cr.mu.RLock()
	defer cr.mu.RUnlock()
	cmds := make([]*Command, 0, len(cr.commands))
	for _, cmd := range cr.commands {
		cmds = append(cmds, cmd)
	}
	return cmds
}

// Execute executes a command
func (cr *CommandRegistry) Execute(ctx context.Context, name string, args []string) error {
	cmd, exists := cr.Get(name)
	if !exists {
		return fmt.Errorf("command %s not found", name)
	}
	return cmd.Handler(ctx, args)
}

// CLI represents the command line interface
type CLI struct {
	registry *CommandRegistry
	mu       sync.RWMutex
}

// NewCLI creates a new CLI instance
func NewCLI() *CLI {
	return &CLI{
		registry: NewCommandRegistry(),
	}
}

// RegisterCommand registers a command with the CLI
func (c *CLI) RegisterCommand(cmd *Command) error {
	return c.registry.Register(cmd)
}

// Execute executes a command
func (c *CLI) Execute(ctx context.Context, name string, args []string) error {
	return c.registry.Execute(ctx, name, args)
}

// ListCommands returns all registered commands
func (c *CLI) ListCommands() []*Command {
	return c.registry.List()
}

// GetCommand retrieves a command by name
func (c *CLI) GetCommand(name string) (*Command, bool) {
	return c.registry.Get(name)
}
