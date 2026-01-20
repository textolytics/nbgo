package core
package core

import (
	"context"
	"sync"
)

// ProviderType represents the type of provider
type ProviderType string

const (
	ProviderTypeC             ProviderType = "c"
	ProviderTypeRust          ProviderType = "rust"
	ProviderTypeGo            ProviderType = "go"
	ProviderTypePython        ProviderType = "python"
	ProviderTypeFlutter       ProviderType = "flutter"
	ProviderTypeRobotFramework ProviderType = "robot_framework"
	ProviderTypeMCP           ProviderType = "mcp"
)

// Provider defines the interface for all providers
type Provider interface {
	// Initialize initializes the provider
	Initialize(ctx context.Context) error
	
	// Start starts the provider
	Start(ctx context.Context) error
	
	// Stop stops the provider gracefully
	Stop(ctx context.Context) error
	
	// GetName returns the provider name
	GetName() string
	
	// GetType returns the provider type
	GetType() ProviderType
	
	// IsHealthy checks if the provider is healthy
	IsHealthy(ctx context.Context) error
}

// BaseProvider provides common functionality for all providers
type BaseProvider struct {
	name   string
	typ    ProviderType
	mu     sync.RWMutex
	active bool
}

// NewBaseProvider creates a new base provider
func NewBaseProvider(name string, typ ProviderType) *BaseProvider {
	return &BaseProvider{
		name:   name,
		typ:    typ,
		active: false,
	}
}

// GetName returns the provider name
func (bp *BaseProvider) GetName() string {
	bp.mu.RLock()
	defer bp.mu.RUnlock()
	return bp.name
}

// GetType returns the provider type
func (bp *BaseProvider) GetType() ProviderType {
	bp.mu.RLock()
	defer bp.mu.RUnlock()
	return bp.typ
}

// SetActive marks the provider as active or inactive
func (bp *BaseProvider) SetActive(active bool) {
	bp.mu.Lock()
	defer bp.mu.Unlock()
	bp.active = active
}

// IsActive checks if the provider is active
func (bp *BaseProvider) IsActive() bool {
	bp.mu.RLock()
	defer bp.mu.RUnlock()
	return bp.active
}

// Registry manages all providers
type Registry struct {
	mu        sync.RWMutex
	providers map[string]Provider
}

// NewRegistry creates a new provider registry
func NewRegistry() *Registry {
	return &Registry{
		providers: make(map[string]Provider),
	}
}

// Register registers a provider
func (r *Registry) Register(name string, provider Provider) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.providers[name] = provider
	return nil
}

// Get retrieves a provider by name
func (r *Registry) Get(name string) (Provider, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	provider, exists := r.providers[name]
	return provider, exists
}

// List returns all registered providers
func (r *Registry) List() []Provider {
	r.mu.RLock()
	defer r.mu.RUnlock()
	providers := make([]Provider, 0, len(r.providers))
	for _, p := range r.providers {
		providers = append(providers, p)
	}
	return providers
}

// StartAll starts all registered providers
func (r *Registry) StartAll(ctx context.Context) error {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, provider := range r.providers {
		if err := provider.Start(ctx); err != nil {
			return err
		}
	}
	return nil
}

// StopAll stops all registered providers
func (r *Registry) StopAll(ctx context.Context) error {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, provider := range r.providers {
		if err := provider.Stop(ctx); err != nil {
			return err
		}
	}
	return nil
}
