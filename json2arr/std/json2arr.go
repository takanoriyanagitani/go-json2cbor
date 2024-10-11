package json2arr

import (
	"encoding/json"
	"io"

	j2a "github.com/takanoriyanagitani/go-json2cbor/json2arr"
)

type JsonToArr struct {
	*json.Decoder
}

func (j JsonToArr) DecodeToArray(buf *[]any) error {
	return j.Decoder.Decode(buf)
}

func (j JsonToArr) ToConverter() j2a.JsonToArray { return j.DecodeToArray }

func JsonToArrNew(rdr io.Reader) JsonToArr {
	return JsonToArr{Decoder: json.NewDecoder(rdr)}
}
