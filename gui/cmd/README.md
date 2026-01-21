# NBGO GUI CLI

Command-Line Interface for NBGO Market Data System GUI Management

## Overview

The NBGO GUI CLI provides three operational modes for managing and interacting with the NBGO system:

1. **TUI Mode** (Terminal User Interface) - Interactive graphical terminal interface
2. **CLI Mode** - Command-line interactive shell
3. **Settings Mode** - Configuration and settings management

## Building

```bash
cd /home/textolytics/nbgo
go build -o nbgo-gui ./gui/cmd
```

## Usage

### Basic Commands

```bash
# Run in TUI mode (default)
./nbgo-gui

# Run in CLI mode
./nbgo-gui -mode cli

# Run in Settings mode
./nbgo-gui -mode settings

# Load custom config
./nbgo-gui -config /path/to/config.yml -mode cli
```

### TUI Mode

Terminal User Interface mode provides an interactive dashboard:

```bash
./nbgo-gui -mode tui
```

Features:
- Multiple view navigation
- System dashboard
- Data explorer
- Debug console
- Monitoring dashboard
- Log viewer
- Configuration editor
- Environment variable display

**Navigation:**
- `Tab` - Switch between views
- `Escape` - Go back
- `:` - Command mode
- `/` - Search mode
- `Ctrl+Q` - Exit

### CLI Mode

Interactive command-line shell mode:

```bash
./nbgo-gui -mode cli
```

**Available Commands:**

```
start <service>       - Start a service
stop <service>        - Stop a service
restart <service>     - Restart a service
status                - Show system status
list [type]           - List resources
config [key] [value]  - Get or set configuration
logs [service]        - View logs
help                  - Show available commands
exit/quit             - Exit the program
```

**Example CLI Session:**

```
nbgo> status
System Status:
  Running: true
  Uptime: 2h 15m 30s
  Managers: 5
  Sessions: 1

nbgo> list modules
Available Modules:
  - system
  - database
  - api
  - messaging
  - monitoring

nbgo> config system log_level debug
Configuration updated: system.log_level = debug

nbgo> exit
```

### Settings Mode

Advanced settings management and validation:

```bash
./nbgo-gui -mode settings
```

Features:
- Load schemas from `schema/` directory
- Aggregate all settings into one tree
- Validate settings against constraints
- Edit settings values
- Save updated settings
- View setting documentation

**Schema Files:**

Settings schemas are defined in YAML or JSON format in the `schema/` directory:

```yaml
name: system
description: System Configuration
version: 1.0.0
settings:
  - name: environment
    type: string
    description: Application environment
    default: development
    required: true
    options:
      - development
      - staging
      - production
```

**Available Schema Files:**

- `schema/system.yaml` - System-wide settings
- `schema/database.yaml` - Database configuration
- `schema/api.yaml` - API settings

## Configuration

### Command Line Flags

- `-mode` - UI mode: `tui`, `cli`, or `settings` (default: `tui`)
- `-config` - Configuration file path (default: `nbgo.yml`)

### Environment Variables

```bash
export NBGO_MODE=cli
export NBGO_CONFIG=/etc/nbgo/config.yml
./nbgo-gui
```

## Settings Management

### Schema Structure

Each schema defines a collection of related settings:

```yaml
name: <section-name>
description: <human-readable description>
version: <semantic version>
settings:
  - name: <setting-name>
    type: <string|integer|boolean|array>
    description: <setting description>
    default: <default value>
    required: <true|false>
    options: [<list of valid options>]
    validate: <validation regex>
```

### Editing Settings

In CLI mode:

```
nbgo> config system environment production
Setting updated: system.environment = production
```

In Settings mode:

```
=== NBGO Settings Manager ===
Settings loaded and validated successfully

Available settings:

[system] System Configuration (v1.0.0)
  environment: Application environment (required)
    Type: string, Default: development
    Options: [development, staging, production]
  log_level: Logging level (required)
    Type: string, Default: info
    Options: [debug, info, warning, error]
```

