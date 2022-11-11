package main

import (
	"net/http"
)

func (app *application) listGameHandler(w http.ResponseWriter, r *http.Request) {

	games, _, err := app.models.Game.GetAll()
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"publishers": games}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
