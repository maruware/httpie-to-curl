package httpietocurl

import (
	"encoding/json"
	"regexp"
	"strconv"
	"strings"
)

// regexp pattern for header
var headerPattern = regexp.MustCompile(`(?P<key>[^:]+):\s*(?P<value>.*)`)

// json string field pattern. e.g. "foo=bar baz"
var jsonFieldPattern = regexp.MustCompile(`(?P<key>[^=]+)=(?P<value>.*)`)

// json non-string field pattern. e.g. age:=29, married:=false, hobbies:='["http", "pies"]', favorite:='{"tool": "HTTPie"}', bookmarks:=@files/data.json, description=@files/text.txt
var jsonNonStringFieldPattern = regexp.MustCompile(`(?P<key>[^:]+):=(?P<value>.*)`)

// parse httpie args
func ParseHttpie(args []string) Request {
	r := Request{}
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

		if jsonNonStringFieldPattern.Match([]byte(arg)) {
			if r.Json == nil {
				r.Json = map[string]any{}
			}
			matches := jsonNonStringFieldPattern.FindStringSubmatch(arg)
			r.Json[matches[1]] = parseJsonNonStringValue(matches[2])
			continue
		}

		if jsonFieldPattern.Match([]byte(arg)) {
			if r.Json == nil {
				r.Json = map[string]any{}
			}
			matches := jsonFieldPattern.FindStringSubmatch(arg)
			r.Json[matches[1]] = matches[2]
			continue
		}

		if headerPattern.Match([]byte(arg)) {
			matches := headerPattern.FindStringSubmatch(arg)
			if r.Headers == nil {
				r.Headers = []Header{}
			}
			r.Headers = append(r.Headers, Header{Key: matches[1], Value: matches[2]})
			continue
		}
	}

	return r
}

func parseJsonNonStringValue(s string) interface{} {
	// int
	if v, err := strconv.ParseInt(s, 0, 64); err == nil {
		return v
	}
	// float
	if v, err := strconv.ParseFloat(s, 64); err == nil {
		return v
	}
	// bool
	if v, err := strconv.ParseBool(s); err == nil {
		return v
	}
	// array
	if strings.HasPrefix(s, "[") && strings.HasSuffix(s, "]") {
		var a []interface{}
		if err := json.Unmarshal([]byte(s), &a); err == nil {
			return a
		}
	}
	// object
	if strings.HasPrefix(s, "{") && strings.HasSuffix(s, "}") {
		var m map[string]interface{}
		if err := json.Unmarshal([]byte(s), &m); err == nil {
			return m
		}
	}

	// otherwise
	return s
}
