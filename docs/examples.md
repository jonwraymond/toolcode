# Examples

## Minimal snippet

```go
res, _ := exec.ExecuteCode(ctx, toolcode.ExecuteParams{
  Language: "go",
  Code:     "__out = 2 + 2",
})
```

## Tool selection + execution

```go
code := `
  tools := SearchTools("get repo", 3)
  if len(tools) == 0 {
    __out = "no tools"
  } else {
    __out = RunTool(tools[0].ID, map[string]any{"owner":"octo","repo":"hello"})
  }
`

res, _ := exec.ExecuteCode(ctx, toolcode.ExecuteParams{Language: "go", Code: code})
```

## Chain helper

```go
steps := []ChainStep{
  {ToolID: "user:get", Args: map[string]any{"user_id": "123"}},
  {ToolID: "orders:list", UsePrevious: true},
}

code := `__out = RunChain(steps)`
```