### Validating Settings

Settings are automatically validated on load:

```
=== NBGO Settings Manager ===
Settings validation successful
```

If validation fails:

```
Settings validation found 3 errors:
  system.environment: required field is empty
  database.host: required field is empty
  api.base_url: required field is empty
```

### Saving Settings

Settings can be saved to YAML files:

```
./nbgo-gui -mode settings
# Make edits through the interface, then:
Settings saved: ./config/system.yaml
Settings saved: ./config/database.yaml
Settings saved: ./config/api.yaml
```

## Architecture

### Components

1. **UIManager** - Manages views and user interface
2. **SettingsManager** - Handles schema loading and settings
3. **CommandDiscovery** - Auto-discovers available commands
4. **KeyboardHandler** - Processes keyboard input
5. **NavigationController** - Controls view navigation
6. **Application** - Main application orchestrator

### Data Flow

```
┌──────────────┐
│ Schema Files │
└──────┬───────┘
       │
       v
┌──────────────────────┐
│ SettingsManager      │
│  - Load schemas      │
│  - Aggregate         │
│  - Validate          │
│  - Edit/Save         │
└──────┬───────────────┘
       │
       v
┌──────────────────────┐
│ Application          │
│  - Initialize        │
│  - Run modes         │
│  - Handle input      │
└──────────────────────┘
```

## Examples

### Example 1: System Configuration

```bash
./nbgo-gui -mode settings
# System Configuration Manager opens
# Edit environment, log level, port, etc.
# Settings validated and saved
```

### Example 2: Interactive CLI Session

```bash
./nbgo-gui -mode cli
nbgo> list modules
nbgo> config system log_level debug
nbgo> config database host db.example.com
nbgo> status
nbgo> exit
```

### Example 3: Custom Configuration

```bash
./nbgo-gui -config /etc/nbgo/production.yml -mode tui
# TUI opens with production configuration
```

## Troubleshooting

### Schema Files Not Loading

Ensure schema files are in the `schema/` directory with `.yaml` or `.json` extension:

```bash
ls -la schema/
# Should show: system.yaml, database.yaml, api.yaml
```

### Validation Errors

Check required settings have values:

```
system.environment: required field is empty
# Set: nbgo> config system environment production
```

### Connection Issues

Verify configuration:

```
nbgo> config
# Shows all current settings
```

## Advanced Usage

### Creating Custom Schema

Create a new file in `schema/` directory:

```yaml
# schema/cache.yaml
name: cache
description: Caching Configuration
version: 1.0.0
settings:
  - name: type
    type: string
    description: Cache backend
    default: redis
    required: true
    options:
      - redis
      - memcached
      - in-memory
  - name: ttl
    type: integer
    description: Default TTL in seconds
    default: 3600
    required: false
```

### Exporting Settings

Export current settings as JSON:

```bash
./nbgo-gui -mode settings > config_export.json
```

### Batch Configuration

Create a configuration script:

```bash
#!/bin/bash
./nbgo-gui -mode cli << EOF
config system environment production
config system log_level warning
config database host prod-db.example.com
config api rate_limit 1000
exit
EOF
```

## Best Practices

1. **Backup Before Changes**: Save current settings before major changes
2. **Validate Settings**: Always run settings validation mode before deploying
3. **Document Custom Schemas**: Add clear descriptions in schema files
4. **Use Required Flags**: Mark critical settings as required
5. **Provide Defaults**: Always include sensible defaults
6. **Test Configuration**: Test settings in development before production

## Performance Considerations

- Schema loading is lazy - only loaded when needed
- Settings validation is cached
- Changes don't apply until explicitly saved
- TUI mode uses virtual rendering for efficiency

## Security Considerations

- Sensitive settings (passwords) are not logged
- Configuration files should have restricted permissions (600)
- Use environment variables for secrets in production
- Validate all user input before applying

## Support

For issues or questions:
- Check schema validation errors
- Review configuration files
- Verify all required settings are set
- Check application logs

## License

MIT License - See LICENSE file for details
