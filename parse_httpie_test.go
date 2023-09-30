package httpietocurl

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
				Headers: []Header{
					{Key: "Content-Type", Value: "application/json"},
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
				Headers: []Header{
					{Key: "Content-Type", Value: "application/json"},
				},
			},
		},
		{
			desc: "json non-string float field",
			args: []string{"http", "post", "http://example.com", "foo:=1.2"},
			want: Request{
				Method: "POST",
				Url:    "http://example.com",
				Json: map[string]any{
					"foo": 1.2,
				},
				Headers: []Header{
					{Key: "Content-Type", Value: "application/json"},
				},
			},
		},
		{
			desc: "json non-string array field",
			args: []string{"http", "post", "http://example.com", "foo:=[1,2,3]"},
			want: Request{
				Method: "POST",
				Url:    "http://example.com",
				Json: map[string]any{
					"foo": []any{float64(1), float64(2), float64(3)},
				},
				Headers: []Header{
					{Key: "Content-Type", Value: "application/json"},
				},
			},
		},
		{
			desc: "json non-string object field",
			args: []string{"http", "post", "http://example.com", "foo:={\"bar\":1}"},
			want: Request{
				Method: "POST",
				Url:    "http://example.com",
				Json: map[string]any{
					"foo": map[string]any{
						"bar": float64(1),
					},
				},
				Headers: []Header{
					{Key: "Content-Type", Value: "application/json"},
				},
			},
		},
		{
			desc: "json non-string bool field",
			args: []string{"http", "post", "http://example.com", "foo:=true"},
			want: Request{
				Method: "POST",
				Url:    "http://example.com",
				Json: map[string]any{
					"foo": true,
				},
				Headers: []Header{
					{Key: "Content-Type", Value: "application/json"},
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
			assert.Equal(t, tt.want.Headers, got.Headers)
			assert.Equal(t, tt.want.Json, got.Json)
			assert.Equal(t, tt.want.Forms, got.Forms)
			assert.Equal(t, tt.want.Queries, got.Queries)
		})
	}
}
