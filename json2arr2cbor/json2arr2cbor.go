package json2arr2cbor

import (
	"context"
	"errors"
	"io"

	a2c "github.com/takanoriyanagitani/go-json2cbor/arr2cbor"
	j2a "github.com/takanoriyanagitani/go-json2cbor/json2arr"
)

type JsonToArrayToCbor struct {
	j2a.JsonToArray
	a2c.ArrayToCbor
}

func (j JsonToArrayToCbor) Convert(buf *[]any) error {
	var edec error = j.JsonToArray(buf)
	if nil != edec {
		return edec
	}

	return j.ArrayToCbor(*buf)
}

func (j JsonToArrayToCbor) ConvertAll(ctx context.Context) error {
	var buf []any
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		buf = buf[:0]
		e := j.Convert(&buf)
		if nil != e {
			if !errors.Is(e, io.EOF) {
				return e
			}
			return nil
		}
	}
}
