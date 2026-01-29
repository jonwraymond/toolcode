# toolcode

`toolcode` executes short orchestration snippets over the tool stack. It exposes
Search/Describe/Run helpers to code, enforces timeouts and limits, and records
tool calls for observability.

[![Docs](https://img.shields.io/badge/docs-ai--tools--stack-blue)](https://jonwraymond.github.io/ai-tools-stack/)

## Deep dives
- Design Notes: [design-notes.md](design-notes.md)
- User Journey: [user-journey.md](user-journey.md)

## Motivation

- **Programmable orchestration**: loops, conditionals, fallbacks
- **Uniform tool surface**: same helpers regardless of backend
- **Traceability**: tool call records for debugging and audits

## Key APIs

- `Executor` interface (`ExecuteCode`)
- `DefaultExecutor` implementation
- `Engine` interface (pluggable runtime)
- `ExecuteParams`, `ExecuteResult`, `ToolCallRecord`

## Quickstart

```go
exec, _ := toolcode.NewDefaultExecutor(toolcode.Config{
  Index:  idx,
  Docs:   docs,
  Run:    runner,
  Engine: engine,
})

res, _ := exec.ExecuteCode(ctx, toolcode.ExecuteParams{
  Language: "go",
  Code:     "__out = 2 + 2",
})
```

## Usability notes

- `ExecuteParams` supports per-call timeouts
- Tool calls are capped via config and params
- `ExecuteResult` includes tool call traces

## Next

- Execution pipeline: `architecture.md`
- Config, limits, and params: `usage.md`
- Examples: `examples.md`
- Design Notes: [design-notes.md](design-notes.md)
- User Journey: [user-journey.md](user-journey.md)

