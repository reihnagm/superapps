package controllers

import (
	"net/http"
	"superapps/models"
	helper "superapps/helpers"
)

func All(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")

	data := &models.News{}

	result, err := data.GetNews(page, limit)

	if err != nil {
		helper.Response(w, 400, true, err.Error(), map[string]interface{}{})
		return
	}

	helper.ResponseWithPagination(w, http.StatusOK, false, "Successfully", 
		result["total"], 
		result["per_page"], 
		result["prev_page"], 
		result["next_page"], 
		result["current_page"], 
		result["next_url"],
		result["prev_url"],
		result["news"],
	)
}
