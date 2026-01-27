# toolcode

`toolcode` is the code-orchestration layer on top of:
- `github.com/jonwraymond/toolindex`
- `github.com/jonwraymond/tooldocs`
- `github.com/jonwraymond/toolrun`

It executes short, constrained snippets through an injected `Engine` while
exposing a small metatool surface (`SearchTools`, `DescribeTool`, `RunTool`,
`RunChain`, `Println`).

## Install

```bash
go get github.com/jonwraymond/toolcode
```

## Quick start

You must provide an `Engine`. The executor handles defaults, limits, and
tool-call tracing.

```go
exec, err := toolcode.NewDefaultExecutor(toolcode.Config{
  Index:  idx,
  Docs:   docs,
  Run:    runner,
  Engine: engine,

  DefaultLanguage: "go",
  DefaultTimeout:  5 * time.Second,
  MaxToolCalls:    20,
  MaxChainSteps:   10,
})
if err != nil {
  log.Fatal(err)
}

res, err := exec.ExecuteCode(ctx, toolcode.ExecuteParams{
  Language: "go",
  Code:     `__out = "ok"`,
})
if err != nil {
  log.Fatal(err)
}

fmt.Println(res.Value, len(res.ToolCalls))
```

## Limits and tracing

- Timeouts are enforced via context deadlines.
- Tool calls and chain steps count against `MaxToolCalls`.
- Chain length can be capped via `MaxChainSteps`.
- All tool calls are recorded in `ExecuteResult.ToolCalls` with `durationMs`.

