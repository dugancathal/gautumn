package web

import (
	"net/http"
	"strings"
)

type Mux struct {
	routeMap HandlerMap
}

func (mux *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if handler, exact := mux.routeMap[r.URL.Path]; exact {
		handler.ServeHTTP(w, r)
		return
	}

	for path, handler := range mux.routeMap {
		if matches(r, path) {
			handler.ServeHTTP(w, r)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func matches(r *http.Request, path string) bool {
	matchParts := strings.Split(path, "/")
	actualParts := strings.Split(r.URL.Path, "/")
	if len(matchParts) != len(actualParts) {
		return false
	}

	for i, actualPart := range actualParts {
		if actualPart == matchParts[i] || matchParts[i][0] == '{' {
			continue
		}
		return false
	}
	return true
}
