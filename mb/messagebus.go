package mb

import (
	"context"
	"fmt"
	"sync"
)

// MessageBusType represents the type of message bus
type MessageBusType string

const (
	MessageBusTypeZMQ  MessageBusType = "zmq"
	MessageBusTypeMQTT MessageBusType = "mqtt"
)

// MessageBus defines the interface for message bus implementations
type MessageBus interface {
	// Connect establishes connection to the message bus
	Connect(ctx context.Context) error

	// Disconnect closes the connection
	Disconnect(ctx context.Context) error

	// Publish publishes a message to a topic
	Publish(ctx context.Context, topic string, message []byte) error

	// Subscribe subscribes to a topic
	Subscribe(ctx context.Context, topic string, handler MessageHandler) error

	// Unsubscribe unsubscribes from a topic
	Unsubscribe(ctx context.Context, topic string) error

	// GetType returns the message bus type
	GetType() MessageBusType

	// IsConnected checks if the bus is connected
	IsConnected() bool
}

// MessageHandler is the callback function for received messages
type MessageHandler func(topic string, message []byte) error

// BaseMessageBus provides common functionality
type BaseMessageBus struct {
	typ       MessageBusType
	connected bool
	mu        sync.RWMutex
	handlers  map[string][]MessageHandler
	config    map[string]interface{}
}

// NewBaseMessageBus creates a new base message bus
func NewBaseMessageBus(typ MessageBusType) *BaseMessageBus {
	return &BaseMessageBus{
		typ:       typ,
		connected: false,
		handlers:  make(map[string][]MessageHandler),
		config:    make(map[string]interface{}),
	}
}

// GetType returns the message bus type
func (bmb *BaseMessageBus) GetType() MessageBusType {
	return bmb.typ
}

// IsConnected checks if the bus is connected
func (bmb *BaseMessageBus) IsConnected() bool {
	bmb.mu.RLock()
	defer bmb.mu.RUnlock()
	return bmb.connected
}

// SetConnected sets the connected state
func (bmb *BaseMessageBus) SetConnected(connected bool) {
	bmb.mu.Lock()
	defer bmb.mu.Unlock()
	bmb.connected = connected
}

// AddHandler adds a message handler for a topic
func (bmb *BaseMessageBus) AddHandler(topic string, handler MessageHandler) {
	bmb.mu.Lock()
	defer bmb.mu.Unlock()
	bmb.handlers[topic] = append(bmb.handlers[topic], handler)
}

// GetHandlers returns all handlers for a topic
func (bmb *BaseMessageBus) GetHandlers(topic string) []MessageHandler {
	bmb.mu.RLock()
	defer bmb.mu.RUnlock()
	return bmb.handlers[topic]
}

// RemoveHandlers removes all handlers for a topic
func (bmb *BaseMessageBus) RemoveHandlers(topic string) {
	bmb.mu.Lock()
	defer bmb.mu.Unlock()
	delete(bmb.handlers, topic)
}

// ZMQBus implements ZeroMQ message bus
type ZMQBus struct {
	*BaseMessageBus
	endpoint string
}

// NewZMQBus creates a new ZMQ message bus
func NewZMQBus(endpoint string) *ZMQBus {
	return &ZMQBus{
		BaseMessageBus: NewBaseMessageBus(MessageBusTypeZMQ),
		endpoint:       endpoint,
	}
}

func (zb *ZMQBus) Connect(ctx context.Context) error {
	// ZMQ connection implementation
	zb.SetConnected(true)
	return nil
}

func (zb *ZMQBus) Disconnect(ctx context.Context) error {
	zb.SetConnected(false)
	return nil
}

func (zb *ZMQBus) Publish(ctx context.Context, topic string, message []byte) error {
	if !zb.IsConnected() {
		return fmt.Errorf("ZMQ bus is not connected")
	}
	// Publish implementation
	return nil
}

func (zb *ZMQBus) Subscribe(ctx context.Context, topic string, handler MessageHandler) error {
	if !zb.IsConnected() {
		return fmt.Errorf("ZMQ bus is not connected")
	}
	zb.AddHandler(topic, handler)
	return nil
}

func (zb *ZMQBus) Unsubscribe(ctx context.Context, topic string) error {
	zb.RemoveHandlers(topic)
	return nil
}

// MQTTBus implements MQTT message bus
type MQTTBus struct {
	*BaseMessageBus
	brokerURL string
	clientID  string
}

// NewMQTTBus creates a new MQTT message bus
func NewMQTTBus(brokerURL, clientID string) *MQTTBus {
	return &MQTTBus{
		BaseMessageBus: NewBaseMessageBus(MessageBusTypeMQTT),
		brokerURL:      brokerURL,
		clientID:       clientID,
	}
}

func (mb *MQTTBus) Connect(ctx context.Context) error {
	// MQTT connection implementation
	mb.SetConnected(true)
	return nil
}

func (mb *MQTTBus) Disconnect(ctx context.Context) error {
	mb.SetConnected(false)
	return nil
}

func (mb *MQTTBus) Publish(ctx context.Context, topic string, message []byte) error {
	if !mb.IsConnected() {
		return fmt.Errorf("MQTT bus is not connected")
	}
	// Publish implementation
	return nil
}

func (mb *MQTTBus) Subscribe(ctx context.Context, topic string, handler MessageHandler) error {
	if !mb.IsConnected() {
		return fmt.Errorf("MQTT bus is not connected")
	}
	mb.AddHandler(topic, handler)
	return nil
}

func (mb *MQTTBus) Unsubscribe(ctx context.Context, topic string) error {
	mb.RemoveHandlers(topic)
	return nil
}

// Registry manages message buses
type Registry struct {
	mu  sync.RWMutex
	bus map[string]MessageBus
}

// NewRegistry creates a new message bus registry
func NewRegistry() *Registry {
	return &Registry{
		bus: make(map[string]MessageBus),
	}
}

// Register registers a message bus
func (r *Registry) Register(name string, bus MessageBus) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.bus[name] = bus
	return nil
}

// Get retrieves a message bus
func (r *Registry) Get(name string) (MessageBus, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	bus, exists := r.bus[name]
	return bus, exists
}

// List returns all registered buses
func (r *Registry) List() []MessageBus {
	r.mu.RLock()
	defer r.mu.RUnlock()
	buses := make([]MessageBus, 0, len(r.bus))
	for _, b := range r.bus {
		buses = append(buses, b)
	}
	return buses
}
