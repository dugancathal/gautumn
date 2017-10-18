package web

import "net/http"

func addAllHeaders(w http.ResponseWriter, headers http.Header) {
	for key, vals := range headers {
		for _, val := range vals {
			w.Header().Add(key, val)
		}
	}
}
