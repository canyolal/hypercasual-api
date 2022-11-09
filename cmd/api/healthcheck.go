package main

import (
	"log"
	"net/http"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	msg := envelope{
		"status": "up",
		"env":    app.config.env,
	}
	err := app.writeJSON(w, http.StatusOK, msg, nil)
	if err != nil {
		log.Print("error writing to JSON")
	}
}
