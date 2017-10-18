package web_test

import (
	"net/http"

	"bytes"
	"net/url"

	"github.com/dugancathal/gautumn/web"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Request", func() {
	Describe("#Params()", func() {
		route := &web.Route{Path: "/foo/bar"}

		It("returns params from the query string", func() {
			r, _ := http.NewRequest("GET", "/foo/bar?baz=bork", nil)

			request := web.Request{Request: r, Route: route}
			Expect(request.Params()["baz"]).To(Equal([]string{"bork"}))
		})

		It("returns params from the form body", func() {
			form := url.Values{"barking": []string{"mad"}}
			r, _ := http.NewRequest("POST", "/foo/bar", bytes.NewBufferString(form.Encode()))
			r.Header.Set("Content-Type", web.ContentTypeForm)

			request := web.Request{Request: r, Route: route}
			Expect(request.Params()["barking"]).To(Equal([]string{"mad"}))
		})

		It("returns url variables", func() {
			form := url.Values{"barking": []string{"mad"}}
			r, _ := http.NewRequest("POST", "/users/3", bytes.NewBufferString(form.Encode()))
			r.Header.Set("Content-Type", web.ContentTypeForm)

			request := web.Request{Request: r, Route: &web.Route{Path: "/users/{user-id}"}}
			Expect(request.Params()["user-id"]).To(Equal([]string{"3"}))
		})

		It("matches nested urls with vars", func() {
			form := url.Values{"barking": []string{"mad"}}
			r, _ := http.NewRequest("POST", "/users/3/blogs/1", bytes.NewBufferString(form.Encode()))
			r.Header.Set("Content-Type", web.ContentTypeForm)

			request := web.Request{Request: r, Route: &web.Route{Path: "/users/{user-id}/blogs/{blog-id}"}}
			Expect(request.Params()["user-id"]).To(Equal([]string{"3"}))
			Expect(request.Params()["blog-id"]).To(Equal([]string{"1"}))
		})
	})
})
