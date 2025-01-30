package tes2

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/N0TTEAM/begos/internal/db"
	"github.com/N0TTEAM/begos/internal/http/model"
	"github.com/N0TTEAM/begos/internal/http/utils/response"
	"github.com/go-playground/validator/v10"
)

func CreateTes() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var tes model.Tes
		err := json.NewDecoder(r.Body).Decode(&tes)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body")))
			return
		}

		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		if err := validator.New().Struct(&tes); err != nil {
			validateErr := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErr))
			return
		}

		if err := db.GetDB().Create(&tes).Error; err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(fmt.Errorf("failed to insert data: %v", err)))
		}

		response.WriteJson(w, http.StatusCreated, map[string]interface{}{"success": "OK", "data": tes})
	}
}

func DeleteTesById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")

		if id == "" {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("ID not found")))
			return
		}

		tesID, err := strconv.Atoi(id)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("ID not valid")))
			return
		}

		var tes model.Tes
		if err := db.GetDB().First(&tes, tesID).Error; err != nil {
			response.WriteJson(w, http.StatusNotFound, response.GeneralError(fmt.Errorf("data not found")))
			return
		}

		if err := db.GetDB().Delete(&tes).Error; err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(fmt.Errorf("failed delete data: %v", err)))
			return
		}
		response.WriteJson(w, http.StatusOK, map[string]interface{}{"success": "OK", "message": "success deleted"})
	}
}
