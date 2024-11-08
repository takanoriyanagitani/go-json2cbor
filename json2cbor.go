package json2cbor

import (
	"context"
	"iter"
)

type JsonMapSource func(context.Context) iter.Seq[map[string]any]
