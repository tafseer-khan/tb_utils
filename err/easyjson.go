package err

import (
	"errors"

	"encoding/json"

	"github.com/mailru/easyjson"

	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson6601e8cdDecodeFunction(in *jlexer.Lexer, out *errStruct) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Error":
			out.Error = errors.New(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6601e8cdEncodeFunction(out *jwriter.Writer, in errStruct) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Error\":"
		out.RawString(prefix[1:])
		out.String(in.Error.Error())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v errStruct) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeFunction(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v errStruct) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeFunction(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *errStruct) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeFunction(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *errStruct) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeFunction(l, v)
}
