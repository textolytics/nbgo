# NBGO Configuration Manager - Quick Start Guide

## Quick Overview

The NBGO Configuration Manager provides three integrated systems:
1. **Schema Discovery** - Automatically discover provider configuration schemas
2. **Annotation Management** - Generate and manage provider configurations
3. **Installation Manager** - Automated installation with dependency resolution

All accessible via **Go API**, **REST API**, **CLI**, or **Jupyter Notebook**.

## Installation

### Prerequisites
- Go 1.22+
- Python 3.11+ (for Jupyter notebook)

### Build the Project

```bash
cd /home/textolytics/nbgo
go mod download
go build -o nbgo ./

# Or build for multiple platforms
./build.sh
```

## Quick Start Examples

### 1. Schema Discovery

#### Via CLI
```bash
# List all discovered schemas
./nbgo schema:list

# Show schema details
./nbgo schema:show schema_CProvider

# Validate data against schema
./nbgo schema:validate schema_CProvider '{"name":"test","version":"1.0"}'
```

#### Via Go API
```go
import (
    "github.com/textolytics/nbgo/schema"
    "context"
)

func main() {
    repo := schema.NewRepository()
    discovery := schema.NewDiscovery(repo, []string{
        "CProvider", "RustProvider", "GoProvider", "PythonProvider",
    })
    
    result, err := discovery.ScanProviders(context.Background())
    if err != nil {
        log.Fatal(err)
    }
    
    log.Printf("Discovered %d schemas", result.SchemasGenerated)
}
```

#### Via REST API
```bash
# Get all schemas
curl http://localhost:8080/api/v1/schemas

# Get schema details
curl http://localhost:8080/api/v1/schemas/schema_CProvider

# Validate data
curl -X POST http://localhost:8080/api/v1/schemas/schema_CProvider/validate \
  -H "Content-Type: application/json" \
  -d '{"name":"test","version":"1.0"}'
```

#### Via Jupyter Notebook
```python
# Run ConfigurationManager.ipynb
# - Schema Discovery system section demonstrates provider scanning
# - Automatic validation of provider schemas
# - Export of discovered schemas
```

### 2. Annotation Management

#### Via CLI
```bash
# List all annotations
./nbgo annotation:list

# List annotations for specific provider
./nbgo annotation:list --provider GoProvider

# Update annotation
./nbgo annotation:update annotation_CProvider '{"compiler":"clang"}'
```

#### Via Go API
```go
import "github.com/textolytics/nbgo/conf"

func main() {
    configMgr := conf.NewConfigManager("config.json")
    
    // Set provider configuration
    configMgr.Set("providers.golang.version", "1.22.3")
    configMgr.Set("providers.python.version", "3.11")
    
    // Get configuration
    version, _ := configMgr.GetString("providers.golang.version")
    
    // Save configuration
    configMgr.SaveJSON("config.json")
}
```

#### Via REST API
```bash
# Get all annotations
curl http://localhost:8080/api/v1/annotations

# Get annotations for provider
curl http://localhost:8080/api/v1/annotations?provider=GoProvider

# Update annotation
curl -X PATCH http://localhost:8080/api/v1/annotations/annotation_CProvider \
  -H "Content-Type: application/json" \
  -d '{"settings":{"compiler":"clang"}}'
```

### 3. Build Configuration

#### Via CLI
```bash
# List build targets
./nbgo build:targets

# Show build configuration
./nbgo build:config build_config_nbgo_1.0.0

# Get build commands
./nbgo build:commands build_config_nbgo_1.0.0
```

#### Via Go API
```go
import "github.com/textolytics/nbgo/build"

func main() {
    buildMgr := build.NewBuilder()
    
    // Get build targets
    targets := buildMgr.ListTargets()
    
    // Create build for provider
    config := buildMgr.CreateBuildConfig(
        "nbgo", 
        "1.0.0",
        []string{"linux_amd64", "darwin_amd64", "windows_amd64"},
    )
    
    // Get build commands
    commands := buildMgr.GetBuildCommands(config.ID)
    for _, cmd := range commands {
        log.Printf("Build: %s\n%s", cmd.Target, cmd.Command)
    }
}
```

#### Via REST API
```bash
# List available build targets
curl http://localhost:8080/api/v1/build/targets

# List build configurations
curl http://localhost:8080/api/v1/build/configs

# Get build commands
curl http://localhost:8080/api/v1/build/configs/build_config_nbgo_1.0.0/commands
```

### 4. Installation Management

#### Via CLI
```bash
# List installation targets
./nbgo install:targets

# Show installation plan
./nbgo install:plan

# Install specific target
./nbgo install:run golang

# Install all targets (with dependency resolution)
./nbgo install:all
```

#### Via Go API
```go
import "github.com/textolytics/nbgo/run"

func main() {
    installMgr := run.NewInstallationManager()
    
    // Register installer for target
    installMgr.RegisterInstaller("golang", 
        run.CommandInstaller("go", "version"),
    )
    
    // Get installation plan
    plan := installMgr.GetInstallationPlan()
    for _, step := range plan.Steps {
        log.Printf("Step %d: %s", step.Order, step.Name)
    }
    
    // Install all targets (respects dependency order)
    results, err := installMgr.InstallAll(context.Background())
    for _, result := range results {
        if result.Success {
            log.Printf("✓ %s installed", result.TargetID)
        } else {
            log.Printf("✗ %s failed: %s", result.TargetID, result.Error)
        }
    }
}
```

#### Via REST API
```bash
# Get installation targets
curl http://localhost:8080/api/v1/install/targets

# Get installation plan
curl http://localhost:8080/api/v1/install/plan

# Install specific target
curl -X POST http://localhost:8080/api/v1/install/targets/golang/install

# Get installation status/results
curl http://localhost:8080/api/v1/install/results
```

## Complete Workflow Example

### Scenario: Deploy NBGO System

```bash
# Step 1: Discover schemas from provider documentation
./nbgo schema:list
# Output: CProvider, RustProvider, GoProvider, PythonProvider, GateGateway, FreedxGateway

# Step 2: View installation plan
./nbgo install:plan
# Output:
#   1. golang (Go SDK)
#   2. python (Python SDK)
#   3. postgres (PostgreSQL)
#   4. redis (Redis)
#   5. clickhouse (ClickHouse) [depends on postgres]
#   6. prometheus (Prometheus)
#   7. grafana (Grafana) [depends on prometheus]
#   8. nbgo (NBGO Application) [depends on golang, postgres, redis]

# Step 3: Execute installation (respects dependency order)
./nbgo install:all
# Output:
#   ✓ golang installed (0.234s)
#   ✓ python installed (0.156s)
#   ✓ postgres installed (1.234s)
#   ✓ redis installed (0.567s)
#   ✓ clickhouse installed (2.341s)
#   ✓ prometheus installed (0.789s)
#   ✓ grafana installed (1.023s)
#   ✓ nbgo installed (1.456s)
#   Summary: 8/8 targets installed successfully

# Step 4: View configuration
./nbgo annotation:list
# Shows all provider configurations

# Step 5: Build for multiple platforms
./nbgo build:commands build_config_nbgo_1.0.0
# Output:
#   linux_amd64: GOOS=linux GOARCH=amd64 go build ...
#   darwin_amd64: GOOS=darwin GOARCH=amd64 go build ...
#   windows_amd64: GOOS=windows GOARCH=amd64 go build ...
```

## Configuration File Examples

### config.json (Provider Configuration)
```json
{
  "providers": {
    "golang": {
      "version": "1.22.3",
      "modules": true,
      "cgo": true
    },
    "python": {
      "version": "3.11",
      "use_venv": true
    },
    "cprovider": {
      "compiler": "gcc",
      "standard": "c99",
      "optimization": "-O2"
    }
  },
  "build": {
    "targets": ["linux_amd64", "darwin_amd64", "windows_amd64"],
    "flags": ["-ldflags", "-w -s", "-trimpath"]
  },
  "installation": {
    "all_targets": true,
    "stop_on_error": false
  }
}
```

### Installation Plan (YAML)
```yaml
installation:
  targets:
    - id: golang
      name: Go SDK
      version: "1.22.3"
      dependencies: []
    
    - id: postgres
      name: PostgreSQL
      version: "15"
      dependencies: []
    
    - id: nbgo
      name: NBGO Application
      version: "1.0.0"
      dependencies:
        - golang
        - postgres
```

## Interactive Development with Jupyter Notebook

```bash
# Start Jupyter
jupyter notebook ConfigurationManager.ipynb

# In notebook:
# 1. Schema Discovery section - discover providers automatically
# 2. Annotation Generation - generate provider configurations
# 3. Build Configuration - set up multi-platform builds
# 4. Installation Planning - see dependency order
# 5. API Testing - test REST endpoints
# 6. CLI Testing - test CLI commands
# 7. Workflow Demonstration - see complete end-to-end scenario
```

## Monitoring & Health Checks

