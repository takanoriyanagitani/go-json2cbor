package map2cbor

import (
	"context"

	jc "github.com/takanoriyanagitani/go-json2cbor"
)

type MapToCbor func(map[string]any) error

func (m MapToCbor) OutputAll(ctx context.Context, ms jc.JsonMapSource) error {
	var err error = nil
	for jm := range ms(ctx) {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		err = m(jm)
		if nil != err {
			return err
		}
	}
	return nil
}
