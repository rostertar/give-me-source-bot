package log

import (
	"errors"
	"fmt"
	"strings"
)

var ErrorUnknownLevelValue = errors.New("Level value is illigal")
var ErrorUnknownLevelName = errors.New("Level name unknown")

type Level int

const (
	LevelDisable = Level(iota)
	LevelError
	LevelWarn
	LevelInfo
	LevelDebug
)

func (l *Level) MarshalText() (text []byte, err error) {
	switch *l {
	case LevelDisable:
		return []byte("Disable"), nil
	case LevelError:
		return []byte("Error"), nil
	case LevelWarn:
		return []byte("Warn"), nil
	case LevelInfo:
		return []byte("Info"), nil
	case LevelDebug:
		return []byte("Debug"), nil
	default:
		return []byte{}, ErrorUnknownLevelValue
	}
}

func (l *Level) UnmarshalText(text []byte) error {
	switch strings.ToLower(string(text)) {
	case "disable":
		*l = LevelDisable
	case "error":
		*l = LevelError
	case "warn":
		*l = LevelWarn
	case "info":
		*l = LevelInfo
	case "debug":
		*l = LevelDebug
	default:
		return fmt.Errorf("%w: %s", ErrorUnknownLevelName, string(text))
	}
	return nil
}

/// Вспомогательные методы для использования в качестве параметров-флагов

func (l *Level) String() string {
	txt, err := l.MarshalText()
	if err != nil {
		return err.Error()
	}
	return string(txt)
}

func (l *Level) Set(txt string) error {
	return l.UnmarshalText([]byte(txt))
}

func (l *Level) Get() any {
	return l
}
