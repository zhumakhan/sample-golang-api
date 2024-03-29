// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package controllers

import (
	models "carwashes/models"
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson5c8673aaDecodeCarwashesControllers(in *jlexer.Lexer, out *Clients) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Clients":
			if in.IsNull() {
				in.Skip()
				out.Clients = nil
			} else {
				in.Delim('[')
				if out.Clients == nil {
					if !in.IsDelim(']') {
						out.Clients = make([]models.Client, 0, 1)
					} else {
						out.Clients = []models.Client{}
					}
				} else {
					out.Clients = (out.Clients)[:0]
				}
				for !in.IsDelim(']') {
					var v1 models.Client
					(v1).UnmarshalEasyJSON(in)
					out.Clients = append(out.Clients, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjson5c8673aaEncodeCarwashesControllers(out *jwriter.Writer, in Clients) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Clients\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Clients == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Clients {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Clients) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5c8673aaEncodeCarwashesControllers(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Clients) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5c8673aaEncodeCarwashesControllers(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Clients) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5c8673aaDecodeCarwashesControllers(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Clients) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5c8673aaDecodeCarwashesControllers(l, v)
}
