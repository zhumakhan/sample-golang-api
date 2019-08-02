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

func easyjson7ae0eefaDecodeCarwashesControllers(in *jlexer.Lexer, out *CarTypes) {
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
		case "CarTypes":
			if in.IsNull() {
				in.Skip()
				out.CarTypes = nil
			} else {
				in.Delim('[')
				if out.CarTypes == nil {
					if !in.IsDelim(']') {
						out.CarTypes = make([]models.CarType, 0, 2)
					} else {
						out.CarTypes = []models.CarType{}
					}
				} else {
					out.CarTypes = (out.CarTypes)[:0]
				}
				for !in.IsDelim(']') {
					var v1 models.CarType
					(v1).UnmarshalEasyJSON(in)
					out.CarTypes = append(out.CarTypes, v1)
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
func easyjson7ae0eefaEncodeCarwashesControllers(out *jwriter.Writer, in CarTypes) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"CarTypes\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.CarTypes == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.CarTypes {
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
func (v CarTypes) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7ae0eefaEncodeCarwashesControllers(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CarTypes) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7ae0eefaEncodeCarwashesControllers(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CarTypes) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7ae0eefaDecodeCarwashesControllers(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CarTypes) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7ae0eefaDecodeCarwashesControllers(l, v)
}
