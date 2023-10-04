package main

import (
	"assignment-mezink/datasources"
	"assignment-mezink/handlers"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	//initialize db
	datasources.ConnectDB()
	if datasources.Db == nil {
		log.Fatalln("fail to connect database")
	}

	//initalize router
	r := chi.NewRouter()
	r.Route("/records", func(r chi.Router) {
        r.Get("/",handlers.GetAllRecordHandler)
		r.Get("/{id}", handlers.GetRecordHandler)
		r.Post("/", handlers.CreateRecordHandler)
		r.Post("/sum", handlers.GetSumRecordHandler)
	})

	http.ListenAndServe(":8080", r)
}
