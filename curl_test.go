package httpietocurl

import "testing"

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
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			got, err := MakeCurlArgs(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			if len(got) != len(tt.want) {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Fatalf("got %v, want %v", got, tt.want)
				}
			}
		})
	}
}
