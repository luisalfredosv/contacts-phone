package routes

import (
	"github.com/luisalfredosv/golang-gorilla/controllers"
	"github.com/gorilla/mux"
)

func SetContactsRoutes(r *mux.Router){
	subrouter := r.PathPrefix("/api").Subrouter()

	subrouter.HandleFunc("/contacts/{id}", controllers.GetContact).Methods("GET")
	subrouter.HandleFunc("/contacts/", controllers.GetContacts).Methods("GET")
	subrouter.HandleFunc("/contacts/", controllers.StoreContact).Methods("POST")
	subrouter.HandleFunc("/contacts/", controllers.UpdateContact).Methods("PUT")
	subrouter.HandleFunc("/contacts/{id}", controllers.GetContact).Methods("DELETE")

}