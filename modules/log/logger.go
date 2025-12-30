package log

import (
	"context"
	"os"
)

var Log Logger
var DefaultLevel = LevelInfo

func init() {
	ctx := context.Background()
	Log = NewDefaultLogger(ctx,
		NewStdoutFormater(ctx, os.Stdout, 40),
		DefaultLevel)
}

type keyForLog struct{}

func ContextWithLogger(ctx context.Context, logger Logger) context.Context {
	return context.WithValue(ctx, keyForLog{}, logger)
}

func To(ctx context.Context) Logger {
	v := ctx.Value(keyForLog{})
	if v == nil {
		return Log
	}
	return v.(Logger)
}

func Debugf(msg string, args ...any) { Log.Debugf(msg, args) }
func Infof(msg string, args ...any)  { Log.Infof(msg, args) }
func Warnf(msg string, args ...any)  { Log.Warnf(msg, args) }
func Errorf(msg string, args ...any) { Log.Errorf(msg, args) }
