package customers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sigolang-sv/app/models"
	"sigolang-sv/helpers"
	"sigolang-sv/schemas"
)

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	decode := json.NewDecoder(r.Body)
	decode.DisallowUnknownFields()

	var customer schemas.Customer
	err := decode.Decode(&customer)
	if err != nil {
		helpers.JSONErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("Error Decoding: %s", err.Error()))
		return
	}
	if customer.Name == "" {
		helpers.JSONErrorResponse(w, http.StatusBadRequest, "Name is required")
		return
	}
	if customer.Age <= 0 {
		helpers.JSONErrorResponse(w, http.StatusBadRequest, "Age is required")
		return
	}

	var profile schemas.Profile

	Risk := 55 - customer.Age
	switch {
	case Risk >= 30:
		profile.StockPercent = 72.5
		profile.BondPercent = 21.5
		profile.MMPercent = 100 - (profile.StockPercent + profile.BondPercent)

	case Risk >= 20:
		profile.StockPercent = 54.5
		profile.BondPercent = 25.5
		profile.MMPercent = 100 - (profile.StockPercent + profile.BondPercent)

	case Risk < 20:
		profile.StockPercent = 34.5
		profile.BondPercent = 45.5
		profile.MMPercent = 100 - (profile.StockPercent + profile.BondPercent)
	}

	customer.Profile = &profile
	result := models.CreateCustomer(customer)
	if result.Error != nil {
		helpers.JSONErrorResponse(w, http.StatusInternalServerError, "Internal Server Error: can't create customer")
		return
	}

	helpers.JSONSuccessResponse(w, customer, "Customer Created")

}