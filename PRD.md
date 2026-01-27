# PRD: toolcode Library

## 1) Purpose and scope
toolcode is the code-mode orchestration layer on top of:
- `toolindex` (discovery + lookup)
- `tooldocs` (schema + examples)
- `toolrun` (execution + chains)

It executes short, constrained code snippets that call metatool-style helper
functions (search, describe, run) while keeping a single, MCP-aligned surface.

Scope: code execution orchestration only. No discovery, docs storage, or tool
execution logic is reimplemented here.

Target MCP protocol version: **2025-11-25** (via `toolmodel.MCPVersion`).

## 2) Goals and non-goals
### Goals
- Provide a safe, small execution surface for agentic orchestration via code
- Reuse the existing libraries rather than duplicating behavior
- Record tool calls for observability and debugging
- Enforce timeouts and bounded execution
- Keep the code engine pluggable
- Align cleanly with MCP metatools, especially `execute_code`

### Non-goals
- No general-purpose sandbox for arbitrary workloads
- No network transports implemented here
- No replacement for `toolindex`, `tooldocs`, or `toolrun`
- No MCP server wiring in this module (metatools server handles that)

## 3) Dependencies and alignment
- `github.com/jonwraymond/toolindex`
  - Discovery results via `[]toolindex.Summary`
- `github.com/jonwraymond/tooldocs`
  - Detail retrieval via `tooldocs.Store`
- `github.com/jonwraymond/toolrun`
  - Execution via `toolrun.Runner`
- `github.com/jonwraymond/toolmodel`
  - Canonical tool schema and MCP version targeting

toolcode must mirror the existing contracts:
- Tool IDs are canonical `Tool.ToolID()` strings (`namespace:name`)
- Chain semantics match `toolrun` (not redefined here)
- Detail levels match `tooldocs.DetailLevel`

## 4) Users and primary use cases
### Users
- Metatools MCP server implementing `execute_code`
- Agents that want to orchestrate via a short snippet instead of many tool calls
- Test harnesses for orchestration strategies

### Use cases
1) Discover -> describe -> run in a single snippet
2) Small multi-step workflows with explicit control flow
3) Debuggability: inspect tool call traces and intermediate values

## 5) Functional requirements
### 5.1 Metatool environment
toolcode must provide a metatool environment that exposes a minimal, stable
set of helper functions to the executing code:
- `SearchTools(query string, limit int) ([]toolindex.Summary, error)`
- `ListNamespaces() ([]string, error)`
- `DescribeTool(id string, level tooldocs.DetailLevel) (tooldocs.ToolDoc, error)`
- `ListToolExamples(id string, max int) ([]tooldocs.ToolExample, error)`
- `RunTool(ctx context.Context, id string, args map[string]any) (toolrun.RunResult, error)`
- `RunChain(ctx context.Context, steps []toolrun.ChainStep) (toolrun.RunResult, []toolrun.StepResult, error)`
- `Println(args ...any)` (captured to `stdout`)

These are thin wrappers over the injected dependencies.

Chain semantics must not be redefined here. In particular, `RunChain` uses
toolrun's rule: when `UsePrevious` is true, the prior step's structured result
is injected at `args["previous"]`.

The snippet must not have direct access to the network, filesystem, or
environment beyond what is explicitly exposed by the engine.

### 5.2 Supported languages (v1)
Start with a single Go-like scripting mode (for example, `"go"` or
`"goscript"`). The `language` field remains part of the API to allow additional
engines later.

### 5.3 Execution limits and safety
toolcode must support execution limits:
- time limit / timeout
- maximum tool calls
- optional maximum chain steps per call

When limits are exceeded, execution halts with a clear error.

### 5.4 Tool call tracing
toolcode must record tool calls made during execution, including:
- tool ID
- arguments (normalized to MCP-native shapes)
- structured result (when available)
- error (when present)
- duration

This trace is part of the returned result for observability.

### 5.5 Final result convention
Define a single clear convention for the snippet's return value in v1:
- The engine returns the value of a special variable named `__out`

Agent guidance: always set `__out` to the final value.

### 5.6 Error handling
Differentiate:
- snippet syntax/runtime errors
- tool errors surfaced via `RunTool` / `RunChain`

For snippet errors:
- return a non-nil error
- populate `stderr` where possible

Tool errors should be returned as errors from helpers. Engines may allow
language-specific catching, but the default behavior is fail-fast.

### 5.7 Pluggable code engine
toolcode must not hard-code an interpreter. Instead:
- define a small engine interface
- allow a default engine to be injected

The engine receives the metatool environment and is responsible for executing
the snippet safely within the provided constraints.

