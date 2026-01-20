package run

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// RuntimeConfig represents runtime configuration
type RuntimeConfig struct {
	Name                string
	Version             string
	StartTimeout        time.Duration
	ShutdownTimeout     time.Duration
	MaxConnections      int
	MaxRetries          int
	RetryInterval       time.Duration
	HealthCheckInterval time.Duration
}

// Component represents a component in the runtime
type Component interface {
	// Name returns the component name
	Name() string

	// Start starts the component
	Start(ctx context.Context) error

	// Stop stops the component gracefully
	Stop(ctx context.Context) error

	// IsHealthy checks if the component is healthy
	IsHealthy(ctx context.Context) error
}

// Runtime manages application runtime and lifecycle
type Runtime struct {
	mu         sync.RWMutex
	config     RuntimeConfig
	components map[string]Component
	running    bool
	ctx        context.Context
	cancel     context.CancelFunc
	signals    chan os.Signal
}

// NewRuntime creates a new runtime instance
func NewRuntime(config RuntimeConfig) *Runtime {
	ctx, cancel := context.WithCancel(context.Background())
	return &Runtime{
		config:     config,
		components: make(map[string]Component),
		signals:    make(chan os.Signal, 1),
		ctx:        ctx,
		cancel:     cancel,
	}
}

// Register registers a component
func (r *Runtime) Register(component Component) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	name := component.Name()
	if _, exists := r.components[name]; exists {
		return fmt.Errorf("component %s already registered", name)
	}

	r.components[name] = component
	return nil
}

// Get retrieves a component by name
func (r *Runtime) Get(name string) (Component, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	component, exists := r.components[name]
	return component, exists
}

// List returns all registered components
func (r *Runtime) List() []Component {
	r.mu.RLock()
	defer r.mu.RUnlock()
	components := make([]Component, 0, len(r.components))
	for _, comp := range r.components {
		components = append(components, comp)
	}
	return components
}

// Start starts the runtime and all components
func (r *Runtime) Start(ctx context.Context) error {
	r.mu.Lock()
	if r.running {
		r.mu.Unlock()
		return fmt.Errorf("runtime already running")
	}
	r.running = true
	r.mu.Unlock()

	// Setup signal handling
	signal.Notify(r.signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	// Start all components
	startCtx, cancel := context.WithTimeout(ctx, r.config.StartTimeout)
	defer cancel()

	for _, component := range r.List() {
		if err := component.Start(startCtx); err != nil {
			return fmt.Errorf("failed to start component %s: %v", component.Name(), err)
		}
	}

	// Start health check goroutine
	go r.healthCheck()

	// Start signal handler
	go r.handleSignals()

	return nil
}

// Stop stops the runtime and all components
func (r *Runtime) Stop(ctx context.Context) error {
	r.mu.Lock()
	if !r.running {
		r.mu.Unlock()
		return fmt.Errorf("runtime not running")
	}
	r.running = false
	r.mu.Unlock()

	r.cancel()

	// Stop all components in reverse order
	stopCtx, cancel := context.WithTimeout(ctx, r.config.ShutdownTimeout)
	defer cancel()

	components := r.List()
	for i := len(components) - 1; i >= 0; i-- {
		component := components[i]
		if err := component.Stop(stopCtx); err != nil {
			// Log error but continue shutting down
			fmt.Printf("Error stopping component %s: %v\n", component.Name(), err)
		}
	}

	return nil
}

// healthCheck periodically checks component health
func (r *Runtime) healthCheck() {
	ticker := time.NewTicker(r.config.HealthCheckInterval)
	defer ticker.Stop()

	for {
		select {
		case <-r.ctx.Done():
			return
		case <-ticker.C:
			for _, component := range r.List() {
				if err := component.IsHealthy(r.ctx); err != nil {
					fmt.Printf("Component %s is unhealthy: %v\n", component.Name(), err)
				}
			}
		}
	}
}

// handleSignals handles OS signals
func (r *Runtime) handleSignals() {
	for signal := range r.signals {
		switch signal {
		case syscall.SIGINT, syscall.SIGTERM:
			fmt.Printf("Received signal: %v, shutting down...\n", signal)
			ctx, cancel := context.WithTimeout(context.Background(), r.config.ShutdownTimeout)
			r.Stop(ctx)
			cancel()
			os.Exit(0)
		case syscall.SIGHUP:
			fmt.Println("Received SIGHUP, reloading configuration...")
			// Reload configuration logic here
		}
	}
}

// IsRunning checks if runtime is running
func (r *Runtime) IsRunning() bool {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.running
}
