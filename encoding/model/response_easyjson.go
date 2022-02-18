// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package model

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

func easyjson6ff3ac1dDecodeGithubComMinghsu0107GoElasticsearchExampleEncodingModel(in *jlexer.Lexer, out *SearchResponse) {
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
		case "took":
			out.Took = int64(in.Int64())
		case "hits":
			easyjson6ff3ac1dDecode(in, &out.Hits)
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
func easyjson6ff3ac1dEncodeGithubComMinghsu0107GoElasticsearchExampleEncodingModel(out *jwriter.Writer, in SearchResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"took\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.Took))
	}
	{
		const prefix string = ",\"hits\":"
		out.RawString(prefix)
		easyjson6ff3ac1dEncode(out, in.Hits)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SearchResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6ff3ac1dEncodeGithubComMinghsu0107GoElasticsearchExampleEncodingModel(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SearchResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6ff3ac1dEncodeGithubComMinghsu0107GoElasticsearchExampleEncodingModel(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SearchResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6ff3ac1dDecodeGithubComMinghsu0107GoElasticsearchExampleEncodingModel(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SearchResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6ff3ac1dDecodeGithubComMinghsu0107GoElasticsearchExampleEncodingModel(l, v)
}
func easyjson6ff3ac1dDecode(in *jlexer.Lexer, out *struct {
	Total struct{ Value int64 }
	Hits  []*SearchHit
}) {
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
		case "total":
			easyjson6ff3ac1dDecode1(in, &out.Total)
		case "hits":
			if in.IsNull() {
				in.Skip()
				out.Hits = nil
			} else {
				in.Delim('[')
				if out.Hits == nil {
					if !in.IsDelim(']') {
						out.Hits = make([]*SearchHit, 0, 8)
					} else {
						out.Hits = []*SearchHit{}
					}
				} else {
					out.Hits = (out.Hits)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *SearchHit
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(SearchHit)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.Hits = append(out.Hits, v1)
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
func easyjson6ff3ac1dEncode(out *jwriter.Writer, in struct {
	Total struct{ Value int64 }
	Hits  []*SearchHit
}) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"total\":"
		out.RawString(prefix[1:])
		easyjson6ff3ac1dEncode1(out, in.Total)
	}
	{
		const prefix string = ",\"hits\":"
		out.RawString(prefix)
		if in.Hits == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Hits {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					(*v3).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
func easyjson6ff3ac1dDecode1(in *jlexer.Lexer, out *struct{ Value int64 }) {
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
		case "value":
			out.Value = int64(in.Int64())
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
func easyjson6ff3ac1dEncode1(out *jwriter.Writer, in struct{ Value int64 }) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"value\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.Value))
	}
	out.RawByte('}')
}
func easyjson6ff3ac1dDecodeGithubComMinghsu0107GoElasticsearchExampleEncodingModel1(in *jlexer.Lexer, out *SearchHit) {
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
		case "_score":
			out.Score = float64(in.Float64())
		case "_index":
			out.Index = string(in.String())
		case "_type":
			out.Type = string(in.String())
		case "_version":
			out.Version = int64(in.Int64())
		case "_source":
			(out.Source).UnmarshalEasyJSON(in)
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
func easyjson6ff3ac1dEncodeGithubComMinghsu0107GoElasticsearchExampleEncodingModel1(out *jwriter.Writer, in SearchHit) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"_score\":"
		out.RawString(prefix[1:])
		out.Float64(float64(in.Score))
	}
	{
		const prefix string = ",\"_index\":"
		out.RawString(prefix)
		out.String(string(in.Index))
	}
	{
		const prefix string = ",\"_type\":"
		out.RawString(prefix)
		out.String(string(in.Type))
	}
	if in.Version != 0 {
		const prefix string = ",\"_version\":"
		out.RawString(prefix)
		out.Int64(int64(in.Version))
	}
	{
		const prefix string = ",\"_source\":"
		out.RawString(prefix)
		(in.Source).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SearchHit) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6ff3ac1dEncodeGithubComMinghsu0107GoElasticsearchExampleEncodingModel1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SearchHit) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6ff3ac1dEncodeGithubComMinghsu0107GoElasticsearchExampleEncodingModel1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SearchHit) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6ff3ac1dDecodeGithubComMinghsu0107GoElasticsearchExampleEncodingModel1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SearchHit) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6ff3ac1dDecodeGithubComMinghsu0107GoElasticsearchExampleEncodingModel1(l, v)
}
func easyjson6ff3ac1dDecodeGithubComMinghsu0107GoElasticsearchExampleEncodingModel2(in *jlexer.Lexer, out *IndexResponse) {
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
		case "_index":
			out.Index = string(in.String())
		case "_id":
			out.ID = string(in.String())
		case "_version":
			out.Version = int(in.Int())
		case "result":
			out.Result = string(in.String())
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
func easyjson6ff3ac1dEncodeGithubComMinghsu0107GoElasticsearchExampleEncodingModel2(out *jwriter.Writer, in IndexResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"_index\":"
		out.RawString(prefix[1:])
		out.String(string(in.Index))
	}
	{
		const prefix string = ",\"_id\":"
		out.RawString(prefix)
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"_version\":"
		out.RawString(prefix)
		out.Int(int(in.Version))
	}
	{
		const prefix string = ",\"result\":"
		out.RawString(prefix)
		out.String(string(in.Result))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v IndexResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6ff3ac1dEncodeGithubComMinghsu0107GoElasticsearchExampleEncodingModel2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v IndexResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6ff3ac1dEncodeGithubComMinghsu0107GoElasticsearchExampleEncodingModel2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *IndexResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6ff3ac1dDecodeGithubComMinghsu0107GoElasticsearchExampleEncodingModel2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *IndexResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6ff3ac1dDecodeGithubComMinghsu0107GoElasticsearchExampleEncodingModel2(l, v)
}
func easyjson6ff3ac1dDecodeGithubComMinghsu0107GoElasticsearchExampleEncodingModel3(in *jlexer.Lexer, out *ErrorResponse) {
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
		case "error":
			if in.IsNull() {
				in.Skip()
				out.Info = nil
			} else {
				if out.Info == nil {
					out.Info = new(ErrorInfo)
				}
				(*out.Info).UnmarshalEasyJSON(in)
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
func easyjson6ff3ac1dEncodeGithubComMinghsu0107GoElasticsearchExampleEncodingModel3(out *jwriter.Writer, in ErrorResponse) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Info != nil {
		const prefix string = ",\"error\":"
		first = false
		out.RawString(prefix[1:])
		(*in.Info).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ErrorResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6ff3ac1dEncodeGithubComMinghsu0107GoElasticsearchExampleEncodingModel3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ErrorResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6ff3ac1dEncodeGithubComMinghsu0107GoElasticsearchExampleEncodingModel3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ErrorResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6ff3ac1dDecodeGithubComMinghsu0107GoElasticsearchExampleEncodingModel3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ErrorResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6ff3ac1dDecodeGithubComMinghsu0107GoElasticsearchExampleEncodingModel3(l, v)
}
func easyjson6ff3ac1dDecodeGithubComMinghsu0107GoElasticsearchExampleEncodingModel4(in *jlexer.Lexer, out *ErrorInfo) {
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
		case "root_cause":
			if in.IsNull() {
				in.Skip()
				out.RootCause = nil
			} else {
				in.Delim('[')
				if out.RootCause == nil {
					if !in.IsDelim(']') {
						out.RootCause = make([]*ErrorInfo, 0, 8)
					} else {
						out.RootCause = []*ErrorInfo{}
					}
				} else {
					out.RootCause = (out.RootCause)[:0]
				}
				for !in.IsDelim(']') {
					var v4 *ErrorInfo
					if in.IsNull() {
						in.Skip()
						v4 = nil
					} else {
						if v4 == nil {
							v4 = new(ErrorInfo)
						}
						(*v4).UnmarshalEasyJSON(in)
					}
					out.RootCause = append(out.RootCause, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "type":
			out.Type = string(in.String())
		case "reason":
			out.Reason = string(in.String())
		case "phase":
			out.Phase = string(in.String())
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
func easyjson6ff3ac1dEncodeGithubComMinghsu0107GoElasticsearchExampleEncodingModel4(out *jwriter.Writer, in ErrorInfo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"root_cause\":"
		out.RawString(prefix[1:])
		if in.RootCause == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.RootCause {
				if v5 > 0 {
					out.RawByte(',')
				}
				if v6 == nil {
					out.RawString("null")
				} else {
					(*v6).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix)
		out.String(string(in.Type))
	}
	{
		const prefix string = ",\"reason\":"
		out.RawString(prefix)
		out.String(string(in.Reason))
	}
	{
		const prefix string = ",\"phase\":"
		out.RawString(prefix)
		out.String(string(in.Phase))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ErrorInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6ff3ac1dEncodeGithubComMinghsu0107GoElasticsearchExampleEncodingModel4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ErrorInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6ff3ac1dEncodeGithubComMinghsu0107GoElasticsearchExampleEncodingModel4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ErrorInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6ff3ac1dDecodeGithubComMinghsu0107GoElasticsearchExampleEncodingModel4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ErrorInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6ff3ac1dDecodeGithubComMinghsu0107GoElasticsearchExampleEncodingModel4(l, v)
}
