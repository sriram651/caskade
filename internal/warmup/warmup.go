// Package warmup is a throwaway code that needs to be deleted at the end of Phase 0
package warmup

type Writer struct {
	Name   string
	offset int64
}

func NewWriter(name string) *Writer {
	return &Writer{
		Name: name,
	}
}

func (w *Writer) MoveOffset(bytesDelta int64) {
	w.offset += bytesDelta
}
