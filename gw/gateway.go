package gw

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// GatewayType represents the type of gateway
type GatewayType string

const (
	GatewayTypeGate   GatewayType = "gate"
	GatewayTypeFreedx GatewayType = "freedx"
)

// OrderBook represents market order book
type OrderBook struct {
	Symbol    string
	Timestamp time.Time
	Bids      []PriceLevel
	Asks      []PriceLevel
}

// PriceLevel represents a single price level in order book
type PriceLevel struct {
	Price    float64
	Quantity float64
}

// Trade represents a completed trade
type Trade struct {
	ID        string
	Symbol    string
	Timestamp time.Time
	Price     float64
	Quantity  float64
	Side      string // "buy" or "sell"
}

// Kline represents candlestick data
type Kline struct {
	Symbol    string
	Interval  string
	Timestamp time.Time
	Open      float64
	High      float64
	Low       float64
	Close     float64
	Volume    float64
}

// Gateway defines the interface for market data gateways
type Gateway interface {
	// Connect establishes connection to the gateway
	Connect(ctx context.Context) error

	// Disconnect closes the connection
	Disconnect(ctx context.Context) error

	// GetOrderBook retrieves the current order book for a symbol
	GetOrderBook(ctx context.Context, symbol string) (*OrderBook, error)

	// GetTrades retrieves recent trades for a symbol
	GetTrades(ctx context.Context, symbol string, limit int) ([]*Trade, error)

	// GetKlines retrieves candlestick data
	GetKlines(ctx context.Context, symbol, interval string, start, end time.Time) ([]*Kline, error)

	// SubscribeOrderBook subscribes to order book updates
	SubscribeOrderBook(ctx context.Context, symbol string, handler OrderBookHandler) error

	// SubscribeTrades subscribes to trade updates
	SubscribeTrades(ctx context.Context, symbol string, handler TradeHandler) error

	// SubscribeKlines subscribes to kline updates
	SubscribeKlines(ctx context.Context, symbol, interval string, handler KlineHandler) error

	// GetType returns the gateway type
	GetType() GatewayType

	// IsConnected checks if connected
	IsConnected() bool
}

// OrderBookHandler is the callback for order book updates
type OrderBookHandler func(ob *OrderBook) error

// TradeHandler is the callback for trade updates
type TradeHandler func(trade *Trade) error

// KlineHandler is the callback for kline updates
type KlineHandler func(kline *Kline) error

// BaseGateway provides common functionality
type BaseGateway struct {
	typ           GatewayType
	connected     bool
	mu            sync.RWMutex
	config        map[string]interface{}
	obsHandlers   map[string][]OrderBookHandler
	tradeHandlers map[string][]TradeHandler
	klineHandlers map[string][]KlineHandler
}

// NewBaseGateway creates a new base gateway
func NewBaseGateway(typ GatewayType) *BaseGateway {
	return &BaseGateway{
		typ:           typ,
		connected:     false,
		config:        make(map[string]interface{}),
		obsHandlers:   make(map[string][]OrderBookHandler),
		tradeHandlers: make(map[string][]TradeHandler),
		klineHandlers: make(map[string][]KlineHandler),
	}
}

// GetType returns the gateway type
func (bg *BaseGateway) GetType() GatewayType {
	return bg.typ
}

// IsConnected checks if connected
func (bg *BaseGateway) IsConnected() bool {
	bg.mu.RLock()
	defer bg.mu.RUnlock()
	return bg.connected
}

// SetConnected sets the connected state
func (bg *BaseGateway) SetConnected(connected bool) {
	bg.mu.Lock()
	defer bg.mu.Unlock()
	bg.connected = connected
}

// GateGateway implements Gate.io gateway
type GateGateway struct {
	*BaseGateway
	apiKey    string
	apiSecret string
	baseURL   string
}

// NewGateGateway creates a new Gate.io gateway
func NewGateGateway(apiKey, apiSecret string) *GateGateway {
	return &GateGateway{
		BaseGateway: NewBaseGateway(GatewayTypeGate),
		apiKey:      apiKey,
		apiSecret:   apiSecret,
		baseURL:     "https://api.gateio.ws",
	}
}

func (gg *GateGateway) Connect(ctx context.Context) error {
	gg.SetConnected(true)
	return nil
}

func (gg *GateGateway) Disconnect(ctx context.Context) error {
	gg.SetConnected(false)
	return nil
}

func (gg *GateGateway) GetOrderBook(ctx context.Context, symbol string) (*OrderBook, error) {
	if !gg.IsConnected() {
		return nil, fmt.Errorf("Gate gateway not connected")
	}
	return &OrderBook{Symbol: symbol, Bids: []PriceLevel{}, Asks: []PriceLevel{}}, nil
}

func (gg *GateGateway) GetTrades(ctx context.Context, symbol string, limit int) ([]*Trade, error) {
	if !gg.IsConnected() {
		return nil, fmt.Errorf("Gate gateway not connected")
	}
	return []*Trade{}, nil
}

