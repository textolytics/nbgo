# NBGO GUI Module

Market Data System - GUI Manager Provider UI

## Overview

The GUI module provides a comprehensive user interface for managing the NBGO Market Data System. It includes multiple manager UIs for different aspects of system administration, real-time monitoring, and interactive command execution.

## Features

### Core Features

- **Multi-View Dashboard**: Solarized-themed UI with multiple viewport modes
- **Auto-Discovery**: Automatic command discovery from module sources
- **Command Suggestions & Auto-Completion**: Intelligent command suggestions and completions
- **Keyboard Navigation**: Arrow-key navigation with modal support
- **Multi-Window Sessions**: Multiple concurrent UI sessions with environment isolation
- **Theme Support**: Modern Solarized highlighting with customizable colors

### Manager UIs

#### 1. Install Prerequisites Manager
Manages system prerequisites and dependencies:
- Prerequisites discovery and validation
- Dependency checking and installation
- Version management and compatibility verification

#### 2. Environment Manager
Manages system environment initialization:
- Environment discovery and validation
- Environment setup and configuration
- Purge and cleanup operations
- Environment variable management

#### 3. Configure Manager
Manages system configuration:
- Configuration discovery and validation
- Configuration reading and updating
- Schema-based configuration management
- Configuration saving and persistence

#### 4. Build Manager
Manages build operations:
- Build task discovery
- Module building and compilation
- Clean and rebuild operations
- Build validation

#### 5. Install Manager
Manages installation operations:
- Installation task discovery
- Component installation
- Installation validation
- Update operations

### Views

The system provides multiple views for different purposes:

- **Dashboard**: Main system dashboard
- **Data Explorer**: Browse market data and instruments
- **Debug Console**: Debug information and diagnostics
- **CLI Console**: Command-line interface
- **Terminal**: System terminal access
- **API Explorer**: API endpoints and testing
- **Monitoring**: System monitoring and metrics
- **Logs**: System logs and filtering
- **Environment**: Environment variables viewer
- **Configuration**: System configuration viewer

## Architecture

### Component Structure

```
gui/
├── go.mod                      # Go module definition
├── ui.go                       # Core UI Manager
├── view.go                     # View definitions
├── session.go                  # Session and Window management
├── managers.go                 # Manager provider implementations
├── command_discovery.go        # Command discovery and auto-completion
├── keyboard_navigation.go      # Keyboard navigation and key bindings
├── application.go              # Main application
└── README.md                   # This file
```

### Key Classes

#### UIManager
Central manager for all UI operations:
- View management and switching
- Provider registration
- Session creation and management
- Command history tracking
- Theme management

#### View
Represents a UI view/panel:
- Content rendering
- Status tracking
- Data visualization
- Command execution

#### UISession
Manages a UI session with multiple windows:
- Window management
- Environment variable tracking
- Command buffering
- Multi-window coordination

#### ManagerProvider
Interface for management providers:
- CommandDiscovery
- Task execution
- Configuration management
- Status reporting

#### CommandDiscovery
Automatic command discovery system:
- Module scanning
- Command registration
- Suggestion generation
- Auto-completion support

#### KeyboardHandler
Handles keyboard input:
- Key binding registration
- Event routing
- Navigation stack management
- View focus management

#### NavigationController
Controls navigation between views:
- View navigation
- History management
- Mode switching (normal, insert, command, search)
- Back/forward navigation

## Usage

### Basic Setup

```go
import (
    "context"
    "gui"
    "conf"
    "logs"
    "cli"
)

// Create configuration and logger
config := &conf.Config{}
logger := logs.NewLogger()
cmdReg := cli.NewCommandRegistry()

// Create and initialize application
app := gui.NewApplication("NBGO", "1.0.0", config, logger, cmdReg)
ctx := context.Background()

if err := app.Initialize(ctx); err != nil {
    logger.Errorf("Failed to initialize: %v", err)
    return
}

// Start the application
if err := app.Start(ctx); err != nil {
    logger.Errorf("Failed to start: %v", err)
    return
}

// Stop when done
defer app.Stop(ctx)
```

### Creating Custom Views

```go
// Create a custom view
view := gui.NewView("custom", "Custom View", "My custom view")
view.SetDimensions(0, 0, 80, 24)
view.UpdateContent("View content here")

// Add to UI Manager
uiManager := app.GetUIManager()
uiManager.views["custom"] = view
```

### Using Managers

