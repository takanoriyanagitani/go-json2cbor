package main

import (
	"bufio"
	"context"
	"io"
	"log"
	"os"

	jc "github.com/takanoriyanagitani/go-json2cbor"

	l "github.com/takanoriyanagitani/go-json2cbor/lines"
	lj "github.com/takanoriyanagitani/go-json2cbor/lines/lines2jsons2maps"

	mc "github.com/takanoriyanagitani/go-json2cbor/map2cbor"
	ma "github.com/takanoriyanagitani/go-json2cbor/map2cbor/amacker"

	aj "github.com/takanoriyanagitani/go-json2cbor/app/json2map2cbor"
)

var lines2maps lj.LinesToMaps = lj.LinesToMapDefault

type IoConfig struct {
	io.Reader
	io.Writer
}

func (i IoConfig) ToLinesSource() l.LinesSource {
	return func(_ context.Context) l.LineIter {
		return l.ReaderToIter(i.Reader)
	}
}

func (i IoConfig) ToJsonMapSource() jc.JsonMapSource {
	return lines2maps.ToJsonMapSource(i.ToLinesSource())
}

func (i IoConfig) ToMapToCbor() mc.MapToCbor {
	return ma.
		WriterToEncoder(i.Writer).
		ToConverter()
}

func (i IoConfig) ToApp() aj.App {
	return aj.App{
		JsonMapSource: i.ToJsonMapSource(),
		MapToCbor:     i.ToMapToCbor(),
	}
}

func rdr2wtr(ctx context.Context, rdr io.Reader, wtr io.Writer) error {
	icfg := IoConfig{
		Reader: rdr,
		Writer: wtr,
	}
	app := icfg.ToApp()
	return app.OutputAll(ctx)
}

func stdin2stdout(ctx context.Context) error {
	var br io.Reader = bufio.NewReader(os.Stdin)
	var bw *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer bw.Flush()
	return rdr2wtr(ctx, br, bw)
}

func main() {
	e := stdin2stdout(context.Background())
	if nil != e {
		log.Printf("%v\n", e)
	}
}