func (gg *GateGateway) GetKlines(ctx context.Context, symbol, interval string, start, end time.Time) ([]*Kline, error) {
	if !gg.IsConnected() {
		return nil, fmt.Errorf("Gate gateway not connected")
	}
	return []*Kline{}, nil
}

func (gg *GateGateway) SubscribeOrderBook(ctx context.Context, symbol string, handler OrderBookHandler) error {
	if !gg.IsConnected() {
		return fmt.Errorf("Gate gateway not connected")
	}
	gg.mu.Lock()
	defer gg.mu.Unlock()
	gg.obsHandlers[symbol] = append(gg.obsHandlers[symbol], handler)
	return nil
}

func (gg *GateGateway) SubscribeTrades(ctx context.Context, symbol string, handler TradeHandler) error {
	if !gg.IsConnected() {
		return fmt.Errorf("Gate gateway not connected")
	}
	gg.mu.Lock()
	defer gg.mu.Unlock()
	gg.tradeHandlers[symbol] = append(gg.tradeHandlers[symbol], handler)
	return nil
}

func (gg *GateGateway) SubscribeKlines(ctx context.Context, symbol, interval string, handler KlineHandler) error {
	if !gg.IsConnected() {
		return fmt.Errorf("Gate gateway not connected")
	}
	gg.mu.Lock()
	defer gg.mu.Unlock()
	key := symbol + ":" + interval
	gg.klineHandlers[key] = append(gg.klineHandlers[key], handler)
	return nil
}

// FreedxGateway implements Freedx gateway
type FreedxGateway struct {
	*BaseGateway
	apiKey    string
	apiSecret string
	baseURL   string
}

// NewFreedxGateway creates a new Freedx gateway
func NewFreedxGateway(apiKey, apiSecret string) *FreedxGateway {
	return &FreedxGateway{
		BaseGateway: NewBaseGateway(GatewayTypeFreedx),
		apiKey:      apiKey,
		apiSecret:   apiSecret,
		baseURL:     "https://api.freedx.io",
	}
}

func (fg *FreedxGateway) Connect(ctx context.Context) error {
	fg.SetConnected(true)
	return nil
}

func (fg *FreedxGateway) Disconnect(ctx context.Context) error {
	fg.SetConnected(false)
	return nil
}

func (fg *FreedxGateway) GetOrderBook(ctx context.Context, symbol string) (*OrderBook, error) {
	if !fg.IsConnected() {
		return nil, fmt.Errorf("Freedx gateway not connected")
	}
	return &OrderBook{Symbol: symbol, Bids: []PriceLevel{}, Asks: []PriceLevel{}}, nil
}

func (fg *FreedxGateway) GetTrades(ctx context.Context, symbol string, limit int) ([]*Trade, error) {
	if !fg.IsConnected() {
		return nil, fmt.Errorf("Freedx gateway not connected")
	}
	return []*Trade{}, nil
}

func (fg *FreedxGateway) GetKlines(ctx context.Context, symbol, interval string, start, end time.Time) ([]*Kline, error) {
	if !fg.IsConnected() {
		return nil, fmt.Errorf("Freedx gateway not connected")
	}
	return []*Kline{}, nil
}

func (fg *FreedxGateway) SubscribeOrderBook(ctx context.Context, symbol string, handler OrderBookHandler) error {
	if !fg.IsConnected() {
		return fmt.Errorf("Freedx gateway not connected")
	}
	fg.mu.Lock()
	defer fg.mu.Unlock()
	fg.obsHandlers[symbol] = append(fg.obsHandlers[symbol], handler)
	return nil
}

func (fg *FreedxGateway) SubscribeTrades(ctx context.Context, symbol string, handler TradeHandler) error {
	if !fg.IsConnected() {
		return fmt.Errorf("Freedx gateway not connected")
	}
	fg.mu.Lock()
	defer fg.mu.Unlock()
	fg.tradeHandlers[symbol] = append(fg.tradeHandlers[symbol], handler)
	return nil
}

func (fg *FreedxGateway) SubscribeKlines(ctx context.Context, symbol, interval string, handler KlineHandler) error {
	if !fg.IsConnected() {
		return fmt.Errorf("Freedx gateway not connected")
	}
	fg.mu.Lock()
	defer fg.mu.Unlock()
	key := symbol + ":" + interval
	fg.klineHandlers[key] = append(fg.klineHandlers[key], handler)
	return nil
}

// Registry manages gateways
type Registry struct {
	mu       sync.RWMutex
	gateways map[string]Gateway
}

// NewRegistry creates a new gateway registry
func NewRegistry() *Registry {
	return &Registry{
		gateways: make(map[string]Gateway),
	}
}

// Register registers a gateway
func (r *Registry) Register(name string, gateway Gateway) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.gateways[name] = gateway
	return nil
}

// Get retrieves a gateway
func (r *Registry) Get(name string) (Gateway, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	gateway, exists := r.gateways[name]
	return gateway, exists
}

// List returns all registered gateways
func (r *Registry) List() []Gateway {
	r.mu.RLock()
	defer r.mu.RUnlock()
	gateways := make([]Gateway, 0, len(r.gateways))
	for _, g := range r.gateways {
		gateways = append(gateways, g)
	}
	return gateways
}
