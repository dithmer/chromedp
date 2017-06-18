// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package storage

import (
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

func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpStorage(in *jlexer.Lexer, out *GetUsageAndQuotaReturns) {
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
		case "usage":
			out.Usage = float64(in.Float64())
		case "quota":
			out.Quota = float64(in.Float64())
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
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpStorage(out *jwriter.Writer, in GetUsageAndQuotaReturns) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Usage != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"usage\":")
		out.Float64(float64(in.Usage))
	}
	if in.Quota != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"quota\":")
		out.Float64(float64(in.Quota))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetUsageAndQuotaReturns) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpStorage(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetUsageAndQuotaReturns) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpStorage(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetUsageAndQuotaReturns) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpStorage(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetUsageAndQuotaReturns) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpStorage(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpStorage1(in *jlexer.Lexer, out *GetUsageAndQuotaParams) {
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
		case "origin":
			out.Origin = string(in.String())
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
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpStorage1(out *jwriter.Writer, in GetUsageAndQuotaParams) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"origin\":")
	out.String(string(in.Origin))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetUsageAndQuotaParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpStorage1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetUsageAndQuotaParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpStorage1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetUsageAndQuotaParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpStorage1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetUsageAndQuotaParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpStorage1(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpStorage2(in *jlexer.Lexer, out *ClearDataForOriginParams) {
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
		case "origin":
			out.Origin = string(in.String())
		case "storageTypes":
			out.StorageTypes = string(in.String())
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
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpStorage2(out *jwriter.Writer, in ClearDataForOriginParams) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"origin\":")
	out.String(string(in.Origin))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"storageTypes\":")
	out.String(string(in.StorageTypes))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ClearDataForOriginParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpStorage2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ClearDataForOriginParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpStorage2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ClearDataForOriginParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpStorage2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ClearDataForOriginParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpStorage2(l, v)
}
