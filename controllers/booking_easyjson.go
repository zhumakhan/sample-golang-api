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

func easyjson6341a807DecodeCarwashesControllers(in *jlexer.Lexer, out *PastBookingsResponse) {
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
		case "Bookings":
			if in.IsNull() {
				in.Skip()
				out.Bookings = nil
			} else {
				in.Delim('[')
				if out.Bookings == nil {
					if !in.IsDelim(']') {
						out.Bookings = make([]models.PastBooking, 0, 1)
					} else {
						out.Bookings = []models.PastBooking{}
					}
				} else {
					out.Bookings = (out.Bookings)[:0]
				}
				for !in.IsDelim(']') {
					var v1 models.PastBooking
					easyjson6341a807DecodeCarwashesModels(in, &v1)
					out.Bookings = append(out.Bookings, v1)
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
func easyjson6341a807EncodeCarwashesControllers(out *jwriter.Writer, in PastBookingsResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Bookings\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Bookings == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Bookings {
				if v2 > 0 {
					out.RawByte(',')
				}
				easyjson6341a807EncodeCarwashesModels(out, v3)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PastBookingsResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6341a807EncodeCarwashesControllers(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PastBookingsResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6341a807EncodeCarwashesControllers(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PastBookingsResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6341a807DecodeCarwashesControllers(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PastBookingsResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6341a807DecodeCarwashesControllers(l, v)
}
func easyjson6341a807DecodeCarwashesModels(in *jlexer.Lexer, out *models.PastBooking) {
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
		case "ID":
			out.ID = uint(in.Uint())
		case "CarWash":
			out.CarWash = uint(in.Uint())
		case "ClientUUID":
			out.ClientUUID = string(in.String())
		case "ClientFirstName":
			out.ClientFirstName = string(in.String())
		case "ClientSecondName":
			out.ClientSecondName = string(in.String())
		case "ClientMiddleName":
			out.ClientMiddleName = string(in.String())
		case "ClientPhone":
			out.ClientPhone = string(in.String())
		case "CarNumber":
			out.CarNumber = string(in.String())
		case "Cost":
			out.Cost = uint(in.Uint())
		case "CreatedAt":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreatedAt).UnmarshalJSON(data))
			}
		case "UpdatedAt":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.UpdatedAt).UnmarshalJSON(data))
			}
		case "CarModel":
			out.CarModel = string(in.String())
		case "BookingServices":
			out.BookingServices = string(in.String())
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
func easyjson6341a807EncodeCarwashesModels(out *jwriter.Writer, in models.PastBooking) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"ID\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint(uint(in.ID))
	}
	{
		const prefix string = ",\"CarWash\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint(uint(in.CarWash))
	}
	{
		const prefix string = ",\"ClientUUID\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ClientUUID))
	}
	{
		const prefix string = ",\"ClientFirstName\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ClientFirstName))
	}
	{
		const prefix string = ",\"ClientSecondName\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ClientSecondName))
	}
	{
		const prefix string = ",\"ClientMiddleName\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ClientMiddleName))
	}
	{
		const prefix string = ",\"ClientPhone\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ClientPhone))
	}
	{
		const prefix string = ",\"CarNumber\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.CarNumber))
	}
	{
		const prefix string = ",\"Cost\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint(uint(in.Cost))
	}
	{
		const prefix string = ",\"CreatedAt\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.CreatedAt).MarshalJSON())
	}
	{
		const prefix string = ",\"UpdatedAt\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.UpdatedAt).MarshalJSON())
	}
	{
		const prefix string = ",\"CarModel\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.CarModel))
	}
	{
		const prefix string = ",\"BookingServices\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.BookingServices))
	}
	out.RawByte('}')
}
func easyjson6341a807DecodeCarwashesControllers1(in *jlexer.Lexer, out *BookingsResponse) {
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
		case "Bookings":
			if in.IsNull() {
				in.Skip()
				out.Bookings = nil
			} else {
				in.Delim('[')
				if out.Bookings == nil {
					if !in.IsDelim(']') {
						out.Bookings = make([]models.Booking, 0, 1)
					} else {
						out.Bookings = []models.Booking{}
					}
				} else {
					out.Bookings = (out.Bookings)[:0]
				}
				for !in.IsDelim(']') {
					var v4 models.Booking
					(v4).UnmarshalEasyJSON(in)
					out.Bookings = append(out.Bookings, v4)
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
func easyjson6341a807EncodeCarwashesControllers1(out *jwriter.Writer, in BookingsResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Bookings\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Bookings == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Bookings {
				if v5 > 0 {
					out.RawByte(',')
				}
				(v6).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v BookingsResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6341a807EncodeCarwashesControllers1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BookingsResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6341a807EncodeCarwashesControllers1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *BookingsResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6341a807DecodeCarwashesControllers1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BookingsResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6341a807DecodeCarwashesControllers1(l, v)
}
