package controllers

import (
	"net/http"
	"superapps/models"
	helper "superapps/helpers"
)

func All(w http.ResponseWriter, r *http.Request) {

	data := &models.News{}

	result, err := data.GetNews()

	if err != nil {
		helper.Response(w, 400, true, err.Error(), map[string]interface{}{})
		return
	}

	helper.Response(w, http.StatusOK, false, "Successfully", result["news"])
}