```go
// Get a specific manager
prerequisitesMgr, err := app.GetManager(gui.ManagerTypeInstallPrerequisites)
if err != nil {
    return err
}

// Discover tasks
tasks, err := prerequisitesMgr.Discover(ctx)
if err != nil {
    return err
}

// Execute operation
err = prerequisitesMgr.Execute(ctx, "install_go", map[string]interface{}{
    "version": "1.22.3",
})
```

### Command Discovery

```go
// Discover commands from modules
modulePaths := []string{"../core", "../conf", "../cli"}
if err := app.DiscoverCommands(ctx, modulePaths); err != nil {
    logger.Errorf("Discovery failed: %v", err)
}

// Get suggestions
suggestions := app.GetCommandSuggestions("start")

// Get completions
completions := app.GetCommandCompletions("start")

// Execute command
output, err := app.ExecuteCommand(ctx, "start --daemon")
```

### Keyboard Navigation

```go
// Create navigation
navCtrl := gui.NewNavigationController()

// Navigate
navCtrl.Navigate("dashboard")
navCtrl.Navigate("data")

// Go back
navCtrl.GoBack()

// Get current view
currentView := navCtrl.GetCurrentView()
```

### Session Management

```go
// Create session
session := app.uiManager.CreateSession("session1")

// Add window to session
window := gui.NewWindow("window1", "Data Window", view)
session.AddWindow(window)

// Set environment variables
session.SetEnvironmentVariable("NBGO_MODE", "development")

// Get environment
env := session.GetAllEnvironmentVariables()
```

## Configuration

### Theme Configuration

```go
// Create custom theme
theme := &gui.UITheme{
    Name:      "custom",
    Solarized: true,
    Fancy:     true,
    Borders:   true,
    Colors: map[string]string{
        "base03": "#002b36",
        // ... more colors
    },
}

uiManager.SetTheme(theme)
```

### Manager Configuration

Managers are configured through the main configuration file and environment variables:

```yaml
# In nbgo.yml
gui:
  theme: solarized-dark
  auto_discovery: true
  keyboard_navigation: true
  max_history: 1000
  max_log_entries: 10000

managers:
  install_prerequisites:
    enabled: true
  environment:
    enabled: true
  configure:
    enabled: true
  build:
    enabled: true
  install:
    enabled: true
```

## Event System

The application includes an event bus for inter-component communication:

```go
// Register event listener
app.On("command_executed", func(event gui.Event) {
    fmt.Printf("Command executed: %v\n", event.Data)
})

// Emit event
app.Emit(gui.Event{
    Type:      "command_executed",
    Timestamp: time.Now().Unix(),
    Data:      map[string]interface{}{"cmd": "start"},
    Source:    "cli",
})
```

## Keyboard Bindings

Default keyboard bindings:

| Key | Action |
|-----|--------|
| `Tab` | Switch to next view |
| `Escape` | Go back to previous view |
| `:` | Enter command mode |
| `/` | Enter search mode |
| `Ctrl+Q` | Quit application |
| `Arrow Keys` | Navigate within view |
| `Enter` | Execute command/selection |

## Best Practices

1. **Always initialize before starting**: Call `Initialize()` before `Start()`
2. **Use contexts for cancellation**: Pass context to enable graceful shutdown
3. **Handle errors appropriately**: Log and propagate errors correctly
4. **Session isolation**: Use separate sessions for different UI instances
5. **Command discovery**: Scan modules for available commands on startup
6. **Theme consistency**: Use consistent theme throughout application

## Performance Considerations

- Command discovery is asynchronous
- Views are rendered only when visible
- Sessions are isolated with independent environments
- Event listeners run in separate goroutines
- History is limited to configurable size

## Troubleshooting

### Commands not discovered
- Ensure module paths are correct
- Check module structure follows conventions
- Verify commands have proper metadata

### UI not responsive
- Check keyboard event handler registration
- Verify event loop is running
- Ensure goroutines aren't blocked

### Theme not applied
- Verify theme colors are valid
- Check terminal supports Solarized colors
- Reset theme to default and retry

## Future Enhancements

- [ ] Web-based UI support
- [ ] REST API for remote access
- [ ] Recording/playback of sessions
- [ ] Custom scripting language
- [ ] Advanced debugging tools
- [ ] Performance profiling UI
- [ ] Log aggregation and filtering
- [ ] Real-time collaboration

## Contributing

When extending the GUI module:

1. Follow existing code patterns
2. Add proper error handling
3. Include context for cancellation
4. Document new features
5. Add tests for new functionality

## License

MIT License - See LICENSE file for details
