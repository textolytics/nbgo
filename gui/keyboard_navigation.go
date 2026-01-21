package gui
package gui

import (
	"sync"
)

// KeyboardEvent represents a keyboard event
type KeyboardEvent struct {
	Key       string
	Modifiers []string // ctrl, shift, alt
	Timestamp int64
}

// KeyboardHandler handles keyboard navigation
type KeyboardHandler struct {
	mu           sync.RWMutex
	keyBindings  map[string]KeyBinding
	focusedView  string
	viewStack    []string
}

// KeyBinding represents a key binding
type KeyBinding struct {
	Key         string
	Modifiers   []string
	Description string
	Handler     func() error
	Category    string
}

// NewKeyboardHandler creates a new keyboard handler
func NewKeyboardHandler() *KeyboardHandler {
	return &KeyboardHandler{
		keyBindings: make(map[string]KeyBinding),
		viewStack:   make([]string, 0),
	}
}

// RegisterKeyBinding registers a key binding
func (kh *KeyboardHandler) RegisterKeyBinding(binding KeyBinding) {
	kh.mu.Lock()
	defer kh.mu.Unlock()
	kh.keyBindings[binding.Key] = binding
}

// HandleKeyEvent handles a keyboard event
func (kh *KeyboardHandler) HandleKeyEvent(event KeyboardEvent) error {
	kh.mu.RLock()
	binding, exists := kh.keyBindings[event.Key]
	kh.mu.RUnlock()

	if !exists {
		return nil
	}

	if binding.Handler != nil {
		return binding.Handler()
	}
	return nil
}

// SetFocusedView sets the focused view
func (kh *KeyboardHandler) SetFocusedView(viewName string) {
	kh.mu.Lock()
	defer kh.mu.Unlock()
	kh.focusedView = viewName
}

// GetFocusedView returns the focused view
func (kh *KeyboardHandler) GetFocusedView() string {
	kh.mu.RLock()
	defer kh.mu.RUnlock()
	return kh.focusedView
}

// PushView pushes a view onto the stack
func (kh *KeyboardHandler) PushView(viewName string) {
	kh.mu.Lock()
	defer kh.mu.Unlock()
	kh.viewStack = append(kh.viewStack, viewName)
	kh.focusedView = viewName
}

// PopView pops a view from the stack
func (kh *KeyboardHandler) PopView() string {
	kh.mu.Lock()
	defer kh.mu.Unlock()

	if len(kh.viewStack) == 0 {
		return ""
	}

	view := kh.viewStack[len(kh.viewStack)-1]
	kh.viewStack = kh.viewStack[:len(kh.viewStack)-1]

	if len(kh.viewStack) > 0 {
		kh.focusedView = kh.viewStack[len(kh.viewStack)-1]
	}

	return view
}

// ListKeyBindings returns all key bindings
func (kh *KeyboardHandler) ListKeyBindings() map[string]KeyBinding {
	kh.mu.RLock()
	defer kh.mu.RUnlock()

	bindings := make(map[string]KeyBinding)
	for k, v := range kh.keyBindings {
		bindings[k] = v
	}
	return bindings
}

// NavigationMode represents different navigation modes
type NavigationMode string

const (
	NavigationModeNormal  NavigationMode = "normal"
	NavigationModeInsert  NavigationMode = "insert"
	NavigationModeCommand NavigationMode = "command"
	NavigationModeSearch  NavigationMode = "search"
)

// NavigationController handles navigation between views
type NavigationController struct {
	mu              sync.RWMutex
	currentMode     NavigationMode
	currentView     string
	navigationStack []string
	history         []string
	historyIndex    int
}

// NewNavigationController creates a new navigation controller
func NewNavigationController() *NavigationController {
	return &NavigationController{
		currentMode:     NavigationModeNormal,
		navigationStack: make([]string, 0),
		history:         make([]string, 0),
		historyIndex:    -1,
	}
}

// Navigate navigates to a view
func (nc *NavigationController) Navigate(viewName string) error {
	nc.mu.Lock()
	defer nc.mu.Unlock()

	nc.currentView = viewName
	nc.navigationStack = append(nc.navigationStack, viewName)
	nc.history = append(nc.history, viewName)
	nc.historyIndex = len(nc.history) - 1

	return nil
}

// GoBack goes back to the previous view
func (nc *NavigationController) GoBack() error {
	nc.mu.Lock()
	defer nc.mu.Unlock()

	if len(nc.navigationStack) <= 1 {
		return nil
	}

	nc.navigationStack = nc.navigationStack[:len(nc.navigationStack)-1]
	nc.currentView = nc.navigationStack[len(nc.navigationStack)-1]

	return nil
}

// GoForward goes forward in history
func (nc *NavigationController) GoForward() error {
	nc.mu.Lock()
	defer nc.mu.Unlock()

	if nc.historyIndex >= len(nc.history)-1 {
		return nil
	}

	nc.historyIndex++
	nc.currentView = nc.history[nc.historyIndex]

	return nil
}

// SetMode sets the navigation mode
func (nc *NavigationController) SetMode(mode NavigationMode) {
	nc.mu.Lock()
	defer nc.mu.Unlock()
	nc.currentMode = mode
}

// GetMode returns the current navigation mode
func (nc *NavigationController) GetMode() NavigationMode {
	nc.mu.RLock()
	defer nc.mu.RUnlock()
	return nc.currentMode
}

// GetCurrentView returns the current view
func (nc *NavigationController) GetCurrentView() string {
	nc.mu.RLock()
	defer nc.mu.RUnlock()
	return nc.currentView
}

// GetNavigationHistory returns the navigation history
func (nc *NavigationController) GetNavigationHistory() []string {
	nc.mu.RLock()
	defer nc.mu.RUnlock()

	history := make([]string, len(nc.history))
	copy(history, nc.history)
	return history
}
