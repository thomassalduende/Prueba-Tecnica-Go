package routes

import (
	"mutant-checker/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/mutant", controller.MutantHandler).Methods("POST")
	r.HandleFunc("/stats", controller.StatsHandler).Methods("GET")
	http.Handle("/", r)
}
