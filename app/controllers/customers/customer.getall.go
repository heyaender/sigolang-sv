package customers

import (
	"fmt"
	"net/http"
	"sigolang-sv/app/models"
	"sigolang-sv/helpers"
	"strconv"
)

func GetAllCustomer(w http.ResponseWriter, r *http.Request) {
	limit := 10
	offset := 0

	limitQuery := r.URL.Query().Get("limit")
	offsetQuery := r.URL.Query().Get("offset")

	var err error
	if limitQuery != "" {
		limit, err = strconv.Atoi(limitQuery)
		if err != nil {
			helpers.JSONErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("Error Parsing: %s", err.Error()))
			return
		}
	}
	if offsetQuery != "" {
		offset, err = strconv.Atoi(offsetQuery)
		if err != nil {
			helpers.JSONErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("Error Parsing: %s", err.Error()))
			return
		}
	}

	customers := models.GetAllCustomer(limit, offset)
	helpers.JSONSuccessResponse(w, customers, "Customers Retrieved")
}