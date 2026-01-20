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
├── logs/                 # Logging system
├── mb/                   # Message bus implementations (ZMQ, MQTT)
├── mcp/                  # Model Context Protocol server
├── mon/                  # Monitoring implementations (VictoriaMetrics, InfluxDB, Grafana)
├── run/                  # Runtime management
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
