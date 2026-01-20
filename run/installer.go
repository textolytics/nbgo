package run

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"sync"
	"time"
)

// InstallationTarget represents a target for installation
type InstallationTarget struct {
	ID           string
	Name         string
	Description  string
	Dependencies []string
	Type         string // "provider", "service", "tool"
	Version      string
	Enabled      bool
	Installed    bool
	Status       string
	LastAttempt  time.Time
}

// InstallationResult tracks installation results
type InstallationResult struct {
	TargetID     string
	Success      bool
	Error        string
	Output       string
	Duration     time.Duration
	Timestamp    time.Time
	Dependencies []string
	Status       string
}

// InstallationManager manages installation of providers and services
type InstallationManager struct {
	mu           sync.RWMutex
	targets      map[string]*InstallationTarget
	results      []*InstallationResult
	installers   map[string]InstallerFunc
	installOrder []string
	ctx          context.Context
	cancel       context.CancelFunc
}

// InstallerFunc is a function that installs a target
type InstallerFunc func(ctx context.Context, target *InstallationTarget) error

// NewInstallationManager creates a new installation manager
func NewInstallationManager() *InstallationManager {
	ctx, cancel := context.WithCancel(context.Background())
	return &InstallationManager{
		targets:      make(map[string]*InstallationTarget),
		results:      make([]*InstallationResult, 0),
		installers:   make(map[string]InstallerFunc),
		installOrder: make([]string, 0),
		ctx:          ctx,
		cancel:       cancel,
	}
}

// RegisterTarget registers an installation target
func (im *InstallationManager) RegisterTarget(target *InstallationTarget) {
	im.mu.Lock()
	defer im.mu.Unlock()
	im.targets[target.ID] = target
	if target.Enabled {
		im.installOrder = append(im.installOrder, target.ID)
	}
}

// RegisterInstaller registers an installer function for a target
func (im *InstallationManager) RegisterInstaller(targetID string, installer InstallerFunc) {
	im.mu.Lock()
	defer im.mu.Unlock()
	im.installers[targetID] = installer
}

// Install installs a single target
func (im *InstallationManager) Install(targetID string) (*InstallationResult, error) {
	im.mu.RLock()
	target, exists := im.targets[targetID]
	installer, installerExists := im.installers[targetID]
	im.mu.RUnlock()

	if !exists {
		return nil, fmt.Errorf("target not found: %s", targetID)
	}

	result := &InstallationResult{
		TargetID:     targetID,
		Timestamp:    time.Now(),
		Dependencies: target.Dependencies,
	}

	startTime := time.Now()

	// Check dependencies
	for _, dep := range target.Dependencies {
		im.mu.RLock()
		depTarget, depExists := im.targets[dep]
		im.mu.RUnlock()

		if !depExists || !depTarget.Installed {
			result.Success = false
			result.Error = fmt.Sprintf("dependency not installed: %s", dep)
			result.Duration = time.Since(startTime)
			result.Status = "failed"
			return result, nil
		}
	}

	// Check if installer is registered
	if !installerExists {
		result.Success = false
		result.Error = "no installer registered"
		result.Duration = time.Since(startTime)
		result.Status = "failed"
		return result, nil
	}

	// Execute installation
	if err := installer(im.ctx, target); err != nil {
		result.Success = false
		result.Error = err.Error()
		result.Status = "failed"
	} else {
		result.Success = true
		result.Status = "installed"
		im.mu.Lock()
		target.Installed = true
		target.Status = "installed"
		target.LastAttempt = time.Now()
		im.mu.Unlock()
	}

	result.Duration = time.Since(startTime)

	// Store result
	im.mu.Lock()
	im.results = append(im.results, result)
	im.mu.Unlock()

	return result, nil
}

// InstallAll installs all enabled targets in dependency order
func (im *InstallationManager) InstallAll() ([]*InstallationResult, error) {
	results := make([]*InstallationResult, 0)

	im.mu.RLock()
	order := make([]string, len(im.installOrder))
	copy(order, im.installOrder)
	im.mu.RUnlock()

	for _, targetID := range order {
		select {
		case <-im.ctx.Done():
			return results, im.ctx.Err()
		default:
			result, err := im.Install(targetID)
			if err != nil {
				return results, fmt.Errorf("install %s failed: %w", targetID, err)
			}
			results = append(results, result)

			if !result.Success {
				return results, fmt.Errorf("install %s failed: %s", targetID, result.Error)
			}
		}
	}

	return results, nil
}

