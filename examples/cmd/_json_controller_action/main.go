package main

import (
	"net/http"

	"fmt"
	"log"

	"github.com/dugancathal/gautumn/gautumn/web"
)

type foo struct {
	Field1 string `json:"field1"`
	Field2 string `json:"field2"`
}

type fooController struct{}

func (c *fooController) Index() web.JsonControllerAction {
	return func(r *http.Request) (int, http.Header, interface{}) {
		return 200, http.Header{"X-Test": []string{"Hi"}}, &foo{"hello", "hey"}
	}
}

func main() {
	controller := &fooController{}

	web.MapRoutes(web.RouteMap{
		"/movies": controller.Index(),
	})
	fmt.Printf("Serving on %d\n", 8080)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
