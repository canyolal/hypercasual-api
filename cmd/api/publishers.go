package main

import (
	"fmt"
	"net/http"

	"github.com/canyolal/hypercasual-inventories/internal/data"
)

func (app *application) createPublisherHandler(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Name string `json:"name"`
		Link string `json:"link"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.serverErrorResponse(w, r, err)
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
