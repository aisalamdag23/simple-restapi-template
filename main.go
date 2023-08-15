package main

import (
	"net/http"

	"github.com/aisalamdag23/simple-restapi-template/api"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func main() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(render.SetContentType(render.ContentTypeJSON))

	router.Get("/return200", api.GetResponse200)
	router.Get("/return400", api.GetResponse400)
	router.Get("/return500", api.GetResponse500)

	http.ListenAndServe(":8082", router)
}
