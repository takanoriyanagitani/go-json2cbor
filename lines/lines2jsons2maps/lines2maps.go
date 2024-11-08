package lines2maps

import (
	"context"
	"encoding/json"
	"iter"

	jc "github.com/takanoriyanagitani/go-json2cbor"

	l "github.com/takanoriyanagitani/go-json2cbor/lines"
)

type LinesToMaps func(l.LineIter) iter.Seq[map[string]any]

func LinesToMapsStd(lines l.LineIter) iter.Seq[map[string]any] {
	return func(yield func(map[string]any) bool) {
		var buf map[string]any
		var err error = nil

		for line := range lines {
			clear(buf)
			err = json.Unmarshal(line, &buf)
			if nil != err {
				return
			}

			if !yield(buf) {
				return
			}
		}
	}
}

var LinesToMapDefault LinesToMaps = LinesToMapsStd

func (m LinesToMaps) ToJsonMapSource(ls l.LinesSource) jc.JsonMapSource {
	return func(ctx context.Context) iter.Seq[map[string]any] {
		var lines l.LineIter = ls(ctx)
		return m(lines)
	}
}
