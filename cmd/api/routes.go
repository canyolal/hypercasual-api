package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/status", app.healthCheckHandler)

	router.HandlerFunc(http.MethodPost, "/v1/publishers", app.createPublisherHandler)
	router.HandlerFunc(http.MethodGet, "/v1/publishers/:id", app.showPublisherHandler)
	router.HandlerFunc(http.MethodGet, "/v1/publishers", app.listPublisherHandler)

	router.HandlerFunc(http.MethodGet, "/v1/games", app.listGameHandler)

	return app.enableCORS(router)
}
