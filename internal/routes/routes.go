package routes

import (
	"github.com/N0TTEAM/begos/internal/http/handlers/tes2"
	"github.com/gorilla/mux"
)

func RegisterRoute(router *mux.Router) {
	router.HandleFunc("/tes", tes2.CreateTes()).Methods("POST")
	router.HandleFunc("/tes/delete", tes2.DeleteTesById()).Methods("DELETE")
}
