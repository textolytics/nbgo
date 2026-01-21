# Building NBGO GUI

## Quick Build

```bash
cd /home/textolytics/nbgo
cd gui/cmd && go build -o ../../nbgo-gui . && cd ../..
```

Or use this shorthand:

```bash
cd /home/textolytics/nbgo
(cd gui/cmd && go build -o ../../nbgo-gui .)
```

## Running the GUI

```bash
# Terminal UI mode (default)
./nbgo-gui

# CLI interactive mode
./nbgo-gui -mode cli

# Settings management mode
./nbgo-gui -mode settings

# With custom config file
./nbgo-gui -config /path/to/config.yml
```

## Verify All Packages Build

```bash
go build ./...
```

## Notes

- The GUI executable is built from the `gui/cmd` package
- The gui/cmd package imports from the gui module and other dependencies
- All configuration is optional - defaults are used if not provided
- The config file path defaults to `nbgo.yml` in the current directory

## Module Structure

```
nbgo/
├── gui/                  # GUI module
│   ├── cmd/             # CLI executable package
│   │   └── main.go      # Entry point
│   ├── ui.go            # UI Manager
│   ├── view.go          # View definitions
│   ├── session.go       # Session management
│   ├── managers.go      # Manager providers
│   ├── command_discovery.go
│   ├── keyboard_navigation.go
│   ├── application.go   # Main application
│   ├── settings.go      # Settings management
│   └── cmd.go           # CLI runner
├── cli/                 # CLI module
├── conf/                # Configuration module
├── logs/                # Logging module
└── go.mod              # Root module file
```
