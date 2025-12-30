package log

import (
	"iter"
)

type Field struct {
	Key   string
	Value any
	Next  *Field
}

func NewFiled(key string, value any) *Field {
	return &Field{Key: key, Value: value}
}

func (f *Field) Iterate() iter.Seq[*Field] {
	return func(yield func(V *Field) bool) {
		for p := f; p != nil && yield(p); p = p.Next {
		}
	}
}

func (f *Field) IterateKV() iter.Seq2[string, any] {
	return func(yield func(K string, V any) bool) {
		p := f
		for p != nil && yield(p.Key, p.Value) {
			p = p.Next
		}
	}
}

func (f *Field) With(other *Field) *Field {
	for p := f; p != nil; p = p.Next {
		if p.Next == nil {
			p.Next = other
			break
		}
	}
	return f
}
