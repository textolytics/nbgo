package gui

import (
	"sync"
)

// UISession represents a UI session with multiple windows
type UISession struct {
	mu            sync.RWMutex
	ID            string
	Windows       map[string]*Window
	Environment   map[string]string
	Active        bool
	CreatedAt     int64
	LastActivity  int64
	CommandBuffer []string
}

// Window represents a UI window
type Window struct {
	mu          sync.RWMutex
	ID          string
	Title       string
	View        *View
	X           int
	Y           int
	Width       int
	Height      int
	Active      bool
	ZIndex      int
	Content     string
}

// NewWindow creates a new window
func NewWindow(id, title string, view *View) *Window {
	return &Window{
		ID:      id,
		Title:   title,
		View:    view,
		Active:  true,
		ZIndex:  0,
		Content: "",
	}
}

// SetPosition sets the window position
func (w *Window) SetPosition(x, y int) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.X = x
	w.Y = y
}

// SetSize sets the window size
func (w *Window) SetSize(width, height int) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.Width = width
	w.Height = height
}

// GetBounds returns the window bounds
func (w *Window) GetBounds() (int, int, int, int) {
	w.mu.RLock()
	defer w.mu.RUnlock()
	return w.X, w.Y, w.Width, w.Height
}

// SetActive sets the window active state
func (w *Window) SetActive(active bool) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.Active = active
}

// IsActive returns if the window is active
func (w *Window) IsActive() bool {
	w.mu.RLock()
	defer w.mu.RUnlock()
	return w.Active
}

// UpdateContent updates the window content
func (w *Window) UpdateContent(content string) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.Content = content
	if w.View != nil {
		w.View.UpdateContent(content)
	}
}

// Minimize minimizes the window
func (w *Window) Minimize() {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.Active = false
}

// Maximize maximizes the window
func (w *Window) Maximize() {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.Width = 100
	w.Height = 30
	w.X = 0
	w.Y = 0
}

// Focus sets focus to the window
func (w *Window) Focus() {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.Active = true
	w.ZIndex = 100
}

// AddWindow adds a window to the session
func (s *UISession) AddWindow(window *Window) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Windows[window.ID] = window
}

// RemoveWindow removes a window from the session
func (s *UISession) RemoveWindow(windowID string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.Windows, windowID)
}

// GetWindow retrieves a window
func (s *UISession) GetWindow(windowID string) *Window {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.Windows[windowID]
}

// ListWindows returns all windows
func (s *UISession) ListWindows() map[string]*Window {
	s.mu.RLock()
	defer s.mu.RUnlock()

	windows := make(map[string]*Window)
	for k, v := range s.Windows {
		windows[k] = v
	}
	return windows
}

// SetEnvironmentVariable sets an environment variable in the session
func (s *UISession) SetEnvironmentVariable(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Environment[key] = value
}

// GetEnvironmentVariable gets an environment variable
func (s *UISession) GetEnvironmentVariable(key string) string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.Environment[key]
}

// GetAllEnvironmentVariables returns all environment variables
func (s *UISession) GetAllEnvironmentVariables() map[string]string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	env := make(map[string]string)
	for k, v := range s.Environment {
		env[k] = v
	}
	return env
}

// AddCommand adds a command to the session buffer
func (s *UISession) AddCommand(cmd string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.CommandBuffer = append(s.CommandBuffer, cmd)
}

// GetCommandBuffer returns the command buffer
func (s *UISession) GetCommandBuffer() []string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	cmds := make([]string, len(s.CommandBuffer))
	copy(cmds, s.CommandBuffer)
	return cmds
}

// ClearCommandBuffer clears the command buffer
func (s *UISession) ClearCommandBuffer() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.CommandBuffer = make([]string, 0)
}

// IsActive returns if the session is active
func (s *UISession) IsActive() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.Active
}

// Close closes the session
func (s *UISession) Close() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Active = false
}
