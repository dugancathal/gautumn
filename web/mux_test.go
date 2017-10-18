package web_test

import (
	"github.com/dugancathal/gautumn/web"

	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Mux", func() {
	var mux http.Handler

	BeforeEach(func() {
		mux = web.MapRoutes(web.HandlerMap{
			"/users":            getUsers("/users"),
			"/users/{id}":       getUsers("/users/{id}"),
			"/users/{id}/posts": getUsers("/users/{id}/posts"),
		})
	})

	It("matches on exact routes", func() {
		req := httptest.NewRequest(http.MethodGet, "/users", nil)
		rec := httptest.NewRecorder()

		mux.ServeHTTP(rec, req)

		Expect(rec.Code).To(Equal(http.StatusCreated))
		Expect(rec.Body.String()).To(Equal("/users"))
	})

	It("handles single variables in the path", func() {
		req := httptest.NewRequest(http.MethodGet, "/users/5", nil)
		rec := httptest.NewRecorder()

		mux.ServeHTTP(rec, req)

		Expect(rec.Code).To(Equal(http.StatusCreated))
		Expect(rec.Body.String()).To(Equal("/users/{id}"))
	})

	It("handles deeply nested variable/constant paths", func() {
		req := httptest.NewRequest(http.MethodGet, "/users/4/posts", nil)
		rec := httptest.NewRecorder()

		mux.ServeHTTP(rec, req)

		Expect(rec.Code).To(Equal(http.StatusCreated))
		Expect(rec.Body.String()).To(Equal("/users/{id}/posts"))
	})

	It("returns a 404 if none match", func() {
		req := httptest.NewRequest(http.MethodGet, "/lololololol", nil)
		rec := httptest.NewRecorder()

		mux.ServeHTTP(rec, req)

		Expect(rec.Code).To(Equal(http.StatusNotFound))
		Expect(rec.Body.String()).To(Equal(""))
	})
})

func getUsers(res string) web.ControllerAction {
	return func(r *http.Request) (int, http.Header, []byte) {
		return http.StatusCreated, http.Header{}, []byte(res)
	}
}
