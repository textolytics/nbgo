package gui

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"github.com/textolytics/nbgo/logs"
	"gopkg.in/yaml.v2"
)

// Setting represents a single configuration setting
type Setting struct {
	Name        string      `json:"name" yaml:"name"`
	Type        string      `json:"type" yaml:"type"`
	Description string      `json:"description" yaml:"description"`
	Default     interface{} `json:"default" yaml:"default"`
	Value       interface{} `json:"value" yaml:"value"`
	Required    bool        `json:"required" yaml:"required"`
	Validate    string      `json:"validate" yaml:"validate"`
	Options     []string    `json:"options" yaml:"options"`
}

// SettingsSchema represents a schema for settings
type SettingsSchema struct {
	Name        string    `json:"name" yaml:"name"`
	Description string    `json:"description" yaml:"description"`
	Version     string    `json:"version" yaml:"version"`
	Settings    []Setting `json:"settings" yaml:"settings"`
}

// SettingsManager manages application settings
type SettingsManager struct {
	mu                 sync.RWMutex
	logger             logs.Logger
	schemas            map[string]*SettingsSchema
	aggregatedSettings map[string]interface{}
	validationErrors   []string
	editableSections   map[string][]Setting
	saveableSections   map[string][]Setting
}

// NewSettingsManager creates a new settings manager
func NewSettingsManager(logger logs.Logger) *SettingsManager {
	return &SettingsManager{
		logger:             logger,
		schemas:            make(map[string]*SettingsSchema),
		aggregatedSettings: make(map[string]interface{}),
		validationErrors:   make([]string, 0),
		editableSections:   make(map[string][]Setting),
		saveableSections:   make(map[string][]Setting),
	}
}

// LoadSchemasFromDirectory loads all schema files from a directory
func (sm *SettingsManager) LoadSchemasFromDirectory(dirPath string) error {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	sm.logger.Infof("Loading schemas from: %s", dirPath)

	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		if os.IsNotExist(err) {
			sm.logger.Warnf("Schema directory not found: %s", dirPath)
			return nil
		}
		return fmt.Errorf("failed to read schema directory: %w", err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if !isSchemaFile(file.Name()) {
			continue
		}

		filePath := filepath.Join(dirPath, file.Name())
		schema, err := sm.loadSchemaFile(filePath)
		if err != nil {
			sm.logger.Warnf("Failed to load schema %s: %v", file.Name(), err)
			continue
		}

		sm.schemas[schema.Name] = schema
		sm.logger.Infof("Schema loaded: %s (v%s)", schema.Name, schema.Version)
	}

	return nil
}

// loadSchemaFile loads a single schema file
func (sm *SettingsManager) loadSchemaFile(filePath string) (*SettingsSchema, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var schema SettingsSchema
	ext := filepath.Ext(filePath)

	if ext == ".json" {
		if err := json.Unmarshal(data, &schema); err != nil {
			return nil, fmt.Errorf("failed to parse JSON: %w", err)
		}
	} else if ext == ".yaml" || ext == ".yml" {
		if err := yaml.Unmarshal(data, &schema); err != nil {
			return nil, fmt.Errorf("failed to parse YAML: %w", err)
		}
	} else {
		return nil, fmt.Errorf("unsupported file format: %s", ext)
	}

	return &schema, nil
}

// GenerateAggregatedSettings combines all loaded schemas into aggregated settings
func (sm *SettingsManager) GenerateAggregatedSettings() error {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	sm.logger.Info("Generating aggregated settings tree...")

	for schemaName, schema := range sm.schemas {
		sm.aggregatedSettings[schemaName] = map[string]interface{}{
			"description": schema.Description,
			"version":     schema.Version,
			"settings":    make([]map[string]interface{}, 0),
		}

		settingsList := make([]map[string]interface{}, 0)
		for _, setting := range schema.Settings {
			settingMap := map[string]interface{}{
				"name":        setting.Name,
				"type":        setting.Type,
				"description": setting.Description,
				"default":     setting.Default,
				"value":       setting.Value,
				"required":    setting.Required,
			}
			if len(setting.Options) > 0 {
				settingMap["options"] = setting.Options
			}
			settingsList = append(settingsList, settingMap)
		}

		// Update the settings in aggregated
		if sectionData, ok := sm.aggregatedSettings[schemaName].(map[string]interface{}); ok {
			sectionData["settings"] = settingsList
		}

		// Track editable and saveable sections
		sm.editableSections[schemaName] = schema.Settings
		sm.saveableSections[schemaName] = schema.Settings
	}

	sm.logger.Infof("Aggregated settings generated: %d schemas", len(sm.schemas))
	return nil
}

