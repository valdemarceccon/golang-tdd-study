package app

import (
	"os"
)

type Tape struct {
	File *os.File
}

func (t *Tape) Write(p []byte) (n int, err error) {
	_ = t.File.Truncate(0)
	_, _ = t.File.Seek(0, 0)
	return t.File.Write(p)
}
