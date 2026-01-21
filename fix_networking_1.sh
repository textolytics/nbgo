
#!/usr/bin/env bash
set -e

NET=observability-net

echo "=== FIX PROVIDER NETWORKING ==="

# 1. Create shared Docker network
docker network inspect $NET >/dev/null 2>&1 || docker network create $NET

# 2. Firewall: allow container-to-container traffic
sudo iptables -C FORWARD -i docker0 -j ACCEPT 2>/dev/null \
 || sudo iptables -A FORWARD -i docker0 -j ACCEPT

sudo iptables -C FORWARD -o docker0 -j ACCEPT 2>/dev/null \
 || sudo iptables -A FORWARD -o docker0 -j ACCEPT

# Persist iptables if possible
sudo iptables-save | sudo tee /etc/iptables/rules.v4 >/dev/null || true

# 3. Run core providers
docker run -d --name prometheus \
  --network $NET \
  -p 9090:9090 \
  prom/prometheus

docker run -d --name grafana \
  --network $NET \
  -p 3000:3000 \
  grafana/grafana

docker run -d --name clickhouse \
  --network $NET \
  -p 8123:8123 \
  clickhouse/clickhouse-server

docker run -d --name influxdb \
  --network $NET \
  -p 8086:8086 \
  influxdb:1.8

docker run -d --name postgres \
  --network $NET \
  -e POSTGRES_PASSWORD=postgres \
  -p 5432:5432 \
  postgres:15

echo "✅ Networking services started"