// ValidateSettings validates all settings against their schemas
func (sm *SettingsManager) ValidateSettings() error {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	sm.validationErrors = make([]string, 0)

	for schemaName, schema := range sm.schemas {
		for _, setting := range schema.Settings {
			if setting.Required && (setting.Value == nil || setting.Value == "") {
				sm.validationErrors = append(sm.validationErrors,
					fmt.Sprintf("%s.%s: required field is empty", schemaName, setting.Name))
			}
		}
	}

	if len(sm.validationErrors) > 0 {
		sm.logger.Warnf("Settings validation found %d errors", len(sm.validationErrors))
		return fmt.Errorf("validation failed with %d errors", len(sm.validationErrors))
	}

	sm.logger.Info("Settings validation successful")
	return nil
}

// GetSetting retrieves a setting value
func (sm *SettingsManager) GetSetting(section, name string) (interface{}, error) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	schema, exists := sm.schemas[section]
	if !exists {
		return nil, fmt.Errorf("section %s not found", section)
	}

	for _, setting := range schema.Settings {
		if setting.Name == name {
			if setting.Value != nil {
				return setting.Value, nil
			}
			return setting.Default, nil
		}
	}

	return nil, fmt.Errorf("setting %s not found in section %s", name, section)
}

// SetSetting updates a setting value
func (sm *SettingsManager) SetSetting(section, name string, value interface{}) error {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	schema, exists := sm.schemas[section]
	if !exists {
		return fmt.Errorf("section %s not found", section)
	}

	for i, setting := range schema.Settings {
		if setting.Name == name {
			schema.Settings[i].Value = value
			sm.logger.Infof("Setting updated: %s.%s = %v", section, name, value)
			return nil
		}
	}

	return fmt.Errorf("setting %s not found in section %s", name, section)
}

// SaveSettings saves settings to files
func (sm *SettingsManager) SaveSettings(outputDir string) error {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	sm.logger.Infof("Saving settings to: %s", outputDir)

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	for schemaName, schema := range sm.schemas {
		filePath := filepath.Join(outputDir, schemaName+".yaml")

		data, err := yaml.Marshal(schema)
		if err != nil {
			sm.logger.Warnf("Failed to marshal schema %s: %v", schemaName, err)
			continue
		}

		if err := ioutil.WriteFile(filePath, data, 0644); err != nil {
			sm.logger.Warnf("Failed to save schema %s: %v", schemaName, err)
			continue
		}

		sm.logger.Infof("Settings saved: %s", filePath)
	}

	return nil
}

// GetAggregatedSettings returns the aggregated settings tree
func (sm *SettingsManager) GetAggregatedSettings() map[string]interface{} {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	// Create a deep copy
	result := make(map[string]interface{})
	for k, v := range sm.aggregatedSettings {
		result[k] = v
	}
	return result
}

// ListSections returns all available sections
func (sm *SettingsManager) ListSections() []string {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	sections := make([]string, 0, len(sm.schemas))
	for name := range sm.schemas {
		sections = append(sections, name)
	}
	return sections
}

// GetEditableSettings returns settings that can be edited
func (sm *SettingsManager) GetEditableSettings(section string) ([]Setting, error) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	settings, exists := sm.editableSections[section]
	if !exists {
		return nil, fmt.Errorf("section %s not found", section)
	}

	return settings, nil
}

// GetValidationErrors returns all validation errors
func (sm *SettingsManager) GetValidationErrors() []string {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	errors := make([]string, len(sm.validationErrors))
	copy(errors, sm.validationErrors)
	return errors
}

// PrintSettings prints all settings
func (sm *SettingsManager) PrintSettings() {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	for schemaName, schema := range sm.schemas {
		fmt.Printf("\n[%s] %s (v%s)\n", schemaName, schema.Description, schema.Version)
		fmt.Println(string(make([]byte, len(schemaName)+4)))

		for _, setting := range schema.Settings {
			required := ""
			if setting.Required {
				required = " (required)"
			}
			fmt.Printf("  %s: %s%s\n", setting.Name, setting.Description, required)
			fmt.Printf("    Type: %s, Default: %v\n", setting.Type, setting.Default)
			if len(setting.Options) > 0 {
				fmt.Printf("    Options: %v\n", setting.Options)
			}
		}
	}
}

// isSchemaFile checks if a file is a schema file
func isSchemaFile(filename string) bool {
	ext := filepath.Ext(filename)
	return ext == ".json" || ext == ".yaml" || ext == ".yml"
}
