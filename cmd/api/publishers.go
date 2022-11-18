package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/canyolal/hypercasual-inventories/internal/data"
)

// creates a new publisher for POST /v1/publishers
func (app *application) createPublisherHandler(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Name string `json:"name"`
		Link string `json:"link"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	publisher := &data.Publisher{
		Name: input.Name,
		Link: input.Link,
	}

	err = app.models.Publisher.Insert(publisher)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/publishers/%d", publisher.Id))

	err = app.writeJSON(w, http.StatusCreated, envelope{"publisher": publisher}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

// shows a single resource of publisher GET /v1/publishers/:id
func (app *application) showPublisherHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	publisher, err := app.models.Publisher.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	app.writeJSON(w, http.StatusOK, envelope{"publisher": publisher}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

// list publishers GET /v1/publishers
func (app *application) listPublisherHandler(w http.ResponseWriter, r *http.Request) {

	publishers, err := app.models.Publisher.GetAll()
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"publishers": publishers}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
