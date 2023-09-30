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
	args = append(args, r.Url)
	return args, nil
}
