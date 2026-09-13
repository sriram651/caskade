// Package warmup is a throwaway code that needs to be deleted at the end of Phase 0
package warmup

import "errors"

type Writer struct {
	Name   string
	offset int64
}

var ErrNegativeDelta = errors.New("bytes delta cannot be negative")

func NewWriter(name string) *Writer {
	return &Writer{
		Name: name,
	}
}

func (w *Writer) MoveOffset(bytesDelta int64) error {
	if bytesDelta < 0 {
		return ErrNegativeDelta
	}

	w.offset += bytesDelta

	return nil
}
