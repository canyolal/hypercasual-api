package main

import (
	"net/http"

	"github.com/canyolal/hypercasual-inventories/internal/data"
	"github.com/canyolal/hypercasual-inventories/internal/validator"
)

func (app *application) listGameHandler(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Name  string
		Genre string
		data.Filters
	}

	v := validator.New()

	qs := r.URL.Query()

	input.Name = app.readString(qs, "name", "")
	input.Genre = app.readString(qs, "genre", "")

	input.Filters.Page = app.readInt(qs, "page", 1, v)
	input.Filters.PageSize = app.readInt(qs, "page_size", 20, v)

	input.Filters.Sort = app.readString(qs, "sort", "id")

	// Allowed sort options.
	input.Filters.SortSafeList = []string{
		"id", "name", "genre", "publisher_id",
		"-id", "-name", "-genre", "-publisher_id",
	}

	// Validate filter entries
	if data.ValidateFilters(v, &input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	games, _, metadata, err := app.models.Game.GetAllWithFilters(input.Name, input.Genre, input.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"metadata": metadata, "publishers": games}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
