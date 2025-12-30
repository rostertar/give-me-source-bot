package log

import (
	"context"
	"fmt"
	"time"
)

type Logger interface {
	Debugf(msg string, args ...any)
	Infof(msg string, args ...any)
	Warnf(msg string, args ...any)
	Errorf(msg string, args ...any)

	WithField(key string, value any) Logger

	WithLevel(level Level) Logger

	WithOutput(fo FormatedOutput) Logger
}

type logger struct {
	context.Context

	Fields *Field

	Formater FormatedOutput

	debugf func(string, ...any)
	infof  func(string, ...any)
	warnf  func(string, ...any)
	errorf func(string, ...any)
}

func NewDefaultLogger(ctx context.Context, formater FormatedOutput, level Level) Logger {
	lgr := &logger{
		Context:  ctx,
		Formater: formater,
	}
	return lgr.WithLevel(level)
}

func (l *logger) copy() *logger {
	ln := new(logger)
	*ln = *l
	return ln
}

func (l *logger) WithField(key string, value any) Logger {
	ln := l.copy()
	ln.Fields = NewFiled(key, value).With(l.Fields)
	return ln
}

func nop(string, ...any) {}

func (l *logger) WithLevel(level Level) Logger {
	ln := l.copy()
	ln.debugf = nop
	ln.infof = nop
	ln.warnf = nop
	ln.errorf = nop
	switch level {
	case LevelDebug:
		ln.debugf = func(msg string, args ...any) {
			ln.somef(LevelDebug, msg, args)
		}
		fallthrough
	case LevelInfo:
		ln.infof = func(msg string, args ...any) {
			ln.somef(LevelInfo, msg, args)
		}
		fallthrough
	case LevelWarn:
		ln.warnf = func(msg string, args ...any) {
			ln.somef(LevelWarn, msg, args)
		}
		fallthrough
	case LevelError:
		ln.errorf = func(msg string, args ...any) {
			ln.somef(LevelError, msg, args)
		}
	case LevelDisable:
		// nothing to do
	}
	return ln
}

func (l *logger) WithOutput(fo FormatedOutput) Logger {
	ln := l.copy()
	ln.Formater = fo
	return ln
}

func (l *logger) Debugf(msg string, args ...any) { l.debugf(msg, args) }
func (l *logger) Infof(msg string, args ...any)  { l.infof(msg, args) }
func (l *logger) Warnf(msg string, args ...any)  { l.warnf(msg, args) }
func (l *logger) Errorf(msg string, args ...any) { l.errorf(msg, args) }

func (l *logger) somef(level Level, msg string, args ...any) {
	if l.Context != nil && l.Err() != nil {
		return
	}
	ev := &Event{
		Level:     level,
		Timestamp: time.Now(),
		Fields:    l.Fields,
	}
	if len(args) == 0 {
		ev.Msg = fmt.Sprintf(msg, args...)
	} else {
		ev.Msg = msg
	}
	l.Formater.Consume(ev)
}
