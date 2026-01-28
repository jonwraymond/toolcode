# Usage

## Configure an executor

```go
exec, err := toolcode.NewDefaultExecutor(toolcode.Config{
  Index:          idx,
  Docs:           docs,
  Run:            runner,
  Engine:         engine,
  DefaultTimeout: 5 * time.Second,
  MaxToolCalls:   32,
  MaxChainSteps:  8,
})
if err != nil {
  // handle config error
}
```

## Execute a snippet

```go
res, err := exec.ExecuteCode(ctx, toolcode.ExecuteParams{
  Language: "go",
  Code: `
    tools := SearchTools("repo", 3)
    repo := RunTool(tools[0].ID, map[string]any{"owner":"o","repo":"r"})
    __out = repo
  `,
})
```

## Limits

- `MaxToolCalls` caps tool invocations per execution
- `MaxChainSteps` caps chain length
- `DefaultTimeout` applies when params omit a timeout
