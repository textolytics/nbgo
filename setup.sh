#!/bin/bash

# NBGO Project Setup Script
# This script sets up the complete NBGO development environment

set -e

echo "=========================================="
echo "NBGO Market Data System - Setup Script"
echo "=========================================="

# Check prerequisites
echo "Checking prerequisites..."

if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Please install Go 1.22.3 or later."
    exit 1
fi

if ! command -v docker &> /dev/null; then
    echo "❌ Docker is not installed. Please install Docker."
    exit 1
fi

if ! command -v docker-compose &> /dev/null; then
    echo "❌ Docker Compose is not installed. Please install Docker Compose."
    exit 1
fi

echo "✓ All prerequisites met"

# Create directories
echo ""
echo "Creating directory structure..."
mkdir -p build/output build/test-results conf/grafana/provisioning logs data/parquet

echo "✓ Directory structure created"

# Go setup
echo ""
echo "Setting up Go modules..."
go mod download
echo "✓ Go modules downloaded"

# Environment setup
echo ""
echo "Setting up environment..."
if [ ! -f ".env" ]; then
    cp .env.example .env
    echo "✓ Environment file created (.env)"
else
    echo "✓ Environment file already exists"
fi

# Python environment setup
echo ""
echo "Setting up Python environment..."
if command -v python3 &> /dev/null; then
    if [ ! -d "venv" ]; then
        bash setup_venv.sh
        echo "✓ Python virtual environment created"
    else
        echo "✓ Python virtual environment already exists"
    fi
fi

# Docker images
echo ""
echo "Pulling Docker images..."
docker-compose pull
echo "✓ Docker images pulled"

# Create configuration files
echo ""
echo "Creating configuration files..."

cat > conf/prometheus.yml <<'EOF'
global:
  scrape_interval: 15s
  evaluation_interval: 15s

scrape_configs:
  - job_name: 'nbgo'
    static_configs:
      - targets: ['localhost:8080']

alerting:
  alertmanagers:
    - static_configs:
        - targets: []
EOF

echo "✓ Configuration files created"

# Display next steps
echo ""
echo "=========================================="
echo "Setup Complete!"
echo "=========================================="
echo ""
echo "Next steps:"
echo "1. Edit .env file with your settings"
echo "2. Start services: docker-compose up -d"
echo "3. View logs: docker-compose logs -f"
echo "4. Build project: bash build.sh"
echo "5. Run tests: bash test.sh"
echo ""
echo "Documentation:"
echo "- README.md - Project overview"
echo "- DEVELOPMENT.md - Development guide"
echo "- nbgo.yml - Configuration reference"
echo ""
echo "Services available after docker-compose up:"
echo "- App: http://localhost:8080"
echo "- Grafana: http://localhost:3000"
echo "- Prometheus: http://localhost:9090"
echo "- ClickHouse: http://localhost:8123"
echo "- InfluxDB: http://localhost:8086"
echo ""
