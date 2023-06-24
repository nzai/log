package log

import (
	"context"
	"io"
)

type Parameter struct {
	Encoder             Encoder
	Writer              io.Writer
	LogLevel            LogLevel
	StaticFields        []Field
	DynamicFields       func(context.Context) []Field
	DynamicKeyAndValues func(context.Context) []interface{}
}

type Encoder string

const (
	JSON    Encoder = "json"
	Console Encoder = "console"
)

func (e Encoder) String() string {
	return string(e)
}

// Option logger option
type Option func(*Parameter)

func WithEncoder(e Encoder) Option {
	return func(c *Parameter) {
		c.Encoder = e
	}
}

func WithWriter(w io.Writer) Option {
	return func(c *Parameter) {
		c.Writer = w
	}
}

func WithLogLevel(level LogLevel) Option {
	return func(c *Parameter) {
		c.LogLevel = level
	}
}

func WithStaticFields(fields []Field) Option {
	return func(c *Parameter) {
		c.StaticFields = fields
	}
}

func WithDynamicFields(fn func(context.Context) []Field) Option {
	return func(c *Parameter) {
		c.DynamicFields = fn
	}
}

func WithDynamicKeyAndValues(fn func(context.Context) []interface{}) Option {
	return func(c *Parameter) {
		c.DynamicKeyAndValues = fn
	}
}