// Uninstall uninstalls a target
func (im *InstallationManager) Uninstall(targetID string) error {
	im.mu.Lock()
	defer im.mu.Unlock()

	target, exists := im.targets[targetID]
	if !exists {
		return fmt.Errorf("target not found: %s", targetID)
	}

	// Check if other targets depend on this one
	for _, other := range im.targets {
		for _, dep := range other.Dependencies {
			if dep == targetID && other.Installed {
				return fmt.Errorf("target %s is required by %s", targetID, other.ID)
			}
		}
	}

	target.Installed = false
	target.Status = "uninstalled"
	return nil
}

// GetTarget retrieves a target
func (im *InstallationManager) GetTarget(targetID string) (*InstallationTarget, bool) {
	im.mu.RLock()
	defer im.mu.RUnlock()
	target, exists := im.targets[targetID]
	return target, exists
}

// ListTargets returns all targets
func (im *InstallationManager) ListTargets() []*InstallationTarget {
	im.mu.RLock()
	defer im.mu.RUnlock()
	targets := make([]*InstallationTarget, 0, len(im.targets))
	for _, target := range im.targets {
		targets = append(targets, target)
	}
	return targets
}

// GetResults returns installation results
func (im *InstallationManager) GetResults() []*InstallationResult {
	im.mu.RLock()
	defer im.mu.RUnlock()
	return im.results
}

// Stop stops any ongoing installation
func (im *InstallationManager) Stop() {
	im.cancel()
}

// CommandInstaller creates an installer that executes a shell command
func CommandInstaller(cmd string, args ...string) InstallerFunc {
	return func(ctx context.Context, target *InstallationTarget) error {
		command := exec.CommandContext(ctx, cmd, args...)
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr
		return command.Run()
	}
}

// ShellInstaller creates an installer that executes a shell script
func ShellInstaller(script string) InstallerFunc {
	return func(ctx context.Context, target *InstallationTarget) error {
		command := exec.CommandContext(ctx, "sh", "-c", script)
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr
		return command.Run()
	}
}

// HealthCheck checks if installation is valid
type HealthCheck struct {
	TargetID  string
	Name      string
	CheckFunc func(ctx context.Context) error
	Interval  time.Duration
}

// HealthCheckMonitor monitors health of installations
type HealthCheckMonitor struct {
	mu      sync.RWMutex
	checks  map[string]*HealthCheck
	results map[string]error
	ticker  *time.Ticker
	done    chan bool
}

// NewHealthCheckMonitor creates a new health check monitor
func NewHealthCheckMonitor(interval time.Duration) *HealthCheckMonitor {
	return &HealthCheckMonitor{
		checks:  make(map[string]*HealthCheck),
		results: make(map[string]error),
		ticker:  time.NewTicker(interval),
		done:    make(chan bool),
	}
}

// RegisterCheck registers a health check
func (hcm *HealthCheckMonitor) RegisterCheck(check *HealthCheck) {
	hcm.mu.Lock()
	defer hcm.mu.Unlock()
	hcm.checks[check.TargetID] = check
}

// Start starts health checking
func (hcm *HealthCheckMonitor) Start(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				hcm.Stop()
				return
			case <-hcm.done:
				return
			case <-hcm.ticker.C:
				hcm.checkAll(ctx)
			}
		}
	}()
}

// checkAll runs all health checks
func (hcm *HealthCheckMonitor) checkAll(ctx context.Context) {
	hcm.mu.RLock()
	checks := make(map[string]*HealthCheck)
	for k, v := range hcm.checks {
		checks[k] = v
	}
	hcm.mu.RUnlock()

	for targetID, check := range checks {
		if err := check.CheckFunc(ctx); err != nil {
			hcm.mu.Lock()
			hcm.results[targetID] = err
			hcm.mu.Unlock()
		} else {
			hcm.mu.Lock()
			hcm.results[targetID] = nil
			hcm.mu.Unlock()
		}
	}
}

// GetResults returns health check results
func (hcm *HealthCheckMonitor) GetResults() map[string]error {
	hcm.mu.RLock()
	defer hcm.mu.RUnlock()
	results := make(map[string]error)
	for k, v := range hcm.results {
		results[k] = v
	}
	return results
}

// Stop stops health checking
func (hcm *HealthCheckMonitor) Stop() {
	hcm.ticker.Stop()
	select {
	case hcm.done <- true:
	default:
	}
}
