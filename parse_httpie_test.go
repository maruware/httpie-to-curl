package httpietocurl

import (
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
				Method:  "GET",
				Url:     "http://example.com",
				Headers: map[string]string{},
			},
		},
		{
			desc: "lowercase method",
			args: []string{"http", "get", "http://example.com"},
			want: Request{
				Method:  "GET",
				Url:     "http://example.com",
				Headers: map[string]string{},
			},
		},
		{
			desc: "post",
			args: []string{"http", "post", "http://example.com"},
			want: Request{
				Method:  "POST",
				Url:     "http://example.com",
				Headers: map[string]string{},
			},
		},
		{
			desc: "put",
			args: []string{"http", "put", "http://example.com"},
			want: Request{
				Method:  "PUT",
				Url:     "http://example.com",
				Headers: map[string]string{},
			},
		},
		{
			desc: "delete",
			args: []string{"http", "delete", "http://example.com"},
			want: Request{
				Method:  "DELETE",
				Url:     "http://example.com",
				Headers: map[string]string{},
			},
		},
		{
			desc: "patch",
			args: []string{"http", "patch", "http://example.com"},
			want: Request{
				Method:  "PATCH",
				Url:     "http://example.com",
				Headers: map[string]string{},
			},
		},
		{
			desc: "header",
			args: []string{"http", "GET", "http://example.com", "X-Test: 1"},
			want: Request{
				Method: "GET",
				Url:    "http://example.com",
				Headers: map[string]string{
					"X-Test": "1",
				},
			},
		},
		{
			desc: "multiple headers",
			args: []string{"http", "GET", "http://example.com", "X-Test: 1", "X-Test2: 2"},
			want: Request{
				Method: "GET",
				Url:    "http://example.com",
				Headers: map[string]string{
					"X-Test":  "1",
					"X-Test2": "2",
				},
			},
		},
		{
			desc: "header with space",
			args: []string{"http", "GET", "http://example.com", "X-Test: 1 2"},
			want: Request{
				Method: "GET",
				Url:    "http://example.com",
				Headers: map[string]string{
					"X-Test": "1 2",
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
		})
	}
}
