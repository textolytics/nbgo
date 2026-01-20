package mon

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// MonitorType represents the type of monitoring system
type MonitorType string

const (
	MonitorTypeVictoriaMetrics MonitorType = "victoriametrics"
	MonitorTypeInfluxDB        MonitorType = "influxdb"
	MonitorTypeGrafana         MonitorType = "grafana"
)

// Metric represents a single metric data point
type Metric struct {
	Name      string
	Timestamp time.Time
	Value     float64
	Labels    map[string]string
}

// Alert represents an alert rule
type Alert struct {
	ID        string
	Name      string
	Query     string
	Condition string
	Duration  time.Duration
	Enabled   bool
}

// Monitor defines the interface for monitoring systems
type Monitor interface {
	// Connect establishes connection to the monitoring system
	Connect(ctx context.Context) error

	// Disconnect closes the connection
	Disconnect(ctx context.Context) error

	// WriteMetric writes a single metric
	WriteMetric(ctx context.Context, metric *Metric) error

	// WriteMetrics writes multiple metrics
	WriteMetrics(ctx context.Context, metrics []*Metric) error

	// QueryMetrics queries metrics based on filter
	QueryMetrics(ctx context.Context, query string, start, end time.Time) ([]*Metric, error)

	// CreateAlert creates a new alert rule
	CreateAlert(ctx context.Context, alert *Alert) error

	// DeleteAlert deletes an alert rule
	DeleteAlert(ctx context.Context, alertID string) error

	// GetType returns the monitor type
	GetType() MonitorType

	// IsConnected checks if connected
	IsConnected() bool
}

// BaseMonitor provides common functionality
type BaseMonitor struct {
	typ       MonitorType
	connected bool
	mu        sync.RWMutex
	config    map[string]interface{}
	alerts    map[string]*Alert
}

// NewBaseMonitor creates a new base monitor
func NewBaseMonitor(typ MonitorType) *BaseMonitor {
	return &BaseMonitor{
		typ:       typ,
		connected: false,
		config:    make(map[string]interface{}),
		alerts:    make(map[string]*Alert),
	}
}

// GetType returns the monitor type
func (bm *BaseMonitor) GetType() MonitorType {
	return bm.typ
}

// IsConnected checks if connected
func (bm *BaseMonitor) IsConnected() bool {
	bm.mu.RLock()
	defer bm.mu.RUnlock()
	return bm.connected
}

// SetConnected sets the connected state
func (bm *BaseMonitor) SetConnected(connected bool) {
	bm.mu.Lock()
	defer bm.mu.Unlock()
	bm.connected = connected
}

// VictoriaMetricsMonitor implements VictoriaMetrics monitoring
type VictoriaMetricsMonitor struct {
	*BaseMonitor
	url string
}

// NewVictoriaMetricsMonitor creates a new VictoriaMetrics monitor
func NewVictoriaMetricsMonitor(url string) *VictoriaMetricsMonitor {
	return &VictoriaMetricsMonitor{
		BaseMonitor: NewBaseMonitor(MonitorTypeVictoriaMetrics),
		url:         url,
	}
}

func (vm *VictoriaMetricsMonitor) Connect(ctx context.Context) error {
	vm.SetConnected(true)
	return nil
}

func (vm *VictoriaMetricsMonitor) Disconnect(ctx context.Context) error {
	vm.SetConnected(false)
	return nil
}

func (vm *VictoriaMetricsMonitor) WriteMetric(ctx context.Context, metric *Metric) error {
	if !vm.IsConnected() {
		return fmt.Errorf("VictoriaMetrics not connected")
	}
	return nil
}

func (vm *VictoriaMetricsMonitor) WriteMetrics(ctx context.Context, metrics []*Metric) error {
	if !vm.IsConnected() {
		return fmt.Errorf("VictoriaMetrics not connected")
	}
	return nil
}

func (vm *VictoriaMetricsMonitor) QueryMetrics(ctx context.Context, query string, start, end time.Time) ([]*Metric, error) {
	if !vm.IsConnected() {
		return nil, fmt.Errorf("VictoriaMetrics not connected")
	}
	return []*Metric{}, nil
}

func (vm *VictoriaMetricsMonitor) CreateAlert(ctx context.Context, alert *Alert) error {
	if !vm.IsConnected() {
		return fmt.Errorf("VictoriaMetrics not connected")
	}
	vm.mu.Lock()
	defer vm.mu.Unlock()
	vm.alerts[alert.ID] = alert
	return nil
}

func (vm *VictoriaMetricsMonitor) DeleteAlert(ctx context.Context, alertID string) error {
	if !vm.IsConnected() {
		return fmt.Errorf("VictoriaMetrics not connected")
	}
	vm.mu.Lock()
	defer vm.mu.Unlock()
	delete(vm.alerts, alertID)
	return nil
}

// InfluxDBMonitor implements InfluxDB monitoring
type InfluxDBMonitor struct {
	*BaseMonitor
	url    string
	org    string
	bucket string
	token  string
}

