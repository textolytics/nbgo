# NBGO Configuration Manager System

A comprehensive configuration management system for the NBGO Market Data System, providing schema discovery, annotation generation, build configuration management, and automated installation.

## Components

### 1. Schema Module (`schema/schema.go`)

Discovers and manages provider configuration schemas.

**Key Types:**
- `Schema` - Configuration schema definition
- `Field` - Schema field definition  
- `Annotation` - Provider configuration annotation
- `BuildConfiguration` - Build config for providers
- `BuildTarget` - Build target (OS/Arch combination)
- `ProviderMetadata` - Provider metadata with schema

**Key Classes:**
- `Repository` - Schema storage and management
- `Discovery` - Discovers schemas from provider documentation

**Features:**
- Auto-discover provider schemas
- Validate data against schemas
- Export/import schema definitions
- Multi-provider schema management

### 2. Configuration Manager Module (`conf/manager.go`)

Manages application configurations with validation and monitoring.

**Key Types:**
- `ConfigValidator` - Configuration validator interface
- `ConfigWatcher` - Configuration change watcher interface
- `Monitor` - Configuration file monitor

**Key Classes:**
- `ConfigManager` - Central configuration manager
- `Monitor` - File-based configuration monitoring

**Features:**
- Set/get configuration values with validation
- YAML/JSON import/export
- Configuration backups and restoration
- Configuration change notifications
- File-based monitoring with auto-reload

### 3. Installation Manager Module (`run/installer.go`)

Manages installation of providers and services with dependency resolution.

**Key Types:**
- `InstallationTarget` - Installation target definition
- `InstallationResult` - Installation result tracking
- `InstallerFunc` - Installer function type
- `HealthCheck` - Health check definition

**Key Classes:**
- `InstallationManager` - Central installation management
- `HealthCheckMonitor` - Health check monitoring

**Features:**
- Register installation targets
- Dependency resolution and ordering
- Install single or all targets
- Health check integration
- Installation result tracking
- Uninstall with dependency checking

## Usage Examples

### Schema Discovery

```go
// Create discovery instance
discovery := schema.NewDiscovery(repo, []string{"provider1", "provider2"})

// Scan providers
result, err := discovery.ScanProviders(ctx)

// List discovered schemas
schemas := repo.ListSchemas(schema.SchemaTypeProvider)

// Validate data
valid, errors := repo.ValidateSchemaData("schema_id", data)
```

### Configuration Management

```go
// Create config manager
configMgr := config.NewConfigManager("config.json")

// Set configuration
configMgr.Set("app.name", "nbgo")
configMgr.Set("app.debug", true)

// Get configuration
name, exists := configMgr.GetString("app.name")

// Save to file
configMgr.SaveJSON("config.json")

// Load from file
configMgr.LoadJSON("config.json")

// Create backup
configMgr.Backup()
```

### Installation Management

```go
// Create installation manager
installMgr := run.NewInstallationManager()

// Register installer
installMgr.RegisterInstaller("golang", run.CommandInstaller("go", "version"))

// Check installation plan
plan := installMgr.GetInstallationPlan()

// Install target
result, err := installMgr.Install("golang")

// Install all
results, err := installMgr.InstallAll()

// Health checks
monitor := run.NewHealthCheckMonitor(5 * time.Second)
monitor.RegisterCheck(&run.HealthCheck{
    TargetID: "golang",
    CheckFunc: func(ctx context.Context) error {
        return exec.CommandContext(ctx, "go", "version").Run()
    },
})
monitor.Start(ctx)
```

## API Endpoints

### Schema Endpoints
- `GET /api/v1/schemas` - List all schemas
- `GET /api/v1/schemas/{id}` - Get schema details
- `POST /api/v1/schemas/{id}/validate` - Validate data

### Annotation Endpoints
- `GET /api/v1/annotations` - List annotations
- `GET /api/v1/annotations/{provider}` - Get provider annotations
- `PATCH /api/v1/annotations/{id}` - Update annotation

### Build Endpoints
- `GET /api/v1/build/targets` - List build targets
- `GET /api/v1/build/configs` - List build configurations
- `GET /api/v1/build/configs/{id}/commands` - Get build commands

### Installation Endpoints
- `GET /api/v1/install/targets` - List installation targets
- `GET /api/v1/install/plan` - Get installation plan
- `POST /api/v1/install/targets/{id}` - Install target

## CLI Commands

### Schema Commands
```bash
nbgo schema:list                      # List all schemas
nbgo schema:show <schema-id>          # Show schema details
nbgo schema:validate <schema-id> <json>  # Validate data
```

### Annotation Commands
```bash
nbgo annotation:list                  # List annotations
nbgo annotation:list --provider NAME  # List provider annotations
nbgo annotation:update <id> <settings> # Update annotation
```

