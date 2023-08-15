package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("root."))
	})

	http.ListenAndServe(":8082", router)
}
