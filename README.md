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

You must provide an `Engine`. The executor handles defaults, limits, tool-call
tracing, and captured stdout.

```go
import (
  "context"
  "fmt"
  "log"
  "time"

  "github.com/jonwraymond/toolcode"
)

// A minimal engine that just runs one tool.
type oneToolEngine struct{}

func (oneToolEngine) Execute(ctx context.Context, params toolcode.ExecuteParams, tools toolcode.Tools) (toolcode.ExecuteResult, error) {
  res, err := tools.RunTool(ctx, "math:add", map[string]any{"a": 2, "b": 3})
  if err != nil {
    return toolcode.ExecuteResult{}, err
  }
  return toolcode.ExecuteResult{Value: res.Structured}, nil
}

exec, err := toolcode.NewDefaultExecutor(toolcode.Config{
  Index:  idx,
  Docs:   docs,
  Run:    runner,
  Engine: oneToolEngine{},

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

## Integration notes

- Tool IDs should be canonical (`namespace:name`)
- Chain semantics match `toolrun`:
  - `usePrevious` injects `args["previous"]`
  - it overwrites existing `previous`
  - it injects even when the previous result is nil

## Version compatibility (current tags)

- `toolmodel`: `v0.1.0`
- `toolindex`: `v0.1.1`
- `tooldocs`: `v0.1.1`
- `toolrun`: `v0.1.0`
- `toolcode`: `v0.1.0`
