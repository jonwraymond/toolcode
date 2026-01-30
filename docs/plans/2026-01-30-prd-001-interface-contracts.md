# PRD-001 Interface Contracts — toolcode

**Status:** Done
**Date:** 2026-01-30


## Overview
Define explicit interface contracts (GoDoc + documented semantics) for all interfaces in this repo. Contracts must state concurrency guarantees, error semantics, ownership of inputs/outputs, and context handling.


## Goals
- Every interface has explicit GoDoc describing behavioral contract.
- Contract behavior is codified in tests (contract tests).
- Docs/README updated where behavior is user-facing.


## Non-Goals
- No API shape changes unless required to satisfy the contract tests.
- No new features beyond contract clarity and tests.


## Interface Inventory
| Interface | File | Methods |
| --- | --- | --- |
| `Logger` | `toolcode/logger.go:5` | Logf(format string, args ...any) |
| `Tools` | `toolcode/tools.go:19` | SearchTools(ctx context.Context, query string, limit int) ([]toolindex.Summary, error)<br/>ListNamespaces(ctx context.Context) ([]string, error)<br/>DescribeTool(ctx context.Context, id string, level tooldocs.DetailLevel) (tooldocs.ToolDoc, error)<br/>ListToolExamples(ctx context.Context, id string, maxExamples int) ([]tooldocs.ToolExample, error)<br/>RunTool(ctx context.Context, id string, args map[string]any) (toolrun.RunResult, error)<br/>RunChain(ctx context.Context, steps []toolrun.ChainStep) (toolrun.RunResult, []toolrun.StepResult, error)<br/>Println(args ...any) |
| `Engine` | `toolcode/engine.go:14` | Execute(ctx context.Context, params ExecuteParams, tools Tools) (ExecuteResult, error) |
| `Executor` | `toolcode/executor.go:12` | ExecuteCode(ctx context.Context, params ExecuteParams) (ExecuteResult, error) |

## Contract Template (apply per interface)
- **Thread-safety:** explicitly state if safe for concurrent use.
- **Context:** cancellation/deadline handling (if context is a parameter).
- **Errors:** classification, retryability, and wrapping expectations.
- **Ownership:** who owns/allocates inputs/outputs; mutation expectations.
- **Determinism/order:** ordering guarantees for returned slices/maps/streams.
- **Nil/zero handling:** behavior for nil inputs or empty values.


## Acceptance Criteria
- All interfaces have GoDoc with explicit behavioral contract.
- Contract tests exist and pass.
- No interface contract contradictions across repos.
