package web_test

import (
	"github.com/dugancathal/gautumn/web"

	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("JsonControllerAction", func() {
	type foo struct {
		Name string
	}
	var action = web.JsonControllerAction(func(r *http.Request) (int, http.Header, interface{}) {
		return http.StatusAccepted,
			http.Header{"Content-Type": []string{"text/html"}},
			&foo{Name: "imma-foo"}
	})

	It("overrides the content-type to UTF-8 JSON", func() {
		recorder := httptest.NewRecorder()

		action.ServeHTTP(recorder, nil)

		Expect(recorder.Header()["Content-Type"]).To(Equal([]string{web.JsonContentType}))
	})

	It("marshals the 'interface{}' as json", func() {
		recorder := httptest.NewRecorder()

		action.ServeHTTP(recorder, nil)

		Expect(recorder.Body.String()).To(Equal(`{"Name":"imma-foo"}`))
	})

	It("sets the response code", func() {
		recorder := httptest.NewRecorder()

		action.ServeHTTP(recorder, nil)

		Expect(recorder.Code).To(Equal(http.StatusAccepted))
	})
})
