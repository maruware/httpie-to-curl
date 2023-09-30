package httpietocurl

import (
	"encoding/json"
	"testing"
)

func TestParseHttpie(t *testing.T) {
	tests := []struct {
		desc string
		args []string
		want Request
	}{
		{
			desc: "basic",
			args: []string{"http", "GET", "http://example.com"},
			want: Request{
				Method: "GET",
				Url:    "http://example.com",
			},
		},
		{
			desc: "lowercase method",
			args: []string{"http", "get", "http://example.com"},
			want: Request{
				Method: "GET",
				Url:    "http://example.com",
			},
		},
		{
			desc: "post",
			args: []string{"http", "post", "http://example.com"},
			want: Request{
				Method: "POST",
				Url:    "http://example.com",
			},
		},
		{
			desc: "put",
			args: []string{"http", "put", "http://example.com"},
			want: Request{
				Method: "PUT",
				Url:    "http://example.com",
			},
		},
		{
			desc: "delete",
			args: []string{"http", "delete", "http://example.com"},
			want: Request{
				Method: "DELETE",
				Url:    "http://example.com",
			},
		},
		{
			desc: "patch",
			args: []string{"http", "patch", "http://example.com"},
			want: Request{
				Method: "PATCH",
				Url:    "http://example.com",
			},
		},
		{
			desc: "header",
			args: []string{"http", "GET", "http://example.com", "X-Test: 1"},
			want: Request{
				Method: "GET",
				Url:    "http://example.com",
				Headers: []Header{
					{Key: "X-Test", Value: "1"},
				},
			},
		},
		{
			desc: "multiple headers",
			args: []string{"http", "GET", "http://example.com", "X-Test: 1", "X-Test2: 2"},
			want: Request{
				Method: "GET",
				Url:    "http://example.com",
				Headers: []Header{
					{Key: "X-Test", Value: "1"},
					{Key: "X-Test2", Value: "2"},
				},
			},
		},
		{
			desc: "header with space",
			args: []string{"http", "GET", "http://example.com", "X-Test: 1 2"},
			want: Request{
				Method: "GET",
				Url:    "http://example.com",
				Headers: []Header{
					{Key: "X-Test", Value: "1 2"},
				},
			},
		},
		{
			desc: "json string field",
			args: []string{"http", "post", "http://example.com", "foo=bar"},
			want: Request{
				Method: "POST",
				Url:    "http://example.com",
				Json: map[string]any{
					"foo": "bar",
				},
			},
		},
		{
			desc: "json non-string int field",
			args: []string{"http", "post", "http://example.com", "foo:=1"},
			want: Request{
				Method: "POST",
				Url:    "http://example.com",
				Json: map[string]any{
					"foo": 1,
				},
			},
		},
		{
			desc: "query",
			args: []string{"http", "post", "http://example.com", "foo==bar"},
			want: Request{
				Method: "POST",
				Url:    "http://example.com",
				Queries: []Query{
					{Key: "foo", Value: "bar"},
				},
			},
		},
		{
			desc: "multiple queries",
			args: []string{"http", "post", "http://example.com", "foo==bar", "bar==baz"},
			want: Request{
				Method: "POST",
				Url:    "http://example.com",
				Queries: []Query{
					{Key: "foo", Value: "bar"},
					{Key: "bar", Value: "baz"},
				},
			},
		},
		{
			desc: "form",
			args: []string{"http", "--form", "post", "http://example.com", "foo=bar"},
			want: Request{
				Method: "POST",
				Url:    "http://example.com",
				Forms: []Form{
					{Key: "foo", Value: "bar"},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			got := ParseHttpie(tt.args)
			if got.Method != tt.want.Method {
				t.Errorf("ParseHttpie(%v).Method = %v, want %v", tt.args, got.Method, tt.want.Method)
			}
			if got.Url != tt.want.Url {
				t.Errorf("ParseHttpie(%v).Url = %v, want %v", tt.args, got.Url, tt.want.Url)
			}
			if len(got.Headers) != len(tt.want.Headers) {
				t.Errorf("ParseHttpie(%v).Headers = %v, want %v", tt.args, got.Headers, tt.want.Headers)
			}
			for k, v := range got.Headers {
				if tt.want.Headers[k] != v {
					t.Errorf("ParseHttpie(%v).Headers[%v] = %v, want %v", tt.args, k, v, tt.want.Headers[k])
				}
			}
			gotJsonStr, err1 := json.Marshal(got.Json)
			if err1 != nil {
				t.Errorf("Failed to marshal got json: %v", err1)
			}
			wantJsonStr, err2 := json.Marshal(tt.want.Json)
			if err2 != nil {
				t.Errorf("Failed to marshal want json: %v", err2)
			}

			if string(gotJsonStr) != string(wantJsonStr) {
				t.Errorf("ParseHttpie(%v).Json = %v, want %v", tt.args, string(gotJsonStr), string(wantJsonStr))
			}

			if len(got.Queries) != len(tt.want.Queries) {
				t.Errorf("ParseHttpie(%v).Queries = %v, want %v", tt.args, got.Queries, tt.want.Queries)
			}
			for k, v := range got.Queries {
				if tt.want.Queries[k] != v {
					t.Errorf("ParseHttpie(%v).Queries[%v] = %v, want %v", tt.args, k, v, tt.want.Queries[k])
				}
			}
		})
	}
}
