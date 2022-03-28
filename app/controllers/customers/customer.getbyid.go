package customers

import (
	"fmt"
	"net/http"
	"sigolang-sv/app/models"
	"sigolang-sv/helpers"

	"github.com/gorilla/mux"
)

func GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	CustomerID := vars["id"]

	fmt.Println(vars)

	customer, result := models.GetCustomerByID(CustomerID)
	if result.RowsAffected == 1 {
		helpers.JSONSuccessResponse(w, customer, "Customer Retrieved")
	} else {
		helpers.JSONErrorResponse(w, http.StatusNotFound, "Customer Not Found")
	}
}