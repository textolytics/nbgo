package gui

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/textolytics/nbgo/cli"
	"github.com/textolytics/nbgo/conf"
	"github.com/textolytics/nbgo/logs"
)

// Application represents the main GUI application
type Application struct {
	mu               sync.RWMutex
	name             string
	version          string
	config           *conf.Config
	logger           logs.Logger
	uiManager        *UIManager
	commandRegistry  *cli.CommandRegistry
	commandDiscovery *CommandDiscovery
	keyboardHandler  *KeyboardHandler
	navigationCtrl   *NavigationController
	managers         map[ManagerType]ManagerProvider
	running          bool
	startTime        time.Time
	sessionManager   *SessionManager
	eventBus         *EventBus
}

// SessionManager manages UI sessions
type SessionManager struct {
	mu       sync.RWMutex
	sessions map[string]*UISession
}

// EventBus handles events across the application
type EventBus struct {
	mu        sync.RWMutex
	listeners map[string][]func(Event)
}

// Event represents an application event
type Event struct {
	Type      string
	Timestamp int64
	Data      map[string]interface{}
	Source    string
}

// NewApplication creates a new GUI application
func NewApplication(name, version string, config *conf.Config, logger logs.Logger,
	cmdReg *cli.CommandRegistry) *Application {
	return &Application{
		name:             name,
		version:          version,
		config:           config,
		logger:           logger,
		commandRegistry:  cmdReg,
		commandDiscovery: NewCommandDiscovery(),
		keyboardHandler:  NewKeyboardHandler(),
		navigationCtrl:   NewNavigationController(),
		managers:         make(map[ManagerType]ManagerProvider),
		sessionManager: &SessionManager{
			sessions: make(map[string]*UISession),
		},
		eventBus: &EventBus{
			listeners: make(map[string][]func(Event)),
		},
		running: false,
	}
}

// Initialize initializes the application
func (app *Application) Initialize(ctx context.Context) error {
	app.mu.Lock()
	defer app.mu.Unlock()

	app.logger.Infof("Initializing NBGO GUI Application v%s", app.version)

	// Initialize UI Manager
	app.uiManager = NewUIManager(app.logger, app.commandRegistry, app.config)
	if err := app.uiManager.Initialize(ctx); err != nil {
		return fmt.Errorf("failed to initialize UI Manager: %w", err)
	}

	// Initialize Managers
	managers := []ManagerProvider{
		NewInstallPrerequisitesManager(app.logger, app.config, app.commandRegistry, app.uiManager),
		NewEnvironmentManager(app.logger, app.config, app.commandRegistry, app.uiManager),
		NewConfigureManager(app.logger, app.config, app.commandRegistry, app.uiManager),
		NewBuildManager(app.logger, app.config, app.commandRegistry, app.uiManager),
		NewInstallManager(app.logger, app.config, app.commandRegistry, app.uiManager),
	}

	for _, manager := range managers {
		if err := manager.Initialize(ctx); err != nil {
			app.logger.Warnf("Failed to initialize manager: %v", err)
			continue
		}
		app.managers[manager.GetType()] = manager
	}

	// Setup default keyboard bindings
	app.setupKeyboardBindings()

	// Setup default navigation
	if err := app.navigationCtrl.Navigate("dashboard"); err != nil {
		return fmt.Errorf("failed to navigate to dashboard: %w", err)
	}

	app.logger.Info("NBGO GUI Application initialized successfully")
	return nil
}

// setupKeyboardBindings sets up default keyboard bindings
func (app *Application) setupKeyboardBindings() {
	// Navigation bindings
	app.keyboardHandler.RegisterKeyBinding(KeyBinding{
		Key:         "Tab",
		Description: "Switch to next view",
		Category:    "navigation",
	})

	app.keyboardHandler.RegisterKeyBinding(KeyBinding{
		Key:         "Escape",
		Description: "Go back to previous view",
		Category:    "navigation",
	})

	// Command bindings
	app.keyboardHandler.RegisterKeyBinding(KeyBinding{
		Key:         ":",
		Modifiers:   []string{},
		Description: "Enter command mode",
		Category:    "command",
	})

	// Search binding
	app.keyboardHandler.RegisterKeyBinding(KeyBinding{
		Key:         "/",
		Description: "Enter search mode",
		Category:    "search",
	})

	// Quit binding
	app.keyboardHandler.RegisterKeyBinding(KeyBinding{
		Key:         "q",
		Modifiers:   []string{"ctrl"},
		Description: "Quit application",
		Category:    "application",
	})
}

// Start starts the application
func (app *Application) Start(ctx context.Context) error {
	app.mu.Lock()
	if app.running {
		app.mu.Unlock()
		return fmt.Errorf("application already running")
	}
	app.running = true
	app.startTime = time.Now()
	app.mu.Unlock()

	app.logger.Info("Starting NBGO GUI Application")

	// Create default session
	session := app.uiManager.CreateSession("default")
	app.sessionManager.AddSession(session)

	// Setup environment variables in session
	app.setupSessionEnvironment(session)

	app.logger.Info("NBGO GUI Application started successfully")
	return nil
}

