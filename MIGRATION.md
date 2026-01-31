# Migration Guide: toolcode to toolexec/code

This document describes how to migrate from the deprecated `github.com/jonwraymond/toolcode`
package to its new location at `github.com/jonwraymond/toolexec/code`.

## Overview

The `toolcode` package has been consolidated into the `toolexec` repository as the
`code` subpackage. This consolidation simplifies the dependency graph and provides
a more cohesive API for tool execution and code orchestration.

## Import Path Changes

Update your import statements as follows:

| Old Import | New Import |
|------------|------------|
| `github.com/jonwraymond/toolcode` | `github.com/jonwraymond/toolexec/code` |

## Migration Steps

### 1. Update go.mod

Remove the old dependency and add the new one:

```bash
go get github.com/jonwraymond/toolexec/code
go mod tidy
```

This will remove `github.com/jonwraymond/toolcode` from your `go.mod` if it's no
longer referenced.

### 2. Update Import Statements

Replace all imports in your codebase:

**Before:**

```go
import (
  "github.com/jonwraymond/toolcode"
)
```

**After:**

```go
import (
  "github.com/jonwraymond/toolexec/code"
)
```

### 3. Update Type References

If you reference types with the package prefix, update them:

**Before:**

```go
toolcode.Config{}
toolcode.ExecuteParams{}
toolcode.ExecuteResult{}
toolcode.Tools
toolcode.NewDefaultExecutor()
```

**After:**

```go
code.Config{}
code.ExecuteParams{}
code.ExecuteResult{}
code.Tools
code.NewDefaultExecutor()
```

### 4. Automated Migration (Optional)

You can use `gofmt` with the `-r` flag or `gorename` for bulk updates:

```bash
# Using gofmt rewrite rule (for simple cases)
gofmt -w -r 'toolcode.Config -> code.Config' .

# Or use sed for import path replacement
find . -name "*.go" -exec sed -i '' \
  's|github.com/jonwraymond/toolcode|github.com/jonwraymond/toolexec/code|g' {} +
```

## API Compatibility

The API remains unchanged. All types, functions, and interfaces have been preserved
with the same signatures. The only change is the import path.

## Timeline

- **Now**: Both packages are available; `toolcode` shows deprecation warnings.
- **Future**: The `toolcode` repository will be archived after a transition period.

## Questions?

For questions or issues with the migration, please open an issue in the
[toolexec repository](https://github.com/jonwraymond/toolexec/issues).
