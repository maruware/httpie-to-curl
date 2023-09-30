package httpietocurl

type Request struct {
	Method  string
	Url     string
	Headers map[string]string
	Json    map[string]any
}
