## Development Guide

### Setting Up Development Environment

#### Prerequisites
- Go 1.22.3
- Git
- Docker & Docker Compose
- Make (optional)

#### Initial Setup
```bash
# Clone repository
git clone https://github.com/textolytics/nbgo.git
cd nbgo

# Install dependencies
go mod download

# Copy example configuration
cp .env.example .env

# Start development services
docker-compose up -d
```

### Module Architecture

#### Core Module (`core/`)
Provides the foundation for all provider implementations.

**Key Components:**
- `Provider` interface - Base interface for all providers
- `Registry` - Manages provider instances
- SDK-specific implementations (C, Rust, Go, Python, Flutter, Robot Framework, MCP)

#### Message Bus Module (`mb/`)
Handles asynchronous messaging between components.

**Supported Buses:**
- ZeroMQ (ZMQ)
- MQTT

#### Data Warehouse Module (`dw/`)
Manages data persistence and querying.

**Supported Systems:**
- ClickHouse
- InfluxDB
- Parquet

#### Monitoring Module (`mon/`)
Provides metrics collection and alerting.

**Supported Systems:**
- VictoriaMetrics
- InfluxDB
- Grafana

#### Gateway Module (`gw/`)
Integrates with external market data sources.

**Supported Gateways:**
- Gate.io
- Freedx

### Common Development Tasks

#### Adding a New Provider

1. Create provider struct:
```go
type MyProvider struct {
    *BaseProvider
    config map[string]interface{}
}
```

2. Implement Provider interface:
```go
func (mp *MyProvider) Initialize(ctx context.Context) error { }
func (mp *MyProvider) Start(ctx context.Context) error { }
func (mp *MyProvider) Stop(ctx context.Context) error { }
func (mp *MyProvider) IsHealthy(ctx context.Context) error { }
```

3. Register in main.go:
```go
registry.Register("my_provider", NewMyProvider())
```

#### Adding a New CLI Command

```go
cmd := &cli.Command{
    Name:        "mycommand",
    Description: "My command description",
    Usage:       "mycommand [options]",
    Handler: func(ctx context.Context, args []string) error {
        // Implementation
        return nil
    },
}
cliApp.RegisterCommand(cmd)
```

#### Adding Configuration Options

1. Update `conf/config.go` with new struct fields
2. Update `nbgo.yml` example file
3. Update `.env.example` with environment variables

### Testing Guidelines

#### Unit Tests
- Create test files with `_test.go` suffix
- Use standard Go testing package
- Aim for > 80% coverage

#### Integration Tests
- Use Robot Framework in `test/` directory
- Test end-to-end workflows

#### Running Tests
```bash
# All tests
bash test.sh

# Go tests only
go test -v ./...

# With coverage
go test -cover -coverprofile=coverage.out ./...
```

### Code Style

- Use `gofmt` for formatting
- Run `golint` for linting
- Add comments for exported functions
- Follow Go naming conventions

### Documentation

- Keep README.md updated
- Add docstrings to public APIs
- Document configuration options
- Include examples in comments

### Git Workflow

1. Create feature branch: `git checkout -b feature/my-feature`
2. Make changes and test
3. Commit: `git commit -am 'Add my feature'`
4. Push: `git push origin feature/my-feature`
5. Create Pull Request

### Performance Optimization

- Use context for cancellation
- Implement connection pooling
- Batch operations where possible
- Monitor resource usage with Grafana

### Debugging

Enable debug logging:
```bash
export LOG_LEVEL=debug
```

Use pprof for profiling:
```go
import _ "net/http/pprof"
```

### Common Issues

**Issue**: Port already in use
**Solution**: Change port in configuration or stop conflicting service

**Issue**: Database connection failed
**Solution**: Ensure Docker services are running: `docker-compose ps`

**Issue**: Go module not found
**Solution**: Run `go mod download` and ensure correct import paths