// setupSessionEnvironment sets up environment variables in the session
func (app *Application) setupSessionEnvironment(session *UISession) {
	envVars := os.Environ()
	for _, envVar := range envVars {
		key := ""
		value := ""
		for i, c := range envVar {
			if c == '=' {
				key = envVar[:i]
				value = envVar[i+1:]
				break
			}
		}
		if key != "" {
			session.SetEnvironmentVariable(key, value)
		}
	}
}

// Stop stops the application
func (app *Application) Stop(ctx context.Context) error {
	app.mu.Lock()
	if !app.running {
		app.mu.Unlock()
		return fmt.Errorf("application not running")
	}
	app.running = false
	app.mu.Unlock()

	app.logger.Info("Stopping NBGO GUI Application")

	// Close all sessions
	sessions := app.sessionManager.GetAllSessions()
	for _, session := range sessions {
		session.Close()
	}

	app.logger.Info("NBGO GUI Application stopped successfully")
	return nil
}

// IsRunning returns if the application is running
func (app *Application) IsRunning() bool {
	app.mu.RLock()
	defer app.mu.RUnlock()
	return app.running
}

// GetUptime returns the application uptime
func (app *Application) GetUptime() time.Duration {
	app.mu.RLock()
	defer app.mu.RUnlock()
	if !app.running {
		return 0
	}
	return time.Since(app.startTime)
}

// GetManager retrieves a manager by type
func (app *Application) GetManager(managerType ManagerType) (ManagerProvider, error) {
	app.mu.RLock()
	defer app.mu.RUnlock()

	manager, exists := app.managers[managerType]
	if !exists {
		return nil, fmt.Errorf("manager %s not found", managerType)
	}
	return manager, nil
}

// ListManagers returns all managers
func (app *Application) ListManagers() map[ManagerType]ManagerProvider {
	app.mu.RLock()
	defer app.mu.RUnlock()

	managers := make(map[ManagerType]ManagerProvider)
	for k, v := range app.managers {
		managers[k] = v
	}
	return managers
}

// GetUIManager returns the UI manager
func (app *Application) GetUIManager() *UIManager {
	app.mu.RLock()
	defer app.mu.RUnlock()
	return app.uiManager
}

// DiscoverCommands discovers available commands
func (app *Application) DiscoverCommands(ctx context.Context, modulePaths []string) error {
	return app.commandDiscovery.DiscoverCommands(ctx, modulePaths)
}

// GetCommandSuggestions returns suggestions for a command prefix
func (app *Application) GetCommandSuggestions(prefix string) []string {
	return app.commandDiscovery.GetSuggestions(prefix)
}

// GetCommandCompletions returns completions for a command
func (app *Application) GetCommandCompletions(cmdName string) []string {
	return app.commandDiscovery.GetCompletions(cmdName)
}

// ExecuteCommand executes a command
func (app *Application) ExecuteCommand(ctx context.Context, cmdLine string) (string, error) {
	app.logger.Infof("Executing command: %s", cmdLine)
	return app.commandDiscovery.ExecuteWithAutoComplete(ctx, cmdLine)
}

// On registers an event listener
func (app *Application) On(eventType string, listener func(Event)) {
	app.eventBus.mu.Lock()
	defer app.eventBus.mu.Unlock()

	if _, exists := app.eventBus.listeners[eventType]; !exists {
		app.eventBus.listeners[eventType] = make([]func(Event), 0)
	}
	app.eventBus.listeners[eventType] = append(app.eventBus.listeners[eventType], listener)
}

// Emit emits an event
func (app *Application) Emit(event Event) {
	app.eventBus.mu.RLock()
	listeners, exists := app.eventBus.listeners[event.Type]
	app.eventBus.mu.RUnlock()

	if !exists {
		return
	}

	for _, listener := range listeners {
		go listener(event)
	}
}

// GetStatus returns the application status
func (app *Application) GetStatus() map[string]interface{} {
	app.mu.RLock()
	defer app.mu.RUnlock()

	return map[string]interface{}{
		"name":     app.name,
		"version":  app.version,
		"running":  app.running,
		"uptime":   app.GetUptime().String(),
		"managers": len(app.managers),
		"sessions": len(app.sessionManager.sessions),
	}
}

// SessionManager methods

// AddSession adds a session to the manager
func (sm *SessionManager) AddSession(session *UISession) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.sessions[session.ID] = session
}

// RemoveSession removes a session from the manager
func (sm *SessionManager) RemoveSession(sessionID string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	delete(sm.sessions, sessionID)
}

// GetSession retrieves a session
func (sm *SessionManager) GetSession(sessionID string) (*UISession, error) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	session, exists := sm.sessions[sessionID]
	if !exists {
		return nil, fmt.Errorf("session %s not found", sessionID)
	}
	return session, nil
}

// GetAllSessions returns all sessions
func (sm *SessionManager) GetAllSessions() map[string]*UISession {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	sessions := make(map[string]*UISession)
	for k, v := range sm.sessions {
		sessions[k] = v
	}
	return sessions
}

// EventBus methods

// Off unregisters an event listener
func (eb *EventBus) Off(eventType string) {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	delete(eb.listeners, eventType)
}

// GetListeners returns listeners for an event type
func (eb *EventBus) GetListeners(eventType string) int {
	eb.mu.RLock()
	defer eb.mu.RUnlock()

	listeners, exists := eb.listeners[eventType]
	if !exists {
		return 0
	}
	return len(listeners)
}
