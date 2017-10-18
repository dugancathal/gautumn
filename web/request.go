package web

import (
	"net/http"
	"net/url"
	"strings"
)

const ContentTypeForm = "application/x-www-form-urlencoded"

type Route struct {
	Path string
}

func (route *Route) VarsIn(path string) url.Values {
	vars := url.Values{}
	actualParts := strings.Split(path, "/")
	for i, part := range strings.Split(route.Path, "/") {
		if len(part) > 0 && part[0] == '{' {
			vars[part[1:len(part)-1]] = []string{actualParts[i]}
		}
	}
	return vars
}

type Request struct {
	*http.Request
	Route  *Route
	params url.Values
}

func (req *Request) Params() url.Values {
	if len(req.params) > 0 {
		return req.params
	}

	req.ParseForm()
	req.params = req.Form
	for k, val := range req.PostForm {
		req.params[k] = val
	}
	for k, val := range req.extractUrlVals() {
		req.params[k] = val
	}
	return req.params
}

func (req *Request) extractUrlVals() url.Values {
	return req.Route.VarsIn(req.URL.Path)
}
