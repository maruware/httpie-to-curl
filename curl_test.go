package httpietocurl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeCurlArgs(t *testing.T) {
	tests := []struct {
		desc  string
		input Request
		want  []string
	}{
		{
			desc: "basic",
			input: Request{
				Url: "http://example.com",
			},
			want: []string{"http://example.com"},
		},
		{
			desc: "method",
			input: Request{
				Method: "GET",
				Url:    "http://example.com",
			},
			want: []string{"-X", "GET", "http://example.com"},
		},
		{
			desc: "headers",
			input: Request{
				Url: "http://example.com",
				Headers: []Header{
					{"Content-Type", "application/json"},
					{"Accept", "application/json"},
				},
			},
			want: []string{"-H", "Content-Type:application/json", "-H", "Accept:application/json", "http://example.com"},
		},
		{
			desc: "json",
			input: Request{
				Url: "http://example.com",
				Json: map[string]any{
					"foo": "bar",
					"baz": 1,
				},
			},
			want: []string{"-d", `{"baz":1,"foo":"bar"}`, "http://example.com"},
		},
		{
			desc: "query",
			input: Request{
				Url: "http://example.com",
				Queries: []Query{
					{"foo", "bar"},
					{"baz", "1"},
				},
			},
			want: []string{"http://example.com?foo=bar&baz=1"},
		},
		{
			desc: "query with space",
			input: Request{
				Url: "http://example.com",
				Queries: []Query{
					{"foo", "bar 1"},
				},
			},
			want: []string{"http://example.com?foo=bar+1"},
		},
		{
			desc: "query with japanese",
			input: Request{
				Url: "http://example.com",
				Queries: []Query{
					{"foo", "あ"},
				},
			},
			want: []string{"http://example.com?foo=%E3%81%82"},
		},
		{
			desc: "form",
			input: Request{
				Url: "http://example.com",
				Forms: []Form{
					{"foo", "bar"},
					{"baz", "1"},
				},
			},
			want: []string{"--data-urlencode", "foo=bar", "--data-urlencode", "baz=1", "http://example.com"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			got, err := MakeCurlArgs(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
