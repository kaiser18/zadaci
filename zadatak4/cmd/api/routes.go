package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodPost, "/api/inputs", app.createInputHandler)
	router.HandlerFunc(http.MethodGet, "/api/inputs", app.getAllInputsHandler)
	router.HandlerFunc(http.MethodGet, "/api/inputs/:id", app.getInputHandler)
	router.HandlerFunc(http.MethodDelete, "/api/inputs/:id", app.deleteInputHandler)

	return router
}
