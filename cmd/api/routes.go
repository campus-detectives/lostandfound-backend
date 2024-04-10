package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc("GET", "/v1/status", app.statusHandler)
	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)
	router.HandlerFunc(http.MethodPost, "/v1/tokens/authentication", app.createAuthenticationTokenHandler)
	router.HandlerFunc(http.MethodPost, "/v1/found", app.addItem)
	router.HandlerFunc(http.MethodGet, "/v1/found/unclaimed", app.getAllUnclaimed)
	router.HandlerFunc(http.MethodPost, "/v1/found/match", app.searchByImage)
	router.HandlerFunc(http.MethodPost, "/v1/found/claim", app.claimItem)
	router.HandlerFunc(http.MethodPost, "/v1/found/contest", app.contestClaim)

	return app.recoverPanic((app.authenticate(router)))

}
