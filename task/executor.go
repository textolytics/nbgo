package task

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// TaskStatus represents the status of a task
type TaskStatus string

const (
	TaskStatusPending   TaskStatus = "pending"
	TaskStatusRunning   TaskStatus = "running"
	TaskStatusCompleted TaskStatus = "completed"
	TaskStatusFailed    TaskStatus = "failed"
	TaskStatusCanceled  TaskStatus = "canceled"
)

// Task represents a task to be executed
type Task struct {
	ID          string
	Name        string
	Description string
	Status      TaskStatus
	StartTime   time.Time
	EndTime     time.Time
	Error       error
	Result      interface{}
	Priority    int
	Timeout     time.Duration
	Retries     int
	CurrentTry  int
	Handler     TaskHandler
}

// TaskHandler is the callback function for task execution
type TaskHandler func(ctx context.Context, task *Task) error

// Executor manages task execution
type Executor struct {
	mu             sync.RWMutex
	tasks          map[string]*Task
	running        map[string]*Task
	queue          []*Task
	workers        int
	maxRetries     int
	defaultTimeout time.Duration
}

// NewExecutor creates a new task executor
func NewExecutor(workers int, maxRetries int, defaultTimeout time.Duration) *Executor {
	return &Executor{
		tasks:          make(map[string]*Task),
		running:        make(map[string]*Task),
		queue:          make([]*Task, 0),
		workers:        workers,
		maxRetries:     maxRetries,
		defaultTimeout: defaultTimeout,
	}
}

// Submit submits a task for execution
func (e *Executor) Submit(task *Task) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	if _, exists := e.tasks[task.ID]; exists {
		return fmt.Errorf("task %s already exists", task.ID)
	}

	if task.Timeout == 0 {
		task.Timeout = e.defaultTimeout
	}

	task.Status = TaskStatusPending
	e.tasks[task.ID] = task
	e.queue = append(e.queue, task)

	return nil
}

// Get retrieves a task by ID
func (e *Executor) Get(id string) (*Task, bool) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	task, exists := e.tasks[id]
	return task, exists
}

// List returns all tasks
func (e *Executor) List() []*Task {
	e.mu.RLock()
	defer e.mu.RUnlock()
	tasks := make([]*Task, 0, len(e.tasks))
	for _, task := range e.tasks {
		tasks = append(tasks, task)
	}
	return tasks
}

// Execute executes a task
func (e *Executor) Execute(ctx context.Context, task *Task) error {
	e.mu.Lock()
	e.running[task.ID] = task
	e.mu.Unlock()

	defer func() {
		e.mu.Lock()
		delete(e.running, task.ID)
		e.mu.Unlock()
	}()

	task.Status = TaskStatusRunning
	task.StartTime = time.Now()

	// Create context with timeout
	execCtx, cancel := context.WithTimeout(ctx, task.Timeout)
	defer cancel()

	// Execute with retry logic
	var err error
	for attempt := 0; attempt <= task.Retries; attempt++ {
		task.CurrentTry = attempt + 1
		err = task.Handler(execCtx, task)
		if err == nil {
			break
		}
		if attempt < task.Retries {
			select {
			case <-time.After(time.Second * time.Duration(attempt+1)):
				// Exponential backoff
			case <-execCtx.Done():
				err = execCtx.Err()
				break
			}
		}
	}

	task.EndTime = time.Now()

	if err != nil {
		task.Status = TaskStatusFailed
		task.Error = err
	} else {
		task.Status = TaskStatusCompleted
	}

	return err
}

// Cancel cancels a task
func (e *Executor) Cancel(id string) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	task, exists := e.tasks[id]
	if !exists {
		return fmt.Errorf("task %s not found", id)
	}

	if task.Status == TaskStatusRunning {
		task.Status = TaskStatusCanceled
		return nil
	}

	if task.Status == TaskStatusPending {
		task.Status = TaskStatusCanceled
		// Remove from queue
		for i, t := range e.queue {
			if t.ID == id {
				e.queue = append(e.queue[:i], e.queue[i+1:]...)
				break
			}
		}
		return nil
	}

	return fmt.Errorf("cannot cancel task in status: %s", task.Status)
}

// Scenario represents a sequence of tasks
type Scenario struct {
	ID       string
	Name     string
	Tasks    []*Task
	Parallel bool
}

// ScenarioExecutor manages scenario execution
type ScenarioExecutor struct {
	executor *Executor
	mu       sync.RWMutex
}

// NewScenarioExecutor creates a new scenario executor
func NewScenarioExecutor(executor *Executor) *ScenarioExecutor {
	return &ScenarioExecutor{
		executor: executor,
	}
}

// Execute executes a scenario
func (se *ScenarioExecutor) Execute(ctx context.Context, scenario *Scenario) error {
	if scenario.Parallel {
		return se.executeParallel(ctx, scenario)
	}
	return se.executeSequential(ctx, scenario)
}

// executeSequential executes tasks sequentially
func (se *ScenarioExecutor) executeSequential(ctx context.Context, scenario *Scenario) error {
	for _, task := range scenario.Tasks {
		if err := se.executor.Execute(ctx, task); err != nil {
			return fmt.Errorf("task %s failed: %v", task.ID, err)
		}
	}
	return nil
}

// executeParallel executes tasks in parallel
func (se *ScenarioExecutor) executeParallel(ctx context.Context, scenario *Scenario) error {
	var wg sync.WaitGroup
	errChan := make(chan error, len(scenario.Tasks))

	for _, task := range scenario.Tasks {
		wg.Add(1)
		go func(t *Task) {
			defer wg.Done()
			if err := se.executor.Execute(ctx, t); err != nil {
				errChan <- err
			}
		}(task)
	}

	wg.Wait()
	close(errChan)

	for err := range errChan {
		if err != nil {
			return err
		}
	}

	return nil
}
