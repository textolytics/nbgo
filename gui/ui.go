package gui

import (
	"context"
	"fmt"
	"sync"

	"github.com/textolytics/nbgo/cli"
	"github.com/textolytics/nbgo/conf"
	"github.com/textolytics/nbgo/core"
	"github.com/textolytics/nbgo/logs"
)

// UITheme represents the UI theme settings
type UITheme struct {
	Name      string
	Colors    map[string]string
	Solarized bool
	Fancy     bool
	Borders   bool
}

// SolarizedTheme returns a Solarized dark theme
func SolarizedTheme() *UITheme {
	return &UITheme{
		Name:      "solarized-dark",
		Solarized: true,
		Fancy:     true,
		Borders:   true,
		Colors: map[string]string{
			"base03":  "#002b36",
			"base02":  "#073642",
			"base01":  "#586e75",
			"base00":  "#657b83",
			"base0":   "#839496",
			"base1":   "#93a1a1",
			"base2":   "#eee8d5",
			"base3":   "#fdf6e3",
			"yellow":  "#b58900",
			"orange":  "#cb4b16",
			"red":     "#dc322f",
			"magenta": "#d33682",
			"violet":  "#6c71c4",
			"blue":    "#268bd2",
			"cyan":    "#2aa198",
			"green":   "#859900",
			"success": "#859900",
			"error":   "#dc322f",
			"warning": "#b58900",
			"info":    "#268bd2",
		},
	}
}

// UIManager manages the user interface
type UIManager struct {
	mu              sync.RWMutex
	theme           *UITheme
	logger          logs.Logger
	commandRegistry *cli.CommandRegistry
	config          *conf.Config
	providers       map[string]core.Provider
	views           map[string]*View
	activeView      string
	sessions        map[string]*UISession
	commandHistory  []string
	autoDiscovery   bool
	keyboardNav     bool
}

// NewUIManager creates a new UI manager
func NewUIManager(logger logs.Logger, cmdReg *cli.CommandRegistry, cfg *conf.Config) *UIManager {
	return &UIManager{
		theme:           SolarizedTheme(),
		logger:          logger,
		commandRegistry: cmdReg,
		config:          cfg,
		providers:       make(map[string]core.Provider),
		views:           make(map[string]*View),
		sessions:        make(map[string]*UISession),
		commandHistory:  make([]string, 0),
		autoDiscovery:   true,
		keyboardNav:     true,
	}
}

// Initialize initializes the UI manager
func (um *UIManager) Initialize(ctx context.Context) error {
	um.mu.Lock()
	defer um.mu.Unlock()

	um.logger.Info("Initializing UI Manager")

	// Initialize default views
	views := map[string]*View{
		"dashboard":     NewView("dashboard", "Main Dashboard", "System Dashboard"),
		"data":          NewView("data", "Data Explorer", "Browse Market Data"),
		"debug":         NewView("debug", "Debug Console", "Debug Information"),
		"cli":           NewView("cli", "CLI Console", "Command Line Interface"),
		"terminal":      NewView("terminal", "Terminal", "System Terminal"),
		"api":           NewView("api", "API Explorer", "API Endpoints"),
		"monitoring":    NewView("monitoring", "Monitoring", "System Monitoring"),
		"logs":          NewView("logs", "Logs", "System Logs"),
		"environment":   NewView("environment", "Environment", "Environment Variables"),
		"configuration": NewView("configuration", "Configuration", "System Configuration"),
	}

	for name, view := range views {
		um.views[name] = view
	}

	um.activeView = "dashboard"
	um.logger.Info("UI Manager initialized successfully")
	return nil
}

// GetTheme returns the current theme
func (um *UIManager) GetTheme() *UITheme {
	um.mu.RLock()
	defer um.mu.RUnlock()
	return um.theme
}

// SetTheme sets the UI theme
func (um *UIManager) SetTheme(theme *UITheme) {
	um.mu.Lock()
	defer um.mu.Unlock()
	um.theme = theme
	um.logger.Infof("Theme changed to: %s", theme.Name)
}

// RegisterProvider registers a provider with the UI
func (um *UIManager) RegisterProvider(name string, provider core.Provider) error {
	um.mu.Lock()
	defer um.mu.Unlock()

	if _, exists := um.providers[name]; exists {
		return fmt.Errorf("provider %s already registered", name)
	}

	um.providers[name] = provider
	um.logger.Infof("Provider registered: %s", name)
	return nil
}

