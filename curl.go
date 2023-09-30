package httpietocurl

import "encoding/json"

func MakeCurlArgs(r Request) ([]string, error) {
	args := []string{}
	if r.Method != "" {
		args = append(args, "-X", r.Method)
	}
	for _, h := range r.Headers {
		args = append(args, "-H", h.Key+":"+h.Value)
	}
	if r.Json != nil {
		if j, err := json.Marshal(r.Json); err != nil {
			return nil, err
		} else {
			args = append(args, "-d", string(j))
		}
	}
	if r.Forms != nil {
		for _, f := range r.Forms {
			args = append(args, "--data-urlencode", MarshalForm(f))
		}
	}

	args = append(args, makeUrl(r.Url, r.Queries))
	return args, nil
}

func makeUrl(url string, queries []Query) string {
	return url + MarshalQueries(queries)
}
