package schema

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

// SchemaType represents the type of schema
type SchemaType string

const (
	SchemaTypeProvider    SchemaType = "provider"
	SchemaTypeAnnotation  SchemaType = "annotation"
	SchemaTypeBuildConfig SchemaType = "build_config"
	SchemaTypeBuildTarget SchemaType = "build_target"
)

// Field represents a schema field
type Field struct {
	Name        string      `json:"name"`
	Type        string      `json:"type"`
	Description string      `json:"description"`
	Required    bool        `json:"required"`
	Default     interface{} `json:"default,omitempty"`
	Options     []string    `json:"options,omitempty"`
}

// Schema represents a configuration schema
type Schema struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Version     string     `json:"version"`
	Type        SchemaType `json:"type"`
	Description string     `json:"description"`
	Fields      []Field    `json:"fields"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	Valid       bool       `json:"valid"`
	Errors      []string   `json:"errors,omitempty"`
}

// Annotation represents provider configuration annotation
type Annotation struct {
	ID          string                 `json:"id"`
	Provider    string                 `json:"provider"`
	Description string                 `json:"description"`
	Settings    map[string]interface{} `json:"settings"`
	Environment map[string]string      `json:"environment"`
	CreatedAt   time.Time              `json:"created_at"`
	UpdatedAt   time.Time              `json:"updated_at"`
}

// BuildConfiguration represents build configuration for a provider
type BuildConfiguration struct {
	ID          string            `json:"id"`
	Provider    string            `json:"provider"`
	Version     string            `json:"version"`
	Targets     []string          `json:"targets"`
	Flags       []string          `json:"flags"`
	Environment map[string]string `json:"environment"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
}

