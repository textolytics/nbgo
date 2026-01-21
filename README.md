# NBGO Market Data System

Market Data System - NB Go Trading Platform

## Overview

NBGO is a comprehensive market data system built in Go with support for multiple SDKs, message buses, data warehouses, and monitoring solutions.

## Project Structure

```
nbgo/
├── api/                  # API server and handlers
├── build/                # Build utilities and scripts
├── cli/                  # Command-line interface
├── conf/                 # Configuration management
├── core/                 # Core provider interfaces and implementations
├── doc/                  # Documentation generation
├── dw/                   # Data warehouse implementations (ClickHouse, InfluxDB, Parquet)
├── gw/                   # Gateway implementations (Gate.io, Freedx)
├── gui/                  # GUI management interface
│   ├── cmd/              # GUI CLI command entry point
│   ├── ui.go             # UI manager
│   ├── view.go           # View definitions
│   ├── session.go        # Session management
│   ├── managers.go       # Management providers
│   ├── settings.go       # Settings management
│   ├── command_discovery.go  # Auto-discovery
│   ├── keyboard_navigation.go # Input handling
│   └── application.go    # Application orchestrator
├── logs/                 # Logging system
├── mb/                   # Message bus implementations (ZMQ, MQTT)
├── mcp/                  # Model Context Protocol server
├── mon/                  # Monitoring implementations (VictoriaMetrics, InfluxDB, Grafana)
├── run/                  # Runtime management
├── schema/               # Configuration schemas (system, database, api)
├── task/                 # Task execution and scheduling
├── test/                 # Test utilities
├── main.go               # Application entry point
├── go.mod                # Go module definition
├── docker-compose.yml    # Docker Compose configuration
├── Dockerfile            # Docker container definition
└── nbgo.yml              # Application configuration
```

## Features

### SDK Providers
- **C**: ZeroMQ C SDK support
- **Rust**: Rust SDK support via rust-sdk
- **Go**: Go SDK support with Gin web framework
- **Python**: Python SDK support
- **Robot Framework**: Test automation framework
- **Flutter**: Mobile app framework
- **MCP**: Model Context Protocol support

### Message Bus
- **ZeroMQ (ZMQ)**: High-performance asynchronous messaging
- **MQTT**: Publish-subscribe messaging protocol

### Data Warehousing
- **ClickHouse**: Distributed column-oriented database
- **InfluxDB**: Time-series database
- **Parquet**: Apache Parquet file format support

### Monitoring
- **VictoriaMetrics**: Monitoring and alerting
- **InfluxDB**: Time-series metrics
- **Grafana**: Visualization and dashboards

### Market Data Gateways
- **Gate.io**: Cryptocurrency exchange integration
- **Freedx**: Futures exchange integration

## Getting Started

### Prerequisites
- Go 1.22.3 or later
- Docker and Docker Compose
- Python 3.11+ (for testing)

### Installation

1. Clone the repository:
```bash
git clone https://github.com/textolytics/nbgo.git
cd nbgo
```

2. Install dependencies:
```bash
go mod download
```

3. Set up environment:
```bash
cp .env.example .env
```

4. Start services with Docker:
```bash
docker-compose up -d
```

### Building

Build for multiple platforms:
```bash
bash build.sh
```

### Running Tests

```bash
bash test.sh
```

Or run Go tests directly:
```bash
go test -v ./...
```

### Configuration

Edit `nbgo.yml` to configure:
- SDK providers
- Message bus endpoints
- Data warehouse connections
- Monitoring systems
- Market data gateways

## API Usage

The system provides a REST API for interacting with all components.

### Health Check
```bash
curl http://localhost:8080/health
```

### List Providers
```bash
curl http://localhost:8080/api/v1/providers
```

## CLI Commands

- `version` - Display version information
- `health` - Check system health
- `list` - List available providers
- `help` - Display help information

## GUI Management Interface

NBGO includes a comprehensive graphical user interface for system management with three operational modes:

### GUI Modes

**Terminal UI Mode** (default)
```bash
./nbgo-gui
```
Interactive dashboard with:
- Multiple views (dashboard, data explorer, debug console, monitoring)
- Real-time system status
- Configuration editor
- Log viewer
- Keyboard navigation

**CLI Mode** - Interactive shell
```bash
./nbgo-gui -mode cli
```
Commands:
- `start <service>` - Start a service
- `stop <service>` - Stop a service
- `status` - Show system status
- `config [key] [value]` - Manage configuration
- `logs [service]` - View service logs

**Settings Mode** - Configuration management
```bash
./nbgo-gui -mode settings
```
Features:
- Load schemas from `schema/` directory
- Aggregate settings from all modules
- Validate configuration constraints
- Edit and save settings
- View comprehensive documentation

### Building GUI CLI

```bash
go build -o nbgo-gui ./gui/cmd
./nbgo-gui -mode tui
```

### GUI Features

- **Auto-Discovery**: Automatically discovers available commands and modules
- **Command Suggestions**: Intelligent command completion and suggestions
- **Keyboard Navigation**: Full keyboard support with arrow keys and vim-like commands
- **Theme Support**: Solarized dark theme with customizable colors
- **Multi-Window Sessions**: Manage multiple concurrent sessions
- **Real-Time Updates**: Live system monitoring and status updates
- **Settings Management**: Load, validate, edit, and save configuration

For detailed GUI documentation, see [gui/cmd/README.md](gui/cmd/README.md)

## Environment Variables

See `.env.example` for all available configuration options.

## Docker Deployment

The project includes comprehensive Docker support:

```bash
# Build image
docker build -t nbgo:latest .

# Run with Docker Compose
docker-compose up -d

# View logs
docker-compose logs -f nbgo
```

## Python Environment Setup

### Using venv
```bash
bash setup_venv.sh
source venv/bin/activate
```

### Using UV
```bash
bash setup_uv_env.sh
source venv_uv/bin/activate
```

## Testing

### Unit Tests
```bash
go test -v -cover ./...
```

### Integration Tests
```bash
robot test/
```

### Coverage Report
```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## Monitoring

Access monitoring dashboards:
- **Grafana**: http://localhost:3000 (admin/admin)
- **Prometheus**: http://localhost:9090

## Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Support

For issues, questions, or suggestions, please open an issue on GitHub.

## Roadmap

- [ ] Enhanced monitoring dashboards
- [ ] Additional exchange integrations
- [ ] WebSocket streaming improvements
- [ ] Kubernetes deployment templates
- [ ] Performance optimizations
- [ ] Extended test coverage