### 5.8 Observability hooks
Provide optional hooks for logging and metrics around:
- tool calls (ID, latency, success/failure)
- snippet execution duration and limit enforcement

### 5.9 MCP-aligned result shaping
The execution result must be easy to expose via an MCP `execute_code` tool:
- primary value (`value`)
- `stdout` and `stderr`
- tool call trace
- duration and limit metadata (when useful)

## 6) Public Go API (v1)
The API should be explicit and align with the other modules' public types.

```go
package toolcode

import (
  "context"
  "time"

  "github.com/jonwraymond/toolindex"
  "github.com/jonwraymond/tooldocs"
  "github.com/jonwraymond/toolrun"
)

// ToolCallRecord captures one tool call made during code execution.
type ToolCallRecord struct {
  ToolID     string         `json:"toolId"`
  Args       map[string]any `json:"args,omitempty"`
  Structured any            `json:"structured,omitempty"`
  Err        error          `json:"-"`
  Duration   time.Duration  `json:"duration"`
}

// ExecuteParams describes a code execution request.
type ExecuteParams struct {
  Language     string        `json:"language"`
  Code         string        `json:"code"`
  Timeout      time.Duration `json:"timeout"`
  MaxToolCalls int           `json:"maxToolCalls,omitempty"`
}

// ExecuteResult is the normalized output of code execution.
type ExecuteResult struct {
  Value     any              `json:"value,omitempty"`
  Stdout    string           `json:"stdout,omitempty"`
  Stderr    string           `json:"stderr,omitempty"`
  ToolCalls []ToolCallRecord `json:"toolCalls,omitempty"`
  Duration  time.Duration    `json:"duration"`
}

// Tools defines the metatool environment available to executing code.
type Tools interface {
  SearchTools(query string, limit int) ([]toolindex.Summary, error)
  ListNamespaces() ([]string, error)
  DescribeTool(id string, level tooldocs.DetailLevel) (tooldocs.ToolDoc, error)
  ListToolExamples(id string, max int) ([]tooldocs.ToolExample, error)
  RunTool(ctx context.Context, id string, args map[string]any) (toolrun.RunResult, error)
  RunChain(ctx context.Context, steps []toolrun.ChainStep) (toolrun.RunResult, []toolrun.StepResult, error)
  Println(args ...any)
}

// Engine executes code with access to the metatool environment.
type Engine interface {
  Execute(ctx context.Context, params ExecuteParams, tools Tools) (ExecuteResult, error)
}

// Executor is the main entrypoint for callers.
type Executor interface {
  ExecuteCode(ctx context.Context, params ExecuteParams) (ExecuteResult, error)
}

// Logger is an optional observability hook.
type Logger interface {
  Logf(format string, args ...any)
}

// Config wires toolcode to the other modules.
type Config struct {
  Index toolindex.Index
  Docs  tooldocs.Store
  Run   toolrun.Runner

  Engine Engine

  DefaultTimeout  time.Duration
  DefaultLanguage string
  MaxToolCalls    int
  Logger          Logger
}

type DefaultExecutor struct {
  cfg Config
}
```

### Configuration rules
- `Index`, `Docs`, and `Run` are required
- If `Engine` is nil, construction should fail fast
- If `params.Timeout` is zero, use `DefaultTimeout`
- If `params.Language` is empty, use `DefaultLanguage`
- `MaxToolCalls` in params is capped by config `MaxToolCalls` when set
- `Logger` is optional

## 7) MCP-facing schemas (for the metatools server)
toolcode itself is transport-agnostic, but it must map cleanly to an MCP tool.

### execute_code
Input:
- `language: string` (v1: `"go"` only)
- `code: string`
- `timeout_ms: integer` (optional)
- `max_tool_calls: integer` (optional)

Output:
- `value: any`
- `stdout: string`
- `stderr: string`
- `tool_calls: [ { toolId, args, structured, duration_ms, error? } ]`
- `duration_ms: integer`

## 8) Non-functional requirements
- Deterministic, testable orchestration behavior
- Clear limits and failure modes
- No redefinition of discovery/docs/execution semantics
- Thin wrapper cost relative to underlying tool calls
- No direct network/filesystem/environment access beyond exposed helpers

## 9) Example flow
1) Executor receives `ExecuteParams`
2) It constructs a metatool environment backed by `Index`, `Docs`, and `Run`
3) It delegates execution to `Engine`
4) It returns the normalized result plus tool call trace

## 10) Success criteria and gates
### Gates
- `go test ./...`

### Success criteria
- Code snippets can discover, describe, and execute tools via one surface
- Tool calls are traced and surfaced in results
- The result shape cleanly powers an MCP `execute_code` metatool
