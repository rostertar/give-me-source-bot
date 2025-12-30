package log

type FormatedOutput interface {
	Consume(ev *Event)
	Close() error
}
