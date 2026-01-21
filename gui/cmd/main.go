package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/textolytics/nbgo/cli"
	"github.com/textolytics/nbgo/conf"
	"github.com/textolytics/nbgo/gui"
	"github.com/textolytics/nbgo/logs"
)

func main() {
	// Parse command line flags
	mode := flag.String("mode", "tui", "UI mode: tui (terminal), cli (command-line), or settings")
	config := flag.String("config", "nbgo.yml", "Configuration file path")
	flag.Parse()

	// Create logger
	logger := logs.NewStandardLogger(os.Stdout)

	// Load configuration using config manager
	cfgMgr := conf.NewManager()
	if err := cfgMgr.LoadJSON(*config); err != nil {
		logger.Warnf("Failed to load configuration file %s, using defaults: %v", *config, err)
	}
	cfg := &conf.Config{
		Environment: make(map[string]string),
		Gateways:    make(map[string]interface{}),
	}

	// Create command registry
	cmdReg := cli.NewCommandRegistry()

	// Create application
	app := gui.NewApplication("NBGO GUI", "1.0.0", cfg, logger, cmdReg)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Initialize application
	if err := app.Initialize(ctx); err != nil {
		logger.Errorf("Failed to initialize application: %v", err)
		os.Exit(1)
	}

	// Start application
	if err := app.Start(ctx); err != nil {
		logger.Errorf("Failed to start application: %v", err)
		os.Exit(1)
	}

	// Handle signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Run based on mode
	switch *mode {
	case "cli":
		runCLIMode(app, logger)
	case "settings":
		runSettingsMode(app, logger)
	case "tui":
		fallthrough
	default:
		runTUIMode(app, logger)
	}

	// Wait for signal
	<-sigChan
	logger.Info("Shutting down...")

	// Stop application
	if err := app.Stop(ctx); err != nil {
		logger.Errorf("Error during shutdown: %v", err)
	}
}

// runCLIMode runs the CLI interactive mode
func runCLIMode(app *gui.Application, logger logs.Logger) {
	logger.Info("Starting CLI mode...")
	fmt.Println("=== NBGO CLI Interface ===")
	fmt.Println("Type 'help' for commands, 'exit' to quit")
	fmt.Println("")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("nbgo> ")
		if !scanner.Scan() {
			break
		}

		cmdLine := scanner.Text()
		if cmdLine == "exit" || cmdLine == "quit" {
			break
		}

		if cmdLine == "help" {
			printCLIHelp()
			continue
		}

		if cmdLine == "" {
			continue
		}

		// Execute command
		output, err := app.ExecuteCommand(context.Background(), cmdLine)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
		if output != "" {
			fmt.Println(output)
		}
	}
}

// runTUIMode runs the terminal UI mode
func runTUIMode(app *gui.Application, logger logs.Logger) {
	logger.Info("Starting TUI mode...")
	fmt.Println("=== NBGO Terminal UI ===")
	fmt.Println("Views available:")

	views := app.GetUIManager().ListViews()
	for name, view := range views {
		fmt.Printf("  [%s] - %s\n", name, view.Description)
	}

	fmt.Println("\nManagers available:")
	managers := app.ListManagers()
	for managerType := range managers {
		fmt.Printf("  - %s\n", managerType)
	}

	fmt.Println("\nPress Ctrl+C to exit")
}

// runSettingsMode runs the settings management mode
func runSettingsMode(app *gui.Application, logger logs.Logger) {
	logger.Info("Starting Settings mode...")
	fmt.Println("=== NBGO Settings Manager ===")

	settingsMgr := gui.NewSettingsManager(logger)
	if err := settingsMgr.LoadSchemasFromDirectory("schema"); err != nil {
		logger.Errorf("Failed to load schemas: %v", err)
		return
	}

	if err := settingsMgr.GenerateAggregatedSettings(); err != nil {
		logger.Errorf("Failed to generate aggregated settings: %v", err)
		return
	}

	if err := settingsMgr.ValidateSettings(); err != nil {
		logger.Errorf("Settings validation failed: %v", err)
		return
	}

	fmt.Println("Settings loaded and validated successfully")
	fmt.Println("\nAvailable settings:")
	settingsMgr.PrintSettings()
}

// printCLIHelp prints CLI help
func printCLIHelp() {
	help := `
Available Commands:
  start <service>       - Start a service
  stop <service>        - Stop a service
  restart <service>     - Restart a service
  status                - Show system status
  list [type]           - List resources
  config [key] [value]  - Get or set configuration
  logs [service]        - View logs
  help                  - Show this help
  exit/quit             - Exit the program
`
	fmt.Print(help)
}