// BuildTarget represents a build target (OS/Arch combination)
type BuildTarget struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	OS         string    `json:"os"`
	Arch       string    `json:"arch"`
	Enabled    bool      `json:"enabled"`
	BuildFlags []string  `json:"build_flags,omitempty"`
	OutputPath string    `json:"output_path,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// ProviderMetadata contains metadata about a provider
type ProviderMetadata struct {
	ID               string              `json:"id"`
	Name             string              `json:"name"`
	Type             string              `json:"type"`
	Version          string              `json:"version"`
	DocumentationURL string              `json:"documentation_url"`
	Schema           *Schema             `json:"schema,omitempty"`
	Annotation       *Annotation         `json:"annotation,omitempty"`
	BuildConfig      *BuildConfiguration `json:"build_config,omitempty"`
	BuildTargets     []BuildTarget       `json:"build_targets,omitempty"`
	LastScanned      time.Time           `json:"last_scanned"`
}

// Repository manages schemas, annotations, and configurations
type Repository struct {
	mu               sync.RWMutex
	schemas          map[string]*Schema
	annotations      map[string]*Annotation
	buildConfigs     map[string]*BuildConfiguration
	buildTargets     map[string]*BuildTarget
	providerMetadata map[string]*ProviderMetadata
	validators       map[SchemaType]SchemaValidator
}

// SchemaValidator validates schemas of a specific type
type SchemaValidator interface {
	Validate(schema *Schema) (bool, []string)
}

// NewRepository creates a new schema repository
func NewRepository() *Repository {
	return &Repository{
		schemas:          make(map[string]*Schema),
		annotations:      make(map[string]*Annotation),
		buildConfigs:     make(map[string]*BuildConfiguration),
		buildTargets:     make(map[string]*BuildTarget),
		providerMetadata: make(map[string]*ProviderMetadata),
		validators:       make(map[SchemaType]SchemaValidator),
	}
}

// RegisterValidator registers a schema validator
func (r *Repository) RegisterValidator(schemaType SchemaType, validator SchemaValidator) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.validators[schemaType] = validator
}

// SaveSchema saves a schema
func (r *Repository) SaveSchema(schema *Schema) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Validate schema
	if validator, exists := r.validators[schema.Type]; exists {
		valid, errors := validator.Validate(schema)
		schema.Valid = valid
		schema.Errors = errors
		if !valid {
			return fmt.Errorf("schema validation failed: %v", errors)
		}
	}

	schema.UpdatedAt = time.Now()
	r.schemas[schema.ID] = schema
	return nil
}

// GetSchema retrieves a schema by ID
func (r *Repository) GetSchema(id string) (*Schema, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	schema, exists := r.schemas[id]
	return schema, exists
}

// ListSchemas returns all schemas of a type
func (r *Repository) ListSchemas(schemaType SchemaType) []*Schema {
	r.mu.RLock()
	defer r.mu.RUnlock()
	schemas := make([]*Schema, 0)
	for _, schema := range r.schemas {
		if schema.Type == schemaType {
			schemas = append(schemas, schema)
		}
	}
	return schemas
}

// SaveAnnotation saves an annotation
func (r *Repository) SaveAnnotation(annotation *Annotation) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	annotation.UpdatedAt = time.Now()
	r.annotations[annotation.ID] = annotation
	return nil
}

// GetAnnotation retrieves an annotation by ID
func (r *Repository) GetAnnotation(id string) (*Annotation, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	annotation, exists := r.annotations[id]
	return annotation, exists
}

// ListAnnotations returns all annotations for a provider
func (r *Repository) ListAnnotations(provider string) []*Annotation {
	r.mu.RLock()
	defer r.mu.RUnlock()
	annotations := make([]*Annotation, 0)
	for _, annotation := range r.annotations {
		if annotation.Provider == provider {
			annotations = append(annotations, annotation)
		}
	}
	return annotations
}

// SaveBuildConfiguration saves a build configuration
func (r *Repository) SaveBuildConfiguration(config *BuildConfiguration) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	config.UpdatedAt = time.Now()
	r.buildConfigs[config.ID] = config
	return nil
}

// GetBuildConfiguration retrieves a build configuration
func (r *Repository) GetBuildConfiguration(id string) (*BuildConfiguration, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	config, exists := r.buildConfigs[id]
	return config, exists
}

// SaveBuildTarget saves a build target
func (r *Repository) SaveBuildTarget(target *BuildTarget) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	target.UpdatedAt = time.Now()
	r.buildTargets[target.ID] = target
	return nil
}

// GetBuildTarget retrieves a build target
func (r *Repository) GetBuildTarget(id string) (*BuildTarget, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	target, exists := r.buildTargets[id]
	return target, exists
}

// ListBuildTargets returns all enabled build targets
func (r *Repository) ListBuildTargets() []*BuildTarget {
	r.mu.RLock()
	defer r.mu.RUnlock()
	targets := make([]*BuildTarget, 0)
	for _, target := range r.buildTargets {
		if target.Enabled {
			targets = append(targets, target)
		}
	}
	return targets
}

// SaveProviderMetadata saves provider metadata
func (r *Repository) SaveProviderMetadata(metadata *ProviderMetadata) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	metadata.LastScanned = time.Now()
	r.providerMetadata[metadata.ID] = metadata
	return nil
}

// GetProviderMetadata retrieves provider metadata
func (r *Repository) GetProviderMetadata(id string) (*ProviderMetadata, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	metadata, exists := r.providerMetadata[id]
	return metadata, exists
}

// ListProviderMetadata returns all provider metadata
func (r *Repository) ListProviderMetadata() []*ProviderMetadata {
	r.mu.RLock()
	defer r.mu.RUnlock()
	metadata := make([]*ProviderMetadata, 0, len(r.providerMetadata))
	for _, m := range r.providerMetadata {
		metadata = append(metadata, m)
	}
	return metadata
}

// ExportSchema exports a schema to JSON
func (r *Repository) ExportSchema(schemaID string) ([]byte, error) {
	r.mu.RLock()
	schema, exists := r.schemas[schemaID]
	r.mu.RUnlock()

	if !exists {
		return nil, fmt.Errorf("schema not found: %s", schemaID)
	}

	return json.MarshalIndent(schema, "", "  ")
}

// ValidateSchemaData validates data against a schema
func (r *Repository) ValidateSchemaData(schemaID string, data map[string]interface{}) (bool, []string) {
	r.mu.RLock()
	schema, exists := r.schemas[schemaID]
	r.mu.RUnlock()

	if !exists {
		return false, []string{"schema not found"}
	}

	errors := make([]string, 0)

	// Check required fields
	for _, field := range schema.Fields {
		if field.Required {
			if _, exists := data[field.Name]; !exists {
				errors = append(errors, fmt.Sprintf("required field missing: %s", field.Name))
			}
		}
	}

	return len(errors) == 0, errors
}

// DiscoveryResult contains results from a discovery scan
type DiscoveryResult struct {
	ScanTime          time.Time
	ProvidersFound    int
	SchemasGenerated  int
	AnnotationsFound  int
	BuildTargetsFound int
	Errors            []string
}

// Discovery scans provider documentation and generates schemas/annotations
type Discovery struct {
	repository *Repository
	providers  []string
	mu         sync.RWMutex
}

// NewDiscovery creates a new discovery instance
func NewDiscovery(repository *Repository, providers []string) *Discovery {
	return &Discovery{
		repository: repository,
		providers:  providers,
	}
}

// ScanProviders scans all providers for schema/annotation generation
func (d *Discovery) ScanProviders(ctx context.Context) (*DiscoveryResult, error) {
	result := &DiscoveryResult{
		ScanTime: time.Now(),
		Errors:   make([]string, 0),
	}

	for _, provider := range d.providers {
		select {
		case <-ctx.Done():
			return result, ctx.Err()
		default:
			// Generate provider metadata
			metadata := &ProviderMetadata{
				ID:          provider,
				Name:        provider,
				Type:        "provider",
				LastScanned: time.Now(),
			}

			// Create default schema
			schema := &Schema{
				ID:      fmt.Sprintf("schema_%s", provider),
				Name:    provider,
				Version: "1.0",
				Type:    SchemaTypeProvider,
				Fields: []Field{
					{Name: "name", Type: "string", Required: true},
					{Name: "version", Type: "string", Required: true},
					{Name: "enabled", Type: "boolean", Required: false, Default: true},
				},
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				Valid:     true,
			}

			if err := d.repository.SaveSchema(schema); err != nil {
				result.Errors = append(result.Errors, fmt.Sprintf("error saving schema for %s: %v", provider, err))
				continue
			}

			metadata.Schema = schema
			result.SchemasGenerated++

			// Create default annotation
			annotation := &Annotation{
				ID:          fmt.Sprintf("annotation_%s", provider),
				Provider:    provider,
				Settings:    make(map[string]interface{}),
				Environment: make(map[string]string),
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			}

			if err := d.repository.SaveAnnotation(annotation); err != nil {
				result.Errors = append(result.Errors, fmt.Sprintf("error saving annotation for %s: %v", provider, err))
				continue
			}

			metadata.Annotation = annotation
			result.AnnotationsFound++

			// Save metadata
			if err := d.repository.SaveProviderMetadata(metadata); err != nil {
				result.Errors = append(result.Errors, fmt.Sprintf("error saving metadata for %s: %v", provider, err))
				continue
			}

			result.ProvidersFound++
		}
	}

	return result, nil
}

// GenerateDefaultBuildTargets generates default build targets
func (d *Discovery) GenerateDefaultBuildTargets(ctx context.Context) error {
	targets := []BuildTarget{
		{ID: "linux_amd64", Name: "Linux x86_64", OS: "linux", Arch: "amd64", Enabled: true, CreatedAt: time.Now()},
		{ID: "linux_arm64", Name: "Linux ARM64", OS: "linux", Arch: "arm64", Enabled: true, CreatedAt: time.Now()},
		{ID: "darwin_amd64", Name: "macOS x86_64", OS: "darwin", Arch: "amd64", Enabled: true, CreatedAt: time.Now()},
		{ID: "darwin_arm64", Name: "macOS ARM64", OS: "darwin", Arch: "arm64", Enabled: true, CreatedAt: time.Now()},
		{ID: "windows_amd64", Name: "Windows x86_64", OS: "windows", Arch: "amd64", Enabled: true, CreatedAt: time.Now()},
	}

	for _, target := range targets {
		if err := d.repository.SaveBuildTarget(&target); err != nil {
			return err
		}
	}

	return nil
}
