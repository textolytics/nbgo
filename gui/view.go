package gui

import (
	"sync"
	"time"
)

// View represents a UI view/panel
type View struct {
	mu          sync.RWMutex
	ID          string
	Title       string
	Description string
	Content     string
	ViewType    string
	Visible     bool
	Width       int
	Height      int
	X           int
	Y           int
	LastUpdate  time.Time
	Data        interface{}
	Commands    []string
	Status      string
}

// NewView creates a new view
func NewView(id, title, description string) *View {
	return &View{
		ID:          id,
		Title:       title,
		Description: description,
		Visible:     true,
		Content:     "",
		Status:      "ready",
		Commands:    make([]string, 0),
		LastUpdate:  time.Now(),
	}
}

// UpdateContent updates the view content
func (v *View) UpdateContent(content string) {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.Content = content
	v.LastUpdate = time.Now()
}

// GetContent returns the view content
func (v *View) GetContent() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.Content
}

// SetStatus sets the view status
func (v *View) SetStatus(status string) {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.Status = status
}

// GetStatus returns the view status
func (v *View) GetStatus() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.Status
}

// AddCommand adds a command to the view
func (v *View) AddCommand(cmd string) {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.Commands = append(v.Commands, cmd)
}

// SetDimensions sets the view dimensions
func (v *View) SetDimensions(x, y, width, height int) {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.X = x
	v.Y = y
	v.Width = width
	v.Height = height
}

// GetDimensions returns the view dimensions
func (v *View) GetDimensions() (x, y, width, height int) {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.X, v.Y, v.Width, v.Height
}

// SetVisibility sets the view visibility
func (v *View) SetVisibility(visible bool) {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.Visible = visible
}

// IsVisible returns if the view is visible
func (v *View) IsVisible() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.Visible
}

// SetData sets the view data
func (v *View) SetData(data interface{}) {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.Data = data
}

// GetData returns the view data
func (v *View) GetData() interface{} {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.Data
}

// Refresh refreshes the view
func (v *View) Refresh() {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.LastUpdate = time.Now()
}

// DashboardView represents the main dashboard view
type DashboardView struct {
	*View
	Widgets []Widget
}

// DataView represents the data explorer view
type DataView struct {
	*View
	Categories map[string][]string
	CurrentCategory string
}

// DebugView represents the debug console view
type DebugView struct {
	*View
	DebugLogs []string
	MaxLogs   int
}

// Widget represents a dashboard widget
type Widget struct {
	ID      string
	Title   string
	Content string
	Type    string
}

// CLIView represents the CLI console view
type CLIView struct {
	*View
	Input       string
	Output      []string
	Suggestions []string
	History     []string
	HistoryIdx  int
}

// TerminalView represents the terminal view
type TerminalView struct {
	*View
	ProcessID   int
	ProcessName string
	Output      string
}

// APIView represents the API explorer view
type APIView struct {
	*View
	Endpoints []APIEndpoint
	Selected  string
}

// APIEndpoint represents an API endpoint
type APIEndpoint struct {
	Path        string
	Method      string
	Description string
	Params      map[string]string
	Response    string
}

// MonitoringView represents the monitoring view
type MonitoringView struct {
	*View
	Metrics     map[string]interface{}
	Services    map[string]string
	Alerts      []Alert
}

// Alert represents a monitoring alert
type Alert struct {
	ID        string
	Severity  string
	Message   string
	Timestamp int64
}

// LogsView represents the logs view
type LogsView struct {
	*View
	LogEntries []LogEntry
	Filter     string
	MaxEntries int
}

// LogEntry represents a log entry
type LogEntry struct {
	Timestamp int64
	Level     string
	Message   string
	Component string
}

// EnvironmentView represents the environment variables view
type EnvironmentView struct {
	*View
	Variables map[string]string
}

// ConfigurationView represents the configuration view
type ConfigurationView struct {
	*View
	Sections map[string]map[string]interface{}
}
