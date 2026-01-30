# API Reference

## Executor

```go
type Executor interface {
  ExecuteCode(ctx context.Context, params ExecuteParams) (ExecuteResult, error)
}
```

### Executor contract

- Concurrency: implementations are safe for concurrent use.
- Context: honor cancellation/deadlines; deadline exceeded wraps `ErrLimitExceeded`.
- Errors: configuration failures return `ErrConfiguration`.
- Ownership: params are read-only; results are caller-owned snapshots.

## Engine

```go
type Engine interface {
  Execute(ctx context.Context, params ExecuteParams, tools Tools) (ExecuteResult, error)
}
```

### Engine contract

- Concurrency: implementations are safe for concurrent use.
- Context: honor cancellation/deadlines and return `ctx.Err()` when canceled.
- Errors: use `CodeError` where possible to include line/column metadata.
- Ownership: params/tools are read-only; results are caller-owned snapshots.

## Tools surface

```go
type Tools interface {
  SearchTools(ctx context.Context, query string, limit int) ([]toolindex.Summary, error)
  ListNamespaces(ctx context.Context) ([]string, error)
  DescribeTool(ctx context.Context, id string, level tooldocs.DetailLevel) (tooldocs.ToolDoc, error)
  ListToolExamples(ctx context.Context, id string, maxExamples int) ([]tooldocs.ToolExample, error)
  RunTool(ctx context.Context, id string, args map[string]any) (toolrun.RunResult, error)
  RunChain(ctx context.Context, steps []toolrun.ChainStep) (toolrun.RunResult, []toolrun.StepResult, error)
  Println(args ...any)
}
```

### Tools contract

- Concurrency: implementations are safe for concurrent use.
- Context: methods honor cancellation/deadlines and return `ctx.Err()` when canceled.
- Ownership: args are read-only; results are caller-owned snapshots.

## Params + results

```go
type ExecuteParams struct {
  Language     string
  Code         string
  Timeout      time.Duration
  MaxToolCalls int
}

type ExecuteResult struct {
  Value      any
  Stdout     string
  Stderr     string
  ToolCalls  []ToolCallRecord
  DurationMs int64
}
```

## ToolCallRecord

```go
type ToolCallRecord struct {
  ToolID     string
  Args       map[string]any
  Structured any
  BackendKind string
  Error      string
  ErrorOp    string
  DurationMs int64
}
```

### Errors

- `ErrConfiguration`
- `ErrLimitExceeded`
- `CodeError`

## Config

```go
type Config struct {
  Index          toolindex.Index
  Docs           tooldocs.Store
  Run            toolrun.Runner
  Engine         Engine
  DefaultTimeout time.Duration
  DefaultLanguage string
  MaxToolCalls   int
  MaxChainSteps  int
  Logger         Logger
}
```
