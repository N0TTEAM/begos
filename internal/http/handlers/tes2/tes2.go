package tes2

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/N0TTEAM/begos/internal/http/model"
	"github.com/N0TTEAM/begos/internal/http/utils/response"
	"github.com/go-playground/validator/v10"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var tes2 model.Tes
		err := json.NewDecoder(r.Body).Decode(&tes2)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body")))
			return
		}

		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		if err := validator.New().Struct(&tes2); err != nil {
			validateErr := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErr))
			return
		}
		response.WriteJson(w, http.StatusCreated, map[string]string{"success": "Ok"})
	}
}
