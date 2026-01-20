package dw

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// DataWarehouseType represents the type of data warehouse
type DataWarehouseType string

const (
	DataWarehouseTypeClickHouse DataWarehouseType = "clickhouse"
	DataWarehouseTypeInfluxDB   DataWarehouseType = "influxdb"
	DataWarehouseTypeParquet    DataWarehouseType = "parquet"
)

// Row represents a single row of data
type Row map[string]interface{}

// QueryResult represents query results
type QueryResult struct {
	Columns []string
	Rows    []Row
	Total   int64
}

// DataWarehouse defines the interface for data warehouse implementations
type DataWarehouse interface {
	// Connect establishes connection to the warehouse
	Connect(ctx context.Context) error

	// Disconnect closes the connection
	Disconnect(ctx context.Context) error

	// CreateTable creates a new table
	CreateTable(ctx context.Context, tableName string, schema map[string]string) error

	// InsertBatch inserts multiple rows
	InsertBatch(ctx context.Context, tableName string, rows []Row) error

	// Query executes a query and returns results
	Query(ctx context.Context, sql string) (*QueryResult, error)

	// DeleteOldData deletes data older than retention period
	DeleteOldData(ctx context.Context, tableName string, retentionDays int) error

	// GetType returns the warehouse type
	GetType() DataWarehouseType

	// IsConnected checks if connected
	IsConnected() bool
}

// BaseDataWarehouse provides common functionality
type BaseDataWarehouse struct {
	typ       DataWarehouseType
	connected bool
	mu        sync.RWMutex
	config    map[string]interface{}
}

// NewBaseDataWarehouse creates a new base data warehouse
func NewBaseDataWarehouse(typ DataWarehouseType) *BaseDataWarehouse {
	return &BaseDataWarehouse{
		typ:       typ,
		connected: false,
		config:    make(map[string]interface{}),
	}
}

// GetType returns the warehouse type
func (bdw *BaseDataWarehouse) GetType() DataWarehouseType {
	return bdw.typ
}

// IsConnected checks if connected
func (bdw *BaseDataWarehouse) IsConnected() bool {
	bdw.mu.RLock()
	defer bdw.mu.RUnlock()
	return bdw.connected
}

// SetConnected sets the connected state
func (bdw *BaseDataWarehouse) SetConnected(connected bool) {
	bdw.mu.Lock()
	defer bdw.mu.Unlock()
	bdw.connected = connected
}

// ClickHouseWarehouse implements ClickHouse data warehouse
type ClickHouseWarehouse struct {
	*BaseDataWarehouse
	host     string
	port     int
	database string
	username string
	password string
}

// NewClickHouseWarehouse creates a new ClickHouse warehouse
func NewClickHouseWarehouse(host string, port int, database, username, password string) *ClickHouseWarehouse {
	return &ClickHouseWarehouse{
		BaseDataWarehouse: NewBaseDataWarehouse(DataWarehouseTypeClickHouse),
		host:              host,
		port:              port,
		database:          database,
		username:          username,
		password:          password,
	}
}

func (chw *ClickHouseWarehouse) Connect(ctx context.Context) error {
	chw.SetConnected(true)
	return nil
}

func (chw *ClickHouseWarehouse) Disconnect(ctx context.Context) error {
	chw.SetConnected(false)
	return nil
}

func (chw *ClickHouseWarehouse) CreateTable(ctx context.Context, tableName string, schema map[string]string) error {
	if !chw.IsConnected() {
		return fmt.Errorf("ClickHouse not connected")
	}
	return nil
}

func (chw *ClickHouseWarehouse) InsertBatch(ctx context.Context, tableName string, rows []Row) error {
	if !chw.IsConnected() {
		return fmt.Errorf("ClickHouse not connected")
	}
	return nil
}

func (chw *ClickHouseWarehouse) Query(ctx context.Context, sql string) (*QueryResult, error) {
	if !chw.IsConnected() {
		return nil, fmt.Errorf("ClickHouse not connected")
	}
	return &QueryResult{Columns: []string{}, Rows: []Row{}, Total: 0}, nil
}

func (chw *ClickHouseWarehouse) DeleteOldData(ctx context.Context, tableName string, retentionDays int) error {
	if !chw.IsConnected() {
		return fmt.Errorf("ClickHouse not connected")
	}
	return nil
}

// InfluxDBWarehouse implements InfluxDB data warehouse
type InfluxDBWarehouse struct {
	*BaseDataWarehouse
	url    string
	org    string
	bucket string
	token  string
}

