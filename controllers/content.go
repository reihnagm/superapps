package controllers

import (
	"net/http"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	models "superapps/models"
	service "superapps/services"
	helper "superapps/helpers"
)

func All(w http.ResponseWriter, r *http.Request) {

	page	:= r.URL.Query().Get("page")
	limit 	:= r.URL.Query().Get("limit")
	search  := r.URL.Query().Get("search")
	appName := r.URL.Query().Get("app_name")

	result, err := service.GetContent(search, page, limit, appName)

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
		result["data"],
	)
}

func CreateMediaContent(w http.ResponseWriter ,r*http.Request) {

	data := &models.ContentMedia{}

	errCreateMedia := json.NewDecoder(r.Body).Decode(data)

	if errCreateMedia != nil {
		helper.Logger("error", "In Server: " + errCreateMedia.Error())
		return
	}

	result, err := service.CreateMediaContent(data)

	if err != nil {
		helper.Response(w, 400, true, err.Error(), map[string]interface{}{})
		return
	}

	helper.Logger("info", "Create media success")
	helper.Response(w, http.StatusOK, false, "Successfully", result)
}

func CreateContent(w http.ResponseWriter, r *http.Request) {

	data := &models.Content{}

	errCreateContent := json.NewDecoder(r.Body).Decode(data)

	if errCreateContent != nil {
		helper.Logger("error", "In Server: " + errCreateContent.Error())
		return
	}

	tokenHeader := r.Header.Get("Authorization")

	token := helper.DecodeJwt(tokenHeader)

	claims, _ := token.Claims.(jwt.MapClaims)
		
	UserId, _ := claims["id"].(string)

	appName := r.Header.Get("APP_NAME")
	id	  	:= data.Uid
	title 	:= data.Title
	desc  	:= data.Description

	data.AppName = appName
	data.UserId = UserId

	if appName == "" {
		helper.Logger("error", "In Server: APP_NAME headers is required")
		helper.Response(w, 400, true, "APP_NAME headers is required", map[string]interface{}{})
		return
	}

	if id == "" {
		helper.Logger("error", "In Server: id is required")
		helper.Response(w, 400, true, "id is required", map[string]interface{}{})
		return
	}

	if  title == "" {
		helper.Logger("error", "In Server: title field is required")
		helper.Response(w, 400, true, "title field is required", map[string]interface{}{})
		return
	} 

	if  desc == "" {
		helper.Logger("error", "In Server: description field is required")
		helper.Response(w, 400, true, "description field is required", map[string]interface{}{})
		return
	} 

	result, err := service.CreateContent(data)

	if err != nil {
		helper.Response(w, 400, true, err.Error(), map[string]interface{}{})
		return
	}

	helper.Logger("info", "Create Media success")
	helper.Response(w, http.StatusOK, false, "Successfully", result)
}
