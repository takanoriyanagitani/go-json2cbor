package main

import (
	"bufio"
	"context"
	"io"
	"log"
	"os"

	ac "github.com/fxamacker/cbor/v2"

	a2c "github.com/takanoriyanagitani/go-json2cbor/arr2cbor"
	a2ca "github.com/takanoriyanagitani/go-json2cbor/arr2cbor/amacker"

	j2a "github.com/takanoriyanagitani/go-json2cbor/json2arr"
	j2as "github.com/takanoriyanagitani/go-json2cbor/json2arr/std"

	j2c "github.com/takanoriyanagitani/go-json2cbor/json2arr2cbor"
)

func timeModeFromString(ts string) ac.TimeMode {
	switch ts {
	case "TimeUnix":
		return ac.TimeUnix
	case "TimeUnixMicro":
		return ac.TimeUnixMicro
	case "TimeUnixDynamic":
		return ac.TimeUnixDynamic
	case "TimeRFC3339":
		return ac.TimeRFC3339
	case "TimeRFC3339Nano":
		return ac.TimeRFC3339Nano
	default:
		return ac.TimeUnix
	}
}

var timeMode ac.TimeMode = timeModeFromString(os.Getenv("ENV_CBOR_TIME_MODE"))

type app struct {
	j2a.JsonToArray
	a2c.ArrayToCbor
}

func (a app) ToConverter() j2c.JsonToArrayToCbor {
	return j2c.JsonToArrayToCbor{
		JsonToArray: a.JsonToArray,
		ArrayToCbor: a.ArrayToCbor,
	}
}

func (a app) ConvertAll(ctx context.Context) error {
	return a.ToConverter().ConvertAll(ctx)
}

func (a app) WithReader(rdr io.Reader) app {
	a.JsonToArray = j2as.JsonToArrNew(rdr).ToConverter()
	return a
}

func (a app) WithMode(m ac.EncMode, wtr io.Writer) app {
	a.ArrayToCbor = a2ca.ArrToCborNew(m)(wtr).ToConverter()
	return a
}

func (a app) WithOptsMust(o ac.EncOptions, wtr io.Writer) app {
	em, err := o.EncMode()
	if nil != err {
		panic(err)
	}
	return a.WithMode(em, wtr)
}

func (a app) WithTimeModeMust(t ac.TimeMode, wtr io.Writer) app {
	var opts ac.EncOptions = ac.CanonicalEncOptions()
	opts.Time = t
	return a.WithOptsMust(opts, wtr)
}

func reader2writer(ctx context.Context, rdr io.Reader, wtr io.Writer) error {
	var a app = app{}.
		WithReader(rdr).
		WithTimeModeMust(timeMode, wtr)
	var cnv j2c.JsonToArrayToCbor = a.ToConverter()
	return cnv.ConvertAll(ctx)
}

func stdin2stdout(ctx context.Context) error {
	var rdr io.Reader = bufio.NewReader(os.Stdin)
	var wtr *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer wtr.Flush()
	return reader2writer(ctx, rdr, wtr)
}

func main() {
	e := stdin2stdout(context.Background())
	if nil != e {
		log.Printf("%v\n", e)
	}
}
