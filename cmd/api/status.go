package main

import "net/http"

func (app *application) statusHandler(w http.ResponseWriter, r *http.Request) {
	data := envelope{
		"status":  "available",
		"version": apiVersion,
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
