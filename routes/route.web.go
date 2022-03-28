package routes

import (
	"sigolang-sv/app/controllers"
	"sigolang-sv/app/controllers/customers"

	"github.com/gorilla/mux"
)

func Route() *mux.Router {
	Route := mux.NewRouter().StrictSlash(true)
	Route.HandleFunc("/", controllers.Index)
	Route.HandleFunc("/customer", customers.CreateCustomer).Methods("OPTIONS","POST")
	Route.HandleFunc("/customers", customers.GetAllCustomer).Methods("OPTIONS", "GET")
	Route.HandleFunc("/customer/{id}", customers.GetCustomerByID).Methods("OPTIONS", "GET")

	return Route
}