package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dugancathal/gautumn/gautumn"
	"github.com/dugancathal/gautumn/gautumn/web"
)

// application code
type repo interface {
	FindAll() []string
}

type movieRepo struct {
}

func (r *movieRepo) FindAll() []string {
	return []string{"Braveheart"}
}

type moviesController struct{}

func (c *moviesController) Index(repo repo) web.ControllerAction {
	return func(r *http.Request) (int, http.Header, []byte) {
		movies := repo.FindAll()
		body, _ := json.Marshal(movies)
		return 200, http.Header{}, []byte(body)
	}
}

func (c *moviesController) InjectedIndex() web.JsonControllerAction {
	return gautumn.DefaultInjected(c.Index).(web.JsonControllerAction)
}

func main() {
	gautumn.DefaultContainer.RegisterByInterface((*repo)(nil), &movieRepo{})

	controller := &moviesController{}

	http.Handle("/movies", controller.InjectedIndex())
	fmt.Printf("Serving on %d\n", 8080)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
