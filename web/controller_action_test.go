package web_test

import (
	"github.com/dugancathal/gautumn/web"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("ControllerAction.ServeHTTP", func() {
	var action = web.ControllerAction(func (r *http.Request) (int, http.Header, []byte) {
		return http.StatusCreated, http.Header{"Content-Type": []string{"application/json"}}, []byte(`hi mom`)
	})

	It("sets the response code, the headers, and the body", func() {
		recorder := httptest.NewRecorder()
		action.ServeHTTP(recorder, nil)

		Expect(recorder.Code).To(Equal(http.StatusCreated))
		Expect(recorder.Header().Get("Content-Type")).To(Equal("application/json"))
		Expect(recorder.Body.String()).To(Equal("hi mom"))
	})

	It("does not overwrite headers that already exist", func() {
		recorder := httptest.NewRecorder()
		recorder.Header().Set("Content-Type", "application/json+hal")
		action.ServeHTTP(recorder, nil)

		Expect(recorder.Header()["Content-Type"]).To(ContainElement("application/json"))
		Expect(recorder.Header()["Content-Type"]).To(ContainElement("application/json+hal"))
	})
})