### Build Commands
```bash
nbgo build:targets                    # List build targets
nbgo build:config <config-id>         # Show build configuration
nbgo build:commands <config-id>       # Get build commands
```

### Installation Commands
```bash
nbgo install:targets                  # List installation targets
nbgo install:plan                     # Show installation plan
nbgo install:run <target-id>          # Install specific target
nbgo install:all                      # Install all targets
```

## Configuration Files

### Configuration Schema
```yaml
schemas:
  provider:
    - id: schema_provider_name
      name: Provider Name
      version: "1.0"
      fields:
        - name: api_key
          type: string
          required: true
          description: API key for authentication
        - name: endpoint
          type: string
          required: false
          default: https://api.example.com

annotations:
  - provider: CProvider
    settings:
      compiler: gcc
      standard: c99
      optimization: -O2
    environment:
      CC: gcc
      CFLAGS: -Wall -O2

build:
  targets:
    - id: linux_amd64
      name: Linux x86_64
      os: linux
      arch: amd64
      enabled: true
    - id: darwin_arm64
      name: macOS ARM64
      os: darwin
      arch: arm64
      enabled: true
  
  configurations:
    - provider: nbgo
      version: "1.0.0"
      targets:
        - linux_amd64
        - darwin_amd64
        - windows_amd64
      flags:
        - -ldflags
        - -w -s
        - -trimpath

installation:
  targets:
    - id: golang
      name: Go SDK
      version: "1.22.3"
      dependencies: []
      type: tool
      enabled: true
    - id: nbgo
      name: NBGO Application
      version: "1.0.0"
      dependencies:
        - golang
        - postgres
        - redis
      type: provider
      enabled: true
```

## Dependencies

### Go Modules
- `github.com/rs/zerolog` - Structured logging
- Standard library packages

### External Services
- Provider documentation (for schema discovery)
- Build tools (gcc, rustc, go, python)
- Services (PostgreSQL, Redis, ClickHouse, etc.)

## Integration with NBGO

The Configuration Manager integrates with:

1. **Core Modules** - Provider management and initialization
2. **Message Bus** - Configuration changes propagation
3. **Monitoring** - Health checks and status tracking
4. **CLI** - Command-line interface for management
5. **API** - RESTful endpoints for remote management
6. **Docker** - Containerized deployment configuration

## Development

### Adding New Providers

1. Add provider to schema discovery
```go
providers := []string{"NewProvider"}
discovery.ScanProviders(ctx)
```

2. Generate annotation
```go
annotation := annotator.generate_for_provider("NewProvider")
```

3. Create build configuration
```go
config := buildMgr.create_build_config("NewProvider", "1.0.0", targetIds)
```

4. Register installation target
```go
target := &InstallationTarget{
    ID: "new_provider",
    Name: "New Provider",
    Dependencies: []string{"golang"},
}
installMgr.RegisterTarget(target)
```

### Custom Validators

```go
type CustomValidator struct{}

func (cv *CustomValidator) Validate(config interface{}) error {
    // Validation logic
    return nil
}

configMgr.RegisterValidator("custom_key", &CustomValidator{})
```

### Custom Watchers

```go
type CustomWatcher struct{}

func (cw *CustomWatcher) OnChange(key string, oldValue, newValue interface{}) error {
    log.Printf("Config changed: %s = %v", key, newValue)
    return nil
}

configMgr.RegisterWatcher("app.name", &CustomWatcher{})
```

## Testing

### Schema Discovery Tests
```bash
go test ./schema -v
```

### Configuration Manager Tests
```bash
go test ./conf -v
```

### Installation Manager Tests
```bash
go test ./run -v
```

## Performance Considerations

- Schema discovery is cached in memory
- Configuration changes are validated before storing
- Installation dependency resolution uses topological sort
- Health checks run asynchronously with configurable intervals
- File monitoring uses efficient file change detection

## Security

- Configuration values support encrypted storage
- Installation targets require dependency verification
- Schema validation prevents invalid configurations
- File backups are created before modifications
- Access control can be implemented via API authentication

## Future Enhancements

1. **Schema Versioning** - Support multiple schema versions
2. **Configuration Encryption** - Encrypt sensitive configuration values
3. **Audit Logging** - Track all configuration changes
4. **Remote Configuration** - Fetch configurations from remote servers
5. **Configuration Templates** - Pre-defined configurations for common setups
6. **Interactive TUI** - Terminal user interface for configuration management
7. **Configuration Validation Rules** - Custom validation rule engine
8. **Configuration History** - Track configuration changes over time

## References

- [NBGO Core Modules](../core/README.md)
- [Message Bus System](../mb/README.md)
- [Data Warehouse System](../dw/README.md)
- [Monitoring System](../mon/README.md)
- [API Endpoints](../api/README.md)
