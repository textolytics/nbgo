package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"

	"gopkg.in/yaml.v2"
)

// Config represents the application configuration
type Config struct {
	Version     string                 `json:"version" yaml:"version"`
	Name        string                 `json:"name" yaml:"name"`
	Description string                 `json:"description" yaml:"description"`
	Server      *ServerConfig          `json:"server" yaml:"server"`
	Database    *DatabaseConfig        `json:"database" yaml:"database"`
	MessageBus  *MessageBusConfig      `json:"message_bus" yaml:"message_bus"`
	Monitoring  *MonitoringConfig      `json:"monitoring" yaml:"monitoring"`
	Gateways    map[string]interface{} `json:"gateways" yaml:"gateways"`
	Environment map[string]string      `json:"environment" yaml:"environment"`
}

// ServerConfig represents server configuration
type ServerConfig struct {
	Host     string `json:"host" yaml:"host"`
	Port     int    `json:"port" yaml:"port"`
	TLS      bool   `json:"tls" yaml:"tls"`
	CertFile string `json:"cert_file" yaml:"cert_file"`
	KeyFile  string `json:"key_file" yaml:"key_file"`
}

// DatabaseConfig represents database configuration
type DatabaseConfig struct {
	Type     string `json:"type" yaml:"type"`
	Host     string `json:"host" yaml:"host"`
	Port     int    `json:"port" yaml:"port"`
	Name     string `json:"name" yaml:"name"`
	User     string `json:"user" yaml:"user"`
	Password string `json:"password" yaml:"password"`
}

// MessageBusConfig represents message bus configuration
type MessageBusConfig struct {
	Type     string `json:"type" yaml:"type"`
	Endpoint string `json:"endpoint" yaml:"endpoint"`
	Timeout  int    `json:"timeout" yaml:"timeout"`
}

// MonitoringConfig represents monitoring configuration
type MonitoringConfig struct {
	Enabled bool                   `json:"enabled" yaml:"enabled"`
	Type    string                 `json:"type" yaml:"type"`
	URL     string                 `json:"url" yaml:"url"`
	Metrics map[string]interface{} `json:"metrics" yaml:"metrics"`
}

// Manager manages configuration
type Manager struct {
	mu     sync.RWMutex
	config *Config
}

// NewManager creates a new configuration manager
func NewManager() *Manager {
	return &Manager{
		config: &Config{
			Environment: make(map[string]string),
			Gateways:    make(map[string]interface{}),
		},
	}
}

// LoadJSON loads configuration from JSON file
func (m *Manager) LoadJSON(filepath string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return err
	}

	m.config = &config
	return nil
}

// LoadYAML loads configuration from YAML file
func (m *Manager) LoadYAML(filepath string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return err
	}

	m.config = &config
	return nil
}

// SaveJSON saves configuration to JSON file
func (m *Manager) SaveJSON(filepath string) error {
	m.mu.RLock()
	defer m.mu.RUnlock()

	data, err := json.MarshalIndent(m.config, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filepath, data, 0644)
}

// SaveYAML saves configuration to YAML file
func (m *Manager) SaveYAML(filepath string) error {
	m.mu.RLock()
	defer m.mu.RUnlock()

	data, err := yaml.Marshal(m.config)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filepath, data, 0644)
}

// Get returns the current configuration
func (m *Manager) Get() *Config {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.config
}

// Set updates the configuration
func (m *Manager) Set(config *Config) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.config = config
}

// Validate validates the configuration
func (m *Manager) Validate() error {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if m.config.Version == "" {
		return fmt.Errorf("version not specified")
	}
	if m.config.Name == "" {
		return fmt.Errorf("name not specified")
	}
	if m.config.Server == nil {
		return fmt.Errorf("server config not specified")
	}
	if m.config.Server.Port <= 0 || m.config.Server.Port > 65535 {
		return fmt.Errorf("invalid server port: %d", m.config.Server.Port)
	}

	return nil
}
