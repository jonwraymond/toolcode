# Design Notes

This page explains the tradeoffs and error semantics behind `toolcode`.

## Design tradeoffs

- **Engine abstraction.** `toolcode` does not execute code directly. It delegates to a pluggable `Engine` so you can swap in different interpreters or sandboxes without changing the API.
- **Narrow tools surface.** Code snippets only see a constrained `Tools` interface (SearchTools, DescribeTool, RunTool, RunChain, Println). This limits what code can do and keeps behavior consistent.
- **Limits-first design.** Max tool calls and chain steps are enforced in the tools environment. Timeouts are enforced at the executor level.
- **Structured observability.** Every tool call is recorded with duration, backend kind, and error op for auditability.
- **Default language.** If not specified, `DefaultLanguage` is used ("go" by default), but the engine decides how that language is implemented.

## Error semantics

`toolcode` uses sentinel errors for predictable handling:

- `ErrCodeExecution` – syntax/runtime errors in the snippet.
- `ErrConfiguration` – missing required config (Index, Docs, Run, Engine).
- `ErrLimitExceeded` – timeouts, max tool calls, or max chain steps exceeded.

`CodeError` wraps engine errors with optional line/column info.

## Extension points

- **Custom engines:** implement `Engine` for a sandboxed interpreter, WASM runtime, or remote execution service.
- **Custom logging:** provide a `Logger` to track tool calls and execution duration.
- **Runtime integration:** use `toolruntime/toolcodeengine` to run `toolcode` on top of secure runtime backends.

## Operational guidance

- Keep snippets short and deterministic; `ExecuteParams` is intended for orchestration, not full applications.
- Use `MaxToolCalls` to prevent runaway tool usage in long snippets.
- Treat `Stdout` as diagnostic output only; use `Value` for structured results.
