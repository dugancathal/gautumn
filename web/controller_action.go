package web

import "net/http"

type ControllerAction func(r *http.Request) (int, http.Header, []byte)

func (c ControllerAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	code, headers, body := c(r)
	for key, vals := range headers {
		for _, val := range vals {
			w.Header().Set(key, val)
		}
	}
	w.WriteHeader(code)
	w.Write(body)
}
