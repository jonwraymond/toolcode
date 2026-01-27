package toolcode

import (
	"fmt"
	"strings"
	"time"

	"github.com/jonwraymond/tooldocs"
	"github.com/jonwraymond/toolindex"
	"github.com/jonwraymond/toolrun"
)

// Config holds the configuration for a code executor.
type Config struct {
	// Index provides tool discovery and lookup capabilities.
	// Required.
	Index toolindex.Index

	// Docs provides tool documentation.
	// Required.
	Docs tooldocs.Store

	// Run provides tool execution capabilities.
	// Required.
	Run toolrun.Runner

	// Engine is the pluggable code execution engine.
	// Required.
	Engine Engine

	// DefaultTimeout is the default execution timeout when not specified
	// in ExecuteParams. If zero, no default timeout is applied.
	DefaultTimeout time.Duration

	// DefaultLanguage is the default language when not specified in
	// ExecuteParams. Defaults to "go" if empty.
	DefaultLanguage string

	// MaxToolCalls limits the maximum number of tool invocations per
	// execution. Zero means unlimited.
	MaxToolCalls int

	// MaxChainSteps limits the maximum number of steps allowed in a single
	// RunChain call. Zero means unlimited.
	MaxChainSteps int

	// Logger is an optional logger for observability.
	Logger Logger
}

// Validate checks that all required fields are set.
// Returns ErrConfiguration if any required field is missing.
func (c *Config) Validate() error {
	var missing []string

	if c.Index == nil {
		missing = append(missing, "Index")
	}
	if c.Docs == nil {
		missing = append(missing, "Docs")
	}
	if c.Run == nil {
		missing = append(missing, "Run")
	}
	if c.Engine == nil {
		missing = append(missing, "Engine")
	}

	if len(missing) > 0 {
		return fmt.Errorf("%w: missing required fields: %s",
			ErrConfiguration, strings.Join(missing, ", "))
	}
	return nil
}

// applyDefaults sets default values for optional fields.
func (c *Config) applyDefaults() {
	if c.DefaultLanguage == "" {
		c.DefaultLanguage = "go"
	}
}
