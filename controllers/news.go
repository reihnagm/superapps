package controllers

import (
	"net/http"
	service "superapps/services"
	helper "superapps/helpers"
)

func All(w http.ResponseWriter, r *http.Request) {

	page	:= r.URL.Query().Get("page")
	limit 	:= r.URL.Query().Get("limit")
	search  := r.URL.Query().Get("search")

	result, err := service.GetNews(search, page, limit)

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

func CreateNews(w http.ResponseWriter, r *http.Request) {

	// err := json.NewDecoder(r.Body).Decode(data)

	// if err != nil {
	// 	helper.Logger("error", "In Server: " + err.Error())
	// 	helper.Response(w, 400, true, "Internal server error ("+err.Error()+")", map[string]interface{}{})
	// 	return
	// }

}
