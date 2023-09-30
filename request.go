package httpietocurl

type Header struct {
	Key   string
	Value string
}

type Request struct {
	Method  string
	Url     string
	Headers []Header
	Json    map[string]any
}
