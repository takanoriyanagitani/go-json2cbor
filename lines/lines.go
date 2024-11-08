package lines

import (
	"bufio"
	"context"
	"io"
	"iter"
)

type Line []byte

type LineIter iter.Seq[Line]

func ScannerToIter(s *bufio.Scanner) iter.Seq[Line] {
	return func(yield func(Line) bool) {
		for s.Scan() {
			var line []byte = s.Bytes()
			if !yield(line) {
				return
			}
		}
	}
}

func ReaderToIter(rdr io.Reader) LineIter {
	var scanner *bufio.Scanner = bufio.NewScanner(rdr)
	return LineIter(ScannerToIter(scanner))
}

type LinesSource func(context.Context) LineIter
