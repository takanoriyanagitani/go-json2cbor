package map2cbor

import (
	"io"

	ac "github.com/fxamacker/cbor/v2"

	m2c "github.com/takanoriyanagitani/go-json2cbor/map2cbor"
)

type MapToCbor struct {
	*ac.Encoder
}

func (a MapToCbor) Encode(mp map[string]any) error {
	return a.Encoder.Encode(mp)
}

func (a MapToCbor) ToConverter() m2c.MapToCbor { return a.Encode }

func MapToCborNew(mode ac.EncMode) func(io.Writer) MapToCbor {
	return func(wtr io.Writer) MapToCbor {
		return MapToCbor{
			Encoder: mode.NewEncoder(wtr),
		}
	}
}

func WriterToEncoder(wtr io.Writer) MapToCbor {
	return MapToCbor{Encoder: ac.NewEncoder(wtr)}
}
