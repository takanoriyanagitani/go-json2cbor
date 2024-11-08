package json2map2cbor

import (
	"context"

	jc "github.com/takanoriyanagitani/go-json2cbor"

	mc "github.com/takanoriyanagitani/go-json2cbor/map2cbor"
)

type App struct {
	jc.JsonMapSource
	mc.MapToCbor
}

func (a App) OutputAll(ctx context.Context) error {
	return a.MapToCbor.OutputAll(ctx, a.JsonMapSource)
}
