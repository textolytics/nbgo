FROM golang:1.22.3-alpine AS builder

WORKDIR /app

# Install dependencies
RUN apk add --no-cache git ca-certificates tzdata make

# Copy all source code first (including sub-modules)
COPY . .

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download && go mod verify

# Download indirect dependencies and tidy
RUN go mod tidy

# Build application
RUN CGO_ENABLED=0 GOOS=linux go build -v -a -installsuffix cgo -o nbgo .

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/nbgo .

# Copy configuration files
COPY nbgo.yml .
COPY conf/ conf/

# Expose default port
EXPOSE 8080

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD ./nbgo health || exit 1

CMD ["./nbgo"]
