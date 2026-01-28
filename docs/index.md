# toolcode

Code-mode orchestration for tool execution.

## What this repo provides

- Execute short orchestration snippets
- A minimal in-sandbox API
- Timeouts and limits

## Example

```go
res, _ := executor.ExecuteCode(ctx, "go", "__out = 2 + 2", 2*time.Second)
```
