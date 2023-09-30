package httpietocurl

import "net/url"

type Header struct {
	Key   string
	Value string
}

type Query struct {
	Key   string
	Value string
}

type Form struct {
	Key   string
	Value string
}

type Request struct {
	Method  string
	Url     string
	Headers []Header
	Json    map[string]any
	Queries []Query
	Forms   []Form
}

func MarshalQuery(q Query) string {
	return q.Key + "=" + url.QueryEscape(q.Value)
}

func MarshalQueries(queries []Query) string {
	if len(queries) == 0 {
		return ""
	}

	q := ""
	for _, query := range queries {
		if q == "" {
			q += "?"
		} else {
			q += "&"
		}
		q += MarshalQuery(query)
	}
	return q
}

func MarshalForm(f Form) string {
	return f.Key + "=" + f.Value
}

func (r *Request) AddHeader(h Header) {
	if r.Headers == nil {
		r.Headers = []Header{}
	}

	// overwrite if header already exists
	for i, header := range r.Headers {
		if header.Key == h.Key {
			r.Headers[i] = h
			return
		}
	}

	r.Headers = append(r.Headers, h)
}

func (r *Request) AddJsonField(key string, value any) {
	if r.Json == nil {
		r.Json = map[string]any{}
	}
	r.Json[key] = value
}

func (r *Request) AddQuery(q Query) {
	if r.Queries == nil {
		r.Queries = []Query{}
	}
	r.Queries = append(r.Queries, q)
}

func (r *Request) AddForm(f Form) {
	if r.Forms == nil {
		r.Forms = []Form{}
	}
	r.Forms = append(r.Forms, f)
}