// NewInfluxDBWarehouse creates a new InfluxDB warehouse
func NewInfluxDBWarehouse(url, org, bucket, token string) *InfluxDBWarehouse {
	return &InfluxDBWarehouse{
		BaseDataWarehouse: NewBaseDataWarehouse(DataWarehouseTypeInfluxDB),
		url:               url,
		org:               org,
		bucket:            bucket,
		token:             token,
	}
}

func (iw *InfluxDBWarehouse) Connect(ctx context.Context) error {
	iw.SetConnected(true)
	return nil
}

func (iw *InfluxDBWarehouse) Disconnect(ctx context.Context) error {
	iw.SetConnected(false)
	return nil
}

func (iw *InfluxDBWarehouse) CreateTable(ctx context.Context, tableName string, schema map[string]string) error {
	if !iw.IsConnected() {
		return fmt.Errorf("InfluxDB not connected")
	}
	return nil
}

func (iw *InfluxDBWarehouse) InsertBatch(ctx context.Context, tableName string, rows []Row) error {
	if !iw.IsConnected() {
		return fmt.Errorf("InfluxDB not connected")
	}
	return nil
}

func (iw *InfluxDBWarehouse) Query(ctx context.Context, sql string) (*QueryResult, error) {
	if !iw.IsConnected() {
		return nil, fmt.Errorf("InfluxDB not connected")
	}
	return &QueryResult{Columns: []string{}, Rows: []Row{}, Total: 0}, nil
}

func (iw *InfluxDBWarehouse) DeleteOldData(ctx context.Context, tableName string, retentionDays int) error {
	if !iw.IsConnected() {
		return fmt.Errorf("InfluxDB not connected")
	}
	return nil
}

// ParquetWarehouse implements Parquet file-based data warehouse
type ParquetWarehouse struct {
	*BaseDataWarehouse
	basePath string
}

// NewParquetWarehouse creates a new Parquet warehouse
func NewParquetWarehouse(basePath string) *ParquetWarehouse {
	return &ParquetWarehouse{
		BaseDataWarehouse: NewBaseDataWarehouse(DataWarehouseTypeParquet),
		basePath:          basePath,
	}
}

func (pw *ParquetWarehouse) Connect(ctx context.Context) error {
	pw.SetConnected(true)
	return nil
}

func (pw *ParquetWarehouse) Disconnect(ctx context.Context) error {
	pw.SetConnected(false)
	return nil
}

func (pw *ParquetWarehouse) CreateTable(ctx context.Context, tableName string, schema map[string]string) error {
	if !pw.IsConnected() {
		return fmt.Errorf("Parquet not connected")
	}
	return nil
}

func (pw *ParquetWarehouse) InsertBatch(ctx context.Context, tableName string, rows []Row) error {
	if !pw.IsConnected() {
		return fmt.Errorf("Parquet not connected")
	}
	return nil
}

func (pw *ParquetWarehouse) Query(ctx context.Context, sql string) (*QueryResult, error) {
	if !pw.IsConnected() {
		return nil, fmt.Errorf("Parquet not connected")
	}
	return &QueryResult{Columns: []string{}, Rows: []Row{}, Total: 0}, nil
}

func (pw *ParquetWarehouse) DeleteOldData(ctx context.Context, tableName string, retentionDays int) error {
	if !pw.IsConnected() {
		return fmt.Errorf("Parquet not connected")
	}
	// For Parquet files, rotate old files instead
	return nil
}

// RetentionPolicy defines data retention settings
type RetentionPolicy struct {
	Name         string
	Duration     time.Duration
	ReplicaCount int
}

// Registry manages data warehouses
type Registry struct {
	mu        sync.RWMutex
	warehouse map[string]DataWarehouse
}

// NewRegistry creates a new data warehouse registry
func NewRegistry() *Registry {
	return &Registry{
		warehouse: make(map[string]DataWarehouse),
	}
}

// Register registers a warehouse
func (r *Registry) Register(name string, warehouse DataWarehouse) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.warehouse[name] = warehouse
	return nil
}

// Get retrieves a warehouse
func (r *Registry) Get(name string) (DataWarehouse, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	warehouse, exists := r.warehouse[name]
	return warehouse, exists
}

// List returns all registered warehouses
func (r *Registry) List() []DataWarehouse {
	r.mu.RLock()
	defer r.mu.RUnlock()
	warehouses := make([]DataWarehouse, 0, len(r.warehouse))
	for _, w := range r.warehouse {
		warehouses = append(warehouses, w)
	}
	return warehouses
}
