package log

import "time"

type Event struct {
	Level
	Timestamp time.Time
	Msg       string
	Fields    *Field
}
