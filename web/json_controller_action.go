package web

import (
	"encoding/json"
	"net/http"
)

const JsonContentType = "application/json;charset=utf-8"

// This type denotes a ControllerAction that acts just normal ControllerAction
// save that it is guaranteed to return JSON.
// The final argument (interface{}) must be JSON serializable and the handler
// will auto-set the Content-Type header to application/json;charset=utf8.
type JsonControllerAction func(r *http.Request) (int, http.Header, interface{})

func (c JsonControllerAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	code, headers, body := c(r)
	addAllHeaders(w, headers)
	w.Header().Set("Content-Type", JsonContentType)

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		code = http.StatusInternalServerError
		bodyBytes = []byte(`{"error": "json-serialization"}`)
	}

	w.WriteHeader(code)
	w.Write(bodyBytes)
}
