package mcp

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
)

// MCPRequest represents an MCP request
type MCPRequest struct {
	ID      string          `json:"id"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params,omitempty"`
	Version string          `json:"jsonrpc"`
}

// MCPResponse represents an MCP response
type MCPResponse struct {
	ID      string          `json:"id"`
	Result  json.RawMessage `json:"result,omitempty"`
	Error   *MCPError       `json:"error,omitempty"`
	Version string          `json:"jsonrpc"`
}

// MCPError represents an MCP error
type MCPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data,omitempty"`
}

// RequestHandler is the callback for handling MCP requests
type RequestHandler func(ctx context.Context, req *MCPRequest) (*MCPResponse, error)

// ServerCapabilities describes server capabilities
type ServerCapabilities struct {
	Sampling       bool `json:"sampling"`
	RootsChanged   bool `json:"roots_changed"`
	ResourceUpdate bool `json:"resource_update"`
}

// Server implements an MCP server
type Server struct {
	mu       sync.RWMutex
	handlers map[string]RequestHandler
	caps     ServerCapabilities
}

// NewServer creates a new MCP server
func NewServer(caps ServerCapabilities) *Server {
	return &Server{
		handlers: make(map[string]RequestHandler),
		caps:     caps,
	}
}

// RegisterHandler registers a request handler
func (s *Server) RegisterHandler(method string, handler RequestHandler) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers[method] = handler
}

// HandleRequest handles an MCP request
func (s *Server) HandleRequest(ctx context.Context, req *MCPRequest) (*MCPResponse, error) {
	s.mu.RLock()
	handler, exists := s.handlers[req.Method]
	s.mu.RUnlock()

	if !exists {
		return &MCPResponse{
			ID:      req.ID,
			Version: "2.0",
			Error: &MCPError{
				Code:    -32601,
				Message: "Method not found",
			},
		}, nil
	}

	return handler(ctx, req)
}

// Client implements an MCP client
type Client struct {
	mu           sync.RWMutex
	capabilities ServerCapabilities
}

// NewClient creates a new MCP client
func NewClient() *Client {
	return &Client{
		capabilities: ServerCapabilities{},
	}
}

// CreateRequest creates a new MCP request
func (c *Client) CreateRequest(id, method string, params interface{}) (*MCPRequest, error) {
	var paramData json.RawMessage
	if params != nil {
		data, err := json.Marshal(params)
		if err != nil {
			return nil, err
		}
		paramData = data
	}

	return &MCPRequest{
		ID:      id,
		Method:  method,
		Params:  paramData,
		Version: "2.0",
	}, nil
}

// ServerConfig represents server configuration
type ServerConfig struct {
	Name           string
	Version        string
	Capabilities   ServerCapabilities
	MaxMessageSize int
	Timeout        int
}

// Transport defines the message transport interface
type Transport interface {
	// Send sends an MCP request/response
	Send(ctx context.Context, data interface{}) error

	// Receive receives an MCP request/response
	Receive(ctx context.Context) (interface{}, error)

	// Close closes the transport
	Close() error
}

// Proxy acts as a proxy between MCPclient and server
type Proxy struct {
	mu        sync.RWMutex
	server    *Server
	transport Transport
}

// NewProxy creates a new MCP proxy
func NewProxy(server *Server, transport Transport) *Proxy {
	return &Proxy{
		server:    server,
		transport: transport,
	}
}

// Start starts the proxy
func (p *Proxy) Start(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			// Receive request
			msg, err := p.transport.Receive(ctx)
			if err != nil {
				return err
			}

			// Convert to MCPRequest
			var req MCPRequest
			if err := json.Unmarshal(msg.([]byte), &req); err != nil {
				return err
			}

			// Handle request
			resp, err := p.server.HandleRequest(ctx, &req)
			if err != nil {
				resp = &MCPResponse{
					ID:      req.ID,
					Version: "2.0",
					Error: &MCPError{
						Code:    -32603,
						Message: "Internal error",
						Data:    fmt.Sprintf("%v", err),
					},
				}
			}

			// Send response
			respData, err := json.Marshal(resp)
			if err != nil {
				return err
			}

			if err := p.transport.Send(ctx, respData); err != nil {
				return err
			}
		}
	}
}