```go
import (
    "github.com/textolytics/nbgo/run"
    "context"
    "os/exec"
)

func main() {
    // Create health check monitor
    monitor := run.NewHealthCheckMonitor(5 * time.Second)
    
    // Register health check for each component
    monitor.RegisterCheck(&run.HealthCheck{
        TargetID: "golang",
        Name: "Go SDK Health",
        CheckFunc: func(ctx context.Context) error {
            return exec.CommandContext(ctx, "go", "version").Run()
        },
    })
    
    // Start monitoring
    monitor.Start(context.Background())
    
    // Check results
    results := monitor.GetResults()
    for targetID, err := range results {
        if err != nil {
            log.Printf("✗ %s failed: %v", targetID, err)
        } else {
            log.Printf("✓ %s healthy", targetID)
        }
    }
}
```

## Troubleshooting

### Schema Discovery Not Finding Providers
```bash
# Check if providers are registered
./nbgo schema:list

# Verify provider documentation is accessible
ls -la /docs/providers/

# Enable debug logging
NBGO_DEBUG=1 ./nbgo schema:list
```

### Installation Failing Due to Dependencies
```bash
# View installation plan to see dependency order
./nbgo install:plan

# Install specific dependencies first
./nbgo install:run golang
./nbgo install:run postgres

# Then install dependent target
./nbgo install:run nbgo
```

### Configuration Validation Errors
```bash
# Check configuration schema
./nbgo schema:show schema_CProvider

# Validate configuration data
./nbgo schema:validate schema_CProvider '{"name":"test"}'

# View validation errors for specific field
curl -X POST http://localhost:8080/api/v1/schemas/schema_CProvider/validate \
  -H "Content-Type: application/json" \
  -d '{"name":"test"}'
```

### Build Command Failures
```bash
# Check build configuration
./nbgo build:config build_config_nbgo_1.0.0

# Get build commands for debugging
./nbgo build:commands build_config_nbgo_1.0.0

# Try manual build
GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -trimpath
```

## Performance Tips

1. **Cache Schema Discovery Results**
   - Discovery is run once and cached in memory
   - Use repository queries for fast lookups

2. **Batch Configuration Updates**
   - Update multiple configs before saving to file
   - Reduce I/O operations

3. **Parallel Installation**
   - Non-dependent targets can be installed in parallel
   - Dependency-respecting topological sort is automatic

4. **Monitor Health Checks**
   - Run health checks asynchronously
   - Configurable check intervals (default: 5 seconds)

## Security Best Practices

1. **Protect Configuration Files**
   ```bash
   chmod 600 config.json
   chmod 600 .env
   ```

2. **Validate All Input**
   - Schema validation is automatic
   - Use registered validators for custom rules

3. **Audit Installation Changes**
   - All installations are logged
   - Installation results are tracked with timestamps

4. **Backup Before Changes**
   - Automatic backups before configuration saves
   - Manually restore if needed: `./nbgo config:restore backup_file`

## Advanced Usage

### Custom Validators
```go
type MyValidator struct{}

func (m *MyValidator) Validate(config interface{}) error {
    // Custom validation logic
    if config == nil {
        return errors.New("config cannot be nil")
    }
    return nil
}

configMgr.RegisterValidator("custom_key", &MyValidator{})
```

### Custom Installers
```go
customInstaller := func(ctx context.Context, target *run.InstallationTarget) error {
    // Custom installation logic
    log.Printf("Installing %s...", target.Name)
    return nil
}

installMgr.RegisterInstaller("custom_target", customInstaller)
```

### Configuration Watchers
```go
type MyWatcher struct{}

func (m *MyWatcher) OnChange(key string, oldValue, newValue interface{}) error {
    log.Printf("Config %s changed from %v to %v", key, oldValue, newValue)
    return nil
}

configMgr.RegisterWatcher("app.debug", &MyWatcher{})
```

## Next Steps

1. **Explore the Jupyter Notebook** - Run `ConfigurationManager.ipynb` for interactive demonstrations
2. **Build the System** - Execute `./build.sh` to build for all platforms
3. **Run Full Installation** - Use `./nbgo install:all` to deploy complete system
4. **Configure Monitoring** - Set up health checks and monitoring dashboards
5. **Integrate with CI/CD** - Use CLI commands in deployment pipelines

## Support

For detailed documentation:
- See `CONFIGURATION_MANAGER.md` for complete API reference
- See `CONFIGURATION_MANAGER_SUMMARY.md` for implementation details
- Run `./nbgo --help` for CLI help

## References

- **Schema Module**: `schema/schema.go`
- **Configuration Module**: `conf/manager.go`
- **Installation Module**: `run/installer.go`
- **Interactive Demo**: `ConfigurationManager.ipynb`
- **Full Documentation**: `CONFIGURATION_MANAGER.md`
