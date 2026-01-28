# Architecture

`toolcode` wraps tool discovery, docs, and execution into a single programmable
surface. The actual code execution is delegated to an injected `Engine`.

## Execution flow


![Diagram](assets/diagrams/execution-flow.svg)


## Snippet lifecycle


![Diagram](assets/diagrams/execution-flow.svg)


## Control points

- `Config` controls limits and defaults
- `Engine` controls the execution sandbox
- `Tools` surface mediates tool calls and captures traces
