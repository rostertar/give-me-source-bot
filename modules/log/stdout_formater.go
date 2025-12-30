package log

import (
	"context"
	"encoding"
	"fmt"
	"io"
)

type StdoutFormater struct {
	io.Writer

	ctx    context.Context
	cancel context.CancelFunc

	output chan *Event
}

func NewStdoutFormater(ctx context.Context, wr io.Writer, queue_size int) FormatedOutput {
	ctx, cancel := context.WithCancel(ctx)
	sof := &StdoutFormater{
		Writer: wr,
		ctx:    ctx,
		cancel: cancel,
		output: make(chan *Event, queue_size),
	}
	return sof
}

func (sf *StdoutFormater) Close() error {
	if sf.cancel != nil {
		sf.cancel()
	}
	sf.cancel = nil
	return nil
}

func (*StdoutFormater) formatField(f *Field) string {
	start := f.Key + "="
	switch v := f.Value.(type) {
	case string:
		return start + v
	case fmt.Stringer:
		return start + v.String()
	case error:
		return start + "`" + v.Error() + "`"
	case encoding.TextMarshaler:
		s, err := v.MarshalText()
		if err == nil {
			return start + string(s)
		} else {
			return start + "(error marshalling)"
		}
	default:
		return fmt.Sprintf("%s%v", start, v)
	}
}

func (sf *StdoutFormater) fieldsString(fs *Field) string {
	var res string
	for f := range fs.Iterate() {
		if len(res) > 0 {
			res = res + ";"
		}
		res = res + sf.formatField(f)
	}
	return res
}

func (*StdoutFormater) lvlName(l Level) string {
	switch l {
	case LevelError:
		return "ERR"
	case LevelWarn:
		return "WRN"
	case LevelInfo:
		return "INF"
	case LevelDebug:
		return "DBG"
	default:
		return "XXX"
	}
}

func (sf *StdoutFormater) proceed(ctx context.Context) {
	for ctx.Err() == nil {
		select {
		case ev := <-sf.output:
			fields := sf.fieldsString(ev.Fields)
			ts := ev.Timestamp.Format("2006-01-02_15:04:05.000")
			level := sf.lvlName(ev.Level)
			fmt.Fprintf(sf.Writer, "%s %s %s %s\n", ts, level, ev.Msg, fields)
		case <-ctx.Done():
			fmt.Fprintf(sf.Writer, "Logger stoped")
			return
		}
	}
}

func (sf *StdoutFormater) Consume(ev *Event) {
	sf.output <- ev
}
