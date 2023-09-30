package httpietocurl

import (
	"regexp"
	"strings"
)

// regexp pattern for header
var headerPattern = regexp.MustCompile(`(?P<key>[^:]+):\s*(?P<value>.*)`)

// parse httpie args
func ParseHttpie(args []string) Request {
	r := Request{
		Headers: map[string]string{},
	}
	for _, arg := range args {
		if arg == "http" {
			continue
		}
		upper := strings.ToUpper(arg)
		if upper == "GET" || upper == "POST" || upper == "PUT" || upper == "PATCH" || upper == "DELETE" {
			r.Method = upper
			continue
		}
		if strings.HasPrefix(arg, "http://") || strings.HasPrefix(arg, "https://") {
			r.Url = arg
			continue
		}

		if headerPattern.Match([]byte(arg)) {
			matches := headerPattern.FindStringSubmatch(arg)
			r.Headers[matches[1]] = matches[2]
			continue
		}
	}

	return r
}
