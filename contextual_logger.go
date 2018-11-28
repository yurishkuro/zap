// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package zap

import (
	"context"
)

// A ContextualLogger is a decorator around the main Logger that allows
// tagging of the log entries with metadata passed to log methods via
// context.Context.
//
// ContextualLogger requires EncoderConfig.EncodeContext handler to be set.
type ContextualLogger struct {
	base *Logger
}

// Named adds a new path segment to the logger's name. Segments are joined by
// periods. By default, Loggers are unnamed.
func (log *ContextualLogger) Named(s string) *ContextualLogger {
	return &ContextualLogger{log.base.Named(s)}
}

// With creates a child logger and adds structured context to it. Fields added
// to the child don't affect the parent, and vice versa.
func (log *ContextualLogger) With(fields ...Field) *ContextualLogger {
	return &ContextualLogger{log.base.With(fields...)}
}

// Debug logs a message at DebugLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (log *ContextualLogger) Debug(ctx context.Context, msg string, fields ...Field) {
	if ce := log.base.check(DebugLevel, msg); ce != nil {
		ce.Context = ctx
		ce.Write(fields...)
	}
}

// Info logs a message at InfoLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (log *ContextualLogger) Info(ctx context.Context, msg string, fields ...Field) {
	if ce := log.base.check(InfoLevel, msg); ce != nil {
		ce.Context = ctx
		ce.Write(fields...)
	}
}

// Warn logs a message at WarnLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (log *ContextualLogger) Warn(ctx context.Context, msg string, fields ...Field) {
	if ce := log.base.check(WarnLevel, msg); ce != nil {
		ce.Context = ctx
		ce.Write(fields...)
	}
}

// Error logs a message at ErrorLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (log *ContextualLogger) Error(ctx context.Context, msg string, fields ...Field) {
	if ce := log.base.check(ErrorLevel, msg); ce != nil {
		ce.Context = ctx
		ce.Write(fields...)
	}
}

// DPanic logs a message at DPanicLevel. The message includes any fields
// passed at the log site, as well as any fields accumulated on the logger.
//
// If the logger is in development mode, it then panics (DPanic means
// "development panic"). This is useful for catching errors that are
// recoverable, but shouldn't ever happen.
func (log *ContextualLogger) DPanic(ctx context.Context, msg string, fields ...Field) {
	if ce := log.base.check(DPanicLevel, msg); ce != nil {
		ce.Context = ctx
		ce.Write(fields...)
	}
}

// Panic logs a message at PanicLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// The logger then panics, even if logging at PanicLevel is disabled.
func (log *ContextualLogger) Panic(ctx context.Context, msg string, fields ...Field) {
	if ce := log.base.check(PanicLevel, msg); ce != nil {
		ce.Context = ctx
		ce.Write(fields...)
	}
}

// Fatal logs a message at FatalLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// The logger then calls os.Exit(1), even if logging at FatalLevel is
// disabled.
func (log *ContextualLogger) Fatal(ctx context.Context, msg string, fields ...Field) {
	if ce := log.base.check(FatalLevel, msg); ce != nil {
		ce.Context = ctx
		ce.Write(fields...)
	}
}

// Base returns the underlying Logger.
func (log *ContextualLogger) Base() *Logger {
	return log.base
}
