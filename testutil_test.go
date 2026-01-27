package toolcode

import (
	"context"
	"sync"

	"github.com/jonwraymond/tooldocs"
	"github.com/jonwraymond/toolindex"
	"github.com/jonwraymond/toolmodel"
	"github.com/jonwraymond/toolrun"
)

// mockIndex implements toolindex.Index for testing.
type mockIndex struct {
	mu sync.Mutex

	// Configurable returns
	searchResult     []toolindex.Summary
	searchErr        error
	namespacesResult []string
	getToolResult    toolmodel.Tool
	getToolBackend   toolmodel.ToolBackend
	getToolErr       error

	// Call tracking
	searchCalls     []searchCall
	getToolCalls    []string
	namespacesCalls int
}

type searchCall struct {
	query string
	limit int
}

func (m *mockIndex) Search(query string, limit int) ([]toolindex.Summary, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.searchCalls = append(m.searchCalls, searchCall{query, limit})
	return m.searchResult, m.searchErr
}

func (m *mockIndex) ListNamespaces() ([]string, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.namespacesCalls++
	return m.namespacesResult, nil
}

func (m *mockIndex) GetTool(id string) (toolmodel.Tool, toolmodel.ToolBackend, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.getToolCalls = append(m.getToolCalls, id)
	return m.getToolResult, m.getToolBackend, m.getToolErr
}

func (m *mockIndex) GetAllBackends(_ string) ([]toolmodel.ToolBackend, error) {
	return nil, nil
}

func (m *mockIndex) RegisterTool(_ toolmodel.Tool, _ toolmodel.ToolBackend) error {
	return nil
}

func (m *mockIndex) RegisterTools(_ []toolindex.ToolRegistration) error {
	return nil
}

func (m *mockIndex) RegisterToolsFromMCP(_ string, _ []toolmodel.Tool) error {
	return nil
}

func (m *mockIndex) UnregisterBackend(_ string, _ toolmodel.BackendKind, _ string) error {
	return nil
}

// mockStore implements tooldocs.Store for testing.
type mockStore struct {
	mu sync.Mutex

	// Configurable returns
	describeResult tooldocs.ToolDoc
	describeErr    error
	examplesResult []tooldocs.ToolExample
	examplesErr    error

	// Call tracking
	describeCalls []describeCall
	examplesCalls []examplesCall
}

type describeCall struct {
	id    string
	level tooldocs.DetailLevel
}

type examplesCall struct {
	id          string
	maxExamples int
}

func (m *mockStore) DescribeTool(id string, level tooldocs.DetailLevel) (tooldocs.ToolDoc, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.describeCalls = append(m.describeCalls, describeCall{id, level})
	return m.describeResult, m.describeErr
}

func (m *mockStore) ListExamples(id string, maxExamples int) ([]tooldocs.ToolExample, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.examplesCalls = append(m.examplesCalls, examplesCall{id, maxExamples})
	return m.examplesResult, m.examplesErr
}

// mockRunner implements toolrun.Runner for testing.
type mockRunner struct {
	mu sync.Mutex

	// Configurable returns
	runResult    toolrun.RunResult
	runErr       error
	chainResult  toolrun.RunResult
	chainSteps   []toolrun.StepResult
	chainErr     error
	streamEvents []toolrun.StreamEvent
	streamErr    error

	// Call tracking
	runCalls   []runCall
	chainCalls [][]toolrun.ChainStep
}

type runCall struct {
	ctx    context.Context
	toolID string
	args   map[string]any
}

func (m *mockRunner) Run(ctx context.Context, toolID string, args map[string]any) (toolrun.RunResult, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.runCalls = append(m.runCalls, runCall{ctx, toolID, args})
	return m.runResult, m.runErr
}

func (m *mockRunner) RunStream(_ context.Context, _ string, _ map[string]any) (<-chan toolrun.StreamEvent, error) {
	ch := make(chan toolrun.StreamEvent, len(m.streamEvents)+1)
	for _, e := range m.streamEvents {
		ch <- e
	}
	close(ch)
	return ch, m.streamErr
}

func (m *mockRunner) RunChain(ctx context.Context, steps []toolrun.ChainStep) (toolrun.RunResult, []toolrun.StepResult, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.chainCalls = append(m.chainCalls, steps)
	return m.chainResult, m.chainSteps, m.chainErr
}

// mockEngine implements Engine for testing.
type mockEngine struct {
	mu sync.Mutex

	// Configurable returns
	executeResult ExecuteResult
	executeErr    error

	// Call tracking
	executeCalls []executeCall
}

type executeCall struct {
	ctx    context.Context
	params ExecuteParams
	tools  Tools
}

func (m *mockEngine) Execute(ctx context.Context, params ExecuteParams, tools Tools) (ExecuteResult, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.executeCalls = append(m.executeCalls, executeCall{ctx, params, tools})
	return m.executeResult, m.executeErr
}

// mockLogger implements Logger for testing.
type mockLogger struct {
	mu       sync.Mutex
	messages []string
}

func (l *mockLogger) Logf(_ string, _ ...any) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.messages = append(l.messages, "")
}
