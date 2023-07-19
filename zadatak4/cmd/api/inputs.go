package main

import (
	"fmt"
	"net/http"

	"zadatak4.mihailoivic/internal/data"
	"zadatak4.mihailoivic/internal/validator"
)

func (app *application) createInputHandler(w http.ResponseWriter, r *http.Request) {
	var newInput struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Telephone string `json:"telephone"`
	}

	err := app.readJSON(w, r, &newInput)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	input := &data.Input{
		FirstName: newInput.FirstName,
		LastName:  newInput.LastName,
		Telephone: newInput.Telephone,
	}

	v := validator.New()

	if data.ValidateInput(v, input); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = data.Create(input)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	fmt.Fprintf(w, "%+v\n", newInput)
}

func (app *application) getAllInputsHandler(w http.ResponseWriter, r *http.Request) {
	results, err := data.GetAll()
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"results": results}, nil)
	if err != nil {
		app.logger.Print(err)
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getInputHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)

	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	input, err := data.Get(id)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"input": input}, nil)
	if err != nil {
		app.logger.Print(err)
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteInputHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	err = data.Delete(id)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
