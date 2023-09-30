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
