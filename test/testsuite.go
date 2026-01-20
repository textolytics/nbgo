package test

import (
	"context"
	"sync"
	"testing"
	"time"
)

// TestCase represents a test case
type TestCase struct {
	Name        string
	Description string
	Setup       func(t *testing.T)
	Execute     func(t *testing.T)
	Teardown    func(t *testing.T)
	Timeout     time.Duration
}

// TestSuite manages a collection of test cases
type TestSuite struct {
	mu        sync.RWMutex
	name      string
	testCases []*TestCase
	results   map[string]*TestResult
}

// TestResult contains the result of a test
type TestResult struct {
	Name      string
	Passed    bool
	Error     error
	Duration  time.Duration
	Timestamp time.Time
}

// NewTestSuite creates a new test suite
func NewTestSuite(name string) *TestSuite {
	return &TestSuite{
		name:      name,
		testCases: make([]*TestCase, 0),
		results:   make(map[string]*TestResult),
	}
}

// AddTest adds a test case to the suite
func (ts *TestSuite) AddTest(testCase *TestCase) {
	ts.mu.Lock()
	defer ts.mu.Unlock()
	ts.testCases = append(ts.testCases, testCase)
}

// Run runs all tests in the suite
func (ts *TestSuite) Run(t *testing.T) {
	ts.mu.RLock()
	tests := make([]*TestCase, len(ts.testCases))
	copy(tests, ts.testCases)
	ts.mu.RUnlock()

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			ts.runTest(t, tc)
		})
	}
}

// runTest runs a single test case
func (ts *TestSuite) runTest(t *testing.T, tc *TestCase) {
	result := &TestResult{
		Name:      tc.Name,
		Timestamp: time.Now(),
	}

	// Setup
	if tc.Setup != nil {
		tc.Setup(t)
	}

	// Execute with timeout
	ctx, cancel := context.WithTimeout(context.Background(), tc.Timeout)
	defer cancel()

	start := time.Now()
	done := make(chan error, 1)

	go func() {
		defer func() {
			if r := recover(); r != nil {
				// Handle panic
				t.Errorf("Test panic: %v", r)
			}
		}()
		tc.Execute(t)
		done <- nil
	}()

	select {
	case <-ctx.Done():
		result.Error = ctx.Err()
		result.Passed = false
		t.Errorf("Test timeout exceeded")
	case err := <-done:
		result.Error = err
		result.Passed = err == nil && !t.Failed()
		result.Duration = time.Since(start)
	}

	// Teardown
	if tc.Teardown != nil {
		tc.Teardown(t)
	}

	ts.mu.Lock()
	ts.results[tc.Name] = result
	ts.mu.Unlock()
}

// GetResults returns all test results
func (ts *TestSuite) GetResults() []*TestResult {
	ts.mu.RLock()
	defer ts.mu.RUnlock()
	results := make([]*TestResult, 0, len(ts.results))
	for _, r := range ts.results {
		results = append(results, r)
	}
	return results
}

// IntegrationTestRunner runs integration tests
type IntegrationTestRunner struct {
	mu     sync.RWMutex
	suites map[string]*TestSuite
}

// NewIntegrationTestRunner creates a new integration test runner
func NewIntegrationTestRunner() *IntegrationTestRunner {
	return &IntegrationTestRunner{
		suites: make(map[string]*TestSuite),
	}
}

// AddSuite adds a test suite
func (itr *IntegrationTestRunner) AddSuite(suite *TestSuite) {
	itr.mu.Lock()
	defer itr.mu.Unlock()
	itr.suites[suite.name] = suite
}

// RunAll runs all test suites
func (itr *IntegrationTestRunner) RunAll(t *testing.T) {
	itr.mu.RLock()
	suites := make(map[string]*TestSuite)
	for k, v := range itr.suites {
		suites[k] = v
	}
	itr.mu.RUnlock()

	for name, suite := range suites {
		t.Run(name, func(t *testing.T) {
			suite.Run(t)
		})
	}
}
