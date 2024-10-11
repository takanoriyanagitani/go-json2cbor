package arr2cbor

import (
	"io"

	ac "github.com/fxamacker/cbor/v2"

	a2c "github.com/takanoriyanagitani/go-json2cbor/arr2cbor"
)

type ArrToCbor struct {
	*ac.Encoder
}

func (a ArrToCbor) EncodeArray(arr []any) error { return a.Encoder.Encode(arr) }

func (a ArrToCbor) ToConverter() a2c.ArrayToCbor { return a.EncodeArray }

func ArrToCborNew(mode ac.EncMode) func(io.Writer) ArrToCbor {
	return func(wtr io.Writer) ArrToCbor {
		return ArrToCbor{
			Encoder: mode.NewEncoder(wtr),
		}
	}
}
