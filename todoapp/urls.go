package todoapp

import (
	"github.com/gorilla/mux"
)

func SetupRoutes(routes *mux.Router) {
	routes.HandleFunc("/", IndexView)
	routes.HandleFunc("/create", CreateEntry)
	routes.HandleFunc("/{id}/delete", DeleteEntry).Methods("POST")
}
