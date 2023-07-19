package main

import (
	"net/http"

	"zadatak2.mihailoivic/internal/data"
	"zadatak2.mihailoivic/internal/validator"
)

func (app *application) createInputHandler(w http.ResponseWriter, r *http.Request) {
	var newInput struct {
		Operation string  `json:"operation"`
		Data      []int64 `json:"data"`
	}

	err := app.readJSON(w, r, &newInput)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	input := &data.Input{
		Operation: newInput.Operation,
		Data:      newInput.Data,
	}

	v := validator.New()

	if data.ValidateInput(v, input); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	if input.Operation != "deduplicate" && input.Operation != "getPairs" {
		app.errorResponse(w, r, http.StatusBadRequest, "operation not allowed")
		return
	}

	result, err := data.DoTheOperation(input)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"result": result}, nil)
	if err != nil {
		app.logger.Print(err)
		app.serverErrorResponse(w, r, err)
	}
}
