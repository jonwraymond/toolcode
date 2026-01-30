# PRD-001 Execution Plan — toolcode (TDD)

**Status:** Ready
**Date:** 2026-01-30
**PRD:** `2026-01-30-prd-001-interface-contracts.md`


## TDD Workflow (required)
1. Red — write failing contract tests
2. Red verification — run tests
3. Green — minimal code/doc changes
4. Green verification — run tests
5. Commit — one commit per task


## Tasks
### Task 0 — Inventory + contract outline
- Confirm interface list and method signatures.
- Draft explicit contract bullets for each interface.
- Update docs/plans/README.md with this PRD + plan.
### Task 1 — Contract tests (Red/Green)
- Add `*_contract_test.go` with tests for each interface listed below.
- Use stub implementations where needed.
### Task 2 — GoDoc contracts
- Add/expand GoDoc on each interface with explicit contract clauses (thread-safety, errors, context, ownership).
- Update README/design-notes if user-facing.
### Task 3 — Verification
- Run `go test ./...`
- Run linters if configured (golangci-lint / gosec).


## Test Skeletons (contract_test.go)
### Logger
```go
func TestLogger_Contract(t *testing.T) {
    // Methods:
    // - Logf(format string, args ...any)
    // Contract assertions:
    // - Concurrency guarantees documented and enforced
    // - Error semantics (types/wrapping) validated
    // - Context cancellation respected (if applicable)
    // - Deterministic ordering where required
    // - Nil/zero input handling specified
}
```
### Tools
```go
func TestTools_Contract(t *testing.T) {
    // Methods:
    // - SearchTools(ctx context.Context, query string, limit int) ([]toolindex.Summary, error)
    // - ListNamespaces(ctx context.Context) ([]string, error)
    // - DescribeTool(ctx context.Context, id string, level tooldocs.DetailLevel) (tooldocs.ToolDoc, error)
    // - ListToolExamples(ctx context.Context, id string, maxExamples int) ([]tooldocs.ToolExample, error)
    // - RunTool(ctx context.Context, id string, args map[string]any) (toolrun.RunResult, error)
    // - RunChain(ctx context.Context, steps []toolrun.ChainStep) (toolrun.RunResult, []toolrun.StepResult, error)
    // - Println(args ...any)
    // Contract assertions:
    // - Concurrency guarantees documented and enforced
    // - Error semantics (types/wrapping) validated
    // - Context cancellation respected (if applicable)
    // - Deterministic ordering where required
    // - Nil/zero input handling specified
}
```
### Engine
```go
func TestEngine_Contract(t *testing.T) {
    // Methods:
    // - Execute(ctx context.Context, params ExecuteParams, tools Tools) (ExecuteResult, error)
    // Contract assertions:
    // - Concurrency guarantees documented and enforced
    // - Error semantics (types/wrapping) validated
    // - Context cancellation respected (if applicable)
    // - Deterministic ordering where required
    // - Nil/zero input handling specified
}
```
### Executor
```go
func TestExecutor_Contract(t *testing.T) {
    // Methods:
    // - ExecuteCode(ctx context.Context, params ExecuteParams) (ExecuteResult, error)
    // Contract assertions:
    // - Concurrency guarantees documented and enforced
    // - Error semantics (types/wrapping) validated
    // - Context cancellation respected (if applicable)
    // - Deterministic ordering where required
    // - Nil/zero input handling specified
}
```
