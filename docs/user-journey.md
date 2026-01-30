# User Journey

This journey shows how `toolcode` enables end-to-end orchestration using code snippets.

## End-to-end flow (stack view)

![Diagram](assets/diagrams/user-journey.svg)

```mermaid
%%{init: {'theme': 'base', 'themeVariables': {'primaryColor': '#6b46c1', 'primaryTextColor': '#fff'}}}%%
flowchart TB
    subgraph input["Input"]
        Code["💻 Code Snippet<br/><small>SearchTools, RunTool, etc.</small>"]
    end

    subgraph executor["Executor"]
        Exec["⚙️ toolcode.Executor"]
        Limits["🛡️ Limits<br/><small>timeout, maxToolCalls</small>"]
    end

    subgraph engine["Engine"]
        Eng["🔧 Engine.Run()"]
        Tools["🔨 Tools Environment"]
    end

    subgraph metatools["Metatool Surface"]
        Search["🔍 SearchTools()"]
        Describe["📚 DescribeTool()"]
        RunTool["▶️ RunTool()"]
        RunChain["🔗 RunChain()"]
        Println["🖨️ Println()"]
    end

    subgraph output["Output"]
        Result["📦 ExecuteResult<br/><small>Value + Stdout + ToolCalls[]</small>"]
    end

    Code --> Exec --> Limits --> Eng --> Tools
    Tools --> Search
    Tools --> Describe
    Tools --> RunTool
    Tools --> RunChain
    Tools --> Println
    Search --> Result
    RunTool --> Result
    RunChain --> Result

    style input fill:#3182ce,stroke:#2c5282
    style executor fill:#6b46c1,stroke:#553c9a,stroke-width:2px
    style engine fill:#d69e2e,stroke:#b7791f
    style metatools fill:#38a169,stroke:#276749
    style output fill:#3182ce,stroke:#2c5282
```

### Tool Call Recording

```mermaid
%%{init: {'theme': 'base', 'themeVariables': {'primaryColor': '#e53e3e'}}}%%
flowchart LR
    subgraph snippet["Code Snippet"]
        Call1["RunTool('a', args)"]
        Call2["RunTool('b', args)"]
        Call3["RunChain([...])"]
    end

    subgraph recording["Recording"]
        Rec["📝 ToolCallRecord[]<br/><small>ToolID, Args, Result,<br/>Backend, Duration</small>"]
    end

    subgraph result["ExecuteResult"]
        Value["Value: __out"]
        Stdout["Stdout: captured"]
        Calls["ToolCalls: [...]"]
    end

    Call1 --> Rec
    Call2 --> Rec
    Call3 --> Rec
    Rec --> Calls
    snippet --> Value
    snippet --> Stdout

    style snippet fill:#6b46c1,stroke:#553c9a
    style recording fill:#e53e3e,stroke:#c53030
    style result fill:#38a169,stroke:#276749
```

## Step-by-step

1. **Agent submits a snippet** to `execute_code`.
2. **Executor sets limits** (timeout, tool call caps).
3. **Engine runs** the snippet with a `Tools` environment.
4. **Snippet orchestrates** tools via Search/Describe/Run APIs.
5. **Result is returned** as `ExecuteResult` (Value + ToolCalls + Stdout).

## Example: pick a tool and run it

```go
code := `
results, _ := SearchTools("create issue", 5)
choice := results[0]

schema, _ := DescribeTool(choice.ID, "schema")
_ = schema // use schema in real code

res, _ := RunTool(choice.ID, map[string]any{"title": "Bug", "body": "..."})
__out = res.Structured
`

exec, _ := toolcode.NewDefaultExecutor(cfg)
result, err := exec.ExecuteCode(ctx, toolcode.ExecuteParams{Code: code})
```

## Expected outcomes

- A deterministic orchestration surface with strong limits.
- Tool call traces for audit/debugging.
- Easy integration with secure runtimes via `toolruntime` adapters.

## Common failure modes

- `ErrConfiguration` if required dependencies are missing.
- `ErrLimitExceeded` if timeouts or tool call caps are hit.
- `ErrCodeExecution` if the engine rejects the snippet.