// GetProvider retrieves a provider by name
func (um *UIManager) GetProvider(name string) (core.Provider, error) {
	um.mu.RLock()
	defer um.mu.RUnlock()

	provider, exists := um.providers[name]
	if !exists {
		return nil, fmt.Errorf("provider %s not found", name)
	}
	return provider, nil
}

// ListProviders returns all registered providers
func (um *UIManager) ListProviders() map[string]core.Provider {
	um.mu.RLock()
	defer um.mu.RUnlock()

	providers := make(map[string]core.Provider)
	for k, v := range um.providers {
		providers[k] = v
	}
	return providers
}

// SwitchView switches to a different view
func (um *UIManager) SwitchView(viewName string) error {
	um.mu.Lock()
	defer um.mu.Unlock()

	if _, exists := um.views[viewName]; !exists {
		return fmt.Errorf("view %s not found", viewName)
	}

	um.activeView = viewName
	um.logger.Infof("Switched to view: %s", viewName)
	return nil
}

// GetActiveView returns the active view
func (um *UIManager) GetActiveView() *View {
	um.mu.RLock()
	defer um.mu.RUnlock()

	return um.views[um.activeView]
}

// ListViews returns all available views
func (um *UIManager) ListViews() map[string]*View {
	um.mu.RLock()
	defer um.mu.RUnlock()

	views := make(map[string]*View)
	for k, v := range um.views {
		views[k] = v
	}
	return views
}

// EnableAutoDiscovery enables automatic command discovery
func (um *UIManager) EnableAutoDiscovery(enable bool) {
	um.mu.Lock()
	defer um.mu.Unlock()
	um.autoDiscovery = enable
}

// EnableKeyboardNavigation enables keyboard navigation
func (um *UIManager) EnableKeyboardNavigation(enable bool) {
	um.mu.Lock()
	defer um.mu.Unlock()
	um.keyboardNav = enable
}

// ExecuteCommand executes a command with auto-discovery
func (um *UIManager) ExecuteCommand(ctx context.Context, cmdName string, args []string) error {
	um.mu.Lock()
	um.commandHistory = append(um.commandHistory, cmdName)
	um.mu.Unlock()

	um.logger.Info(fmt.Sprintf("Executing command: %s %v", cmdName, args))

	// Command would be executed through the command registry
	return nil
}

// GetCommandHistory returns the command history
func (um *UIManager) GetCommandHistory() []string {
	um.mu.RLock()
	defer um.mu.RUnlock()

	history := make([]string, len(um.commandHistory))
	copy(history, um.commandHistory)
	return history
}

// CreateSession creates a new UI session
func (um *UIManager) CreateSession(sessionID string) *UISession {
	um.mu.Lock()
	defer um.mu.Unlock()

	session := &UISession{
		ID:           sessionID,
		Windows:      make(map[string]*Window),
		Environment:  make(map[string]string),
		Active:       true,
		CreatedAt:    int64(0), // Use proper timestamp
		LastActivity: int64(0),
	}

	um.sessions[sessionID] = session
	um.logger.Infof("Session created: %s", sessionID)
	return session
}

// GetSession retrieves a session
func (um *UIManager) GetSession(sessionID string) (*UISession, error) {
	um.mu.RLock()
	defer um.mu.RUnlock()

	session, exists := um.sessions[sessionID]
	if !exists {
		return nil, fmt.Errorf("session %s not found", sessionID)
	}
	return session, nil
}

// CloseSession closes a session
func (um *UIManager) CloseSession(sessionID string) error {
	um.mu.Lock()
	defer um.mu.Unlock()

	session, exists := um.sessions[sessionID]
	if !exists {
		return fmt.Errorf("session %s not found", sessionID)
	}

	session.Active = false
	um.logger.Infof("Session closed: %s", sessionID)
	return nil
}

// GetAllSessions returns all sessions
func (um *UIManager) GetAllSessions() map[string]*UISession {
	um.mu.RLock()
	defer um.mu.RUnlock()

	sessions := make(map[string]*UISession)
	for k, v := range um.sessions {
		sessions[k] = v
	}
	return sessions
}