// NewInfluxDBMonitor creates a new InfluxDB monitor
func NewInfluxDBMonitor(url, org, bucket, token string) *InfluxDBMonitor {
	return &InfluxDBMonitor{
		BaseMonitor: NewBaseMonitor(MonitorTypeInfluxDB),
		url:         url,
		org:         org,
		bucket:      bucket,
		token:       token,
	}
}

func (im *InfluxDBMonitor) Connect(ctx context.Context) error {
	im.SetConnected(true)
	return nil
}

func (im *InfluxDBMonitor) Disconnect(ctx context.Context) error {
	im.SetConnected(false)
	return nil
}

func (im *InfluxDBMonitor) WriteMetric(ctx context.Context, metric *Metric) error {
	if !im.IsConnected() {
		return fmt.Errorf("InfluxDB not connected")
	}
	return nil
}

func (im *InfluxDBMonitor) WriteMetrics(ctx context.Context, metrics []*Metric) error {
	if !im.IsConnected() {
		return fmt.Errorf("InfluxDB not connected")
	}
	return nil
}

func (im *InfluxDBMonitor) QueryMetrics(ctx context.Context, query string, start, end time.Time) ([]*Metric, error) {
	if !im.IsConnected() {
		return nil, fmt.Errorf("InfluxDB not connected")
	}
	return []*Metric{}, nil
}

func (im *InfluxDBMonitor) CreateAlert(ctx context.Context, alert *Alert) error {
	if !im.IsConnected() {
		return fmt.Errorf("InfluxDB not connected")
	}
	im.mu.Lock()
	defer im.mu.Unlock()
	im.alerts[alert.ID] = alert
	return nil
}

func (im *InfluxDBMonitor) DeleteAlert(ctx context.Context, alertID string) error {
	if !im.IsConnected() {
		return fmt.Errorf("InfluxDB not connected")
	}
	im.mu.Lock()
	defer im.mu.Unlock()
	delete(im.alerts, alertID)
	return nil
}

// GrafanaMonitor implements Grafana monitoring
type GrafanaMonitor struct {
	*BaseMonitor
	url    string
	apiKey string
}

// NewGrafanaMonitor creates a new Grafana monitor
func NewGrafanaMonitor(url, apiKey string) *GrafanaMonitor {
	return &GrafanaMonitor{
		BaseMonitor: NewBaseMonitor(MonitorTypeGrafana),
		url:         url,
		apiKey:      apiKey,
	}
}

func (gm *GrafanaMonitor) Connect(ctx context.Context) error {
	gm.SetConnected(true)
	return nil
}

func (gm *GrafanaMonitor) Disconnect(ctx context.Context) error {
	gm.SetConnected(false)
	return nil
}

func (gm *GrafanaMonitor) WriteMetric(ctx context.Context, metric *Metric) error {
	if !gm.IsConnected() {
		return fmt.Errorf("Grafana not connected")
	}
	return nil
}

func (gm *GrafanaMonitor) WriteMetrics(ctx context.Context, metrics []*Metric) error {
	if !gm.IsConnected() {
		return fmt.Errorf("Grafana not connected")
	}
	return nil
}

func (gm *GrafanaMonitor) QueryMetrics(ctx context.Context, query string, start, end time.Time) ([]*Metric, error) {
	if !gm.IsConnected() {
		return nil, fmt.Errorf("Grafana not connected")
	}
	return []*Metric{}, nil
}

func (gm *GrafanaMonitor) CreateAlert(ctx context.Context, alert *Alert) error {
	if !gm.IsConnected() {
		return fmt.Errorf("Grafana not connected")
	}
	gm.mu.Lock()
	defer gm.mu.Unlock()
	gm.alerts[alert.ID] = alert
	return nil
}

func (gm *GrafanaMonitor) DeleteAlert(ctx context.Context, alertID string) error {
	if !gm.IsConnected() {
		return fmt.Errorf("Grafana not connected")
	}
	gm.mu.Lock()
	defer gm.mu.Unlock()
	delete(gm.alerts, alertID)
	return nil
}

// Registry manages monitors
type Registry struct {
	mu       sync.RWMutex
	monitors map[string]Monitor
}

// NewRegistry creates a new monitor registry
func NewRegistry() *Registry {
	return &Registry{
		monitors: make(map[string]Monitor),
	}
}

// Register registers a monitor
func (r *Registry) Register(name string, monitor Monitor) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.monitors[name] = monitor
	return nil
}

// Get retrieves a monitor
func (r *Registry) Get(name string) (Monitor, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	monitor, exists := r.monitors[name]
	return monitor, exists
}

// List returns all registered monitors
func (r *Registry) List() []Monitor {
	r.mu.RLock()
	defer r.mu.RUnlock()
	monitors := make([]Monitor, 0, len(r.monitors))
	for _, m := range r.monitors {
		monitors = append(monitors, m)
	}
	return monitors
}
