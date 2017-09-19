package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"database/sql"

	"os"

	"github.com/dugancathal/gautumn"
	"github.com/dugancathal/gautumn/web"
	_ "github.com/lib/pq"
)

// application code
type repo interface {
	FindAll() []string
}

type movieRepo struct {
	db *sql.DB
}

func newMovieRepo(db *sql.DB) repo {
	return &movieRepo{db}
}

func (r *movieRepo) FindAll() []string {
	res, err := r.db.Query("SELECT name FROM movies")
	if err != nil {
		panic(err)
	}
	movieNames := []string{}
	defer res.Close()
	for res.Next() {
		var name string
		res.Scan(&name)
		movieNames = append(movieNames, name)
	}
	return movieNames
}

type moviesController struct{}

func (c *moviesController) Index(repo repo) web.ControllerAction {
	return func(r *http.Request) (int, http.Header, []byte) {
		movies := repo.FindAll()
		body, _ := json.Marshal(movies)
		return 200, http.Header{}, []byte(body)
	}
}

func (c *moviesController) InjectedIndex() web.ControllerAction {
	return gautumn.DefaultInjected(c.Index).(web.ControllerAction)
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: go run main.go DBNAME")
	}

	dbName := os.Args[1]
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://@localhost/%s?sslmode=disable", dbName))
	if err != nil {
		log.Fatalf("Failed to open db %#v\n", err)
	}

	gautumn.RegisterByType(db)
	gautumn.RegisterByConstructor(newMovieRepo)

	controller := &moviesController{}

	http.Handle("/movies", controller.InjectedIndex())
	fmt.Printf("Serving on %d\n", 8080)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
