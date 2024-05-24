package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"log"
	"my-app/db"
	"my-app/handlers"
	"net/http"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	if err := db.CreateDatabase(); err != nil {
		log.Fatal(err)
	}

	db.Migrate()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello"))
		if err != nil {
			return
		}
	})

	r.Group(func(r chi.Router) {
		r.Post("/login", handlers.Login)
	})

	r.Group(func(r chi.Router) {
		r.Get("/user", handlers.GetUsers)
		r.Get("/user/{id}", handlers.GetUser)
		r.Post("/user", handlers.StoreUser)
	})

	err := http.ListenAndServe(":4000", r)
	if err != nil {
		return
	}
}
