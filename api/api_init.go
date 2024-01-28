package api

import (

	"github.com/go-pg/pg/v10"

	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func StartApi(pgdb *pg.DB) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger, middleware.WithValue("DB", pgdb))

	router.Route("/coffe/types", func(r chi.Router) {
		r.Post("/", createCoffeType)
		r.Get("/", getCoffeTypes)
	})

	router.Route("/coffe/bags", func(r chi.Router) {
		r.Post("/", createCoffeBag)
		r.Get("/", getCoffeBags)
	})


	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Up and running"))
	})
	return router
}

