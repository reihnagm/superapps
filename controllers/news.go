package controllers

import (
	"net/http"
	"encoding/json"
	"os"
	"strings"
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

	result, err := service.GetNews(search, page, limit, appName)

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

	data := &models.NewsForm{}

	err := json.NewDecoder(r.Body).Decode(data)

	tokenHeader := r.Header.Get("Authorization")

	appName := r.Header.Get("APP_NAME")

	splitted := strings.Split(tokenHeader, " ")

	tokenPart := splitted[1]

	token, err := jwt.Parse(tokenPart, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		helper.Logger("error", "In Server: " + err.Error())
		return
	}

	if err != nil {
		helper.Logger("error", "In Server: " + err.Error())
		helper.Response(w, 400, true, "Internal server error ("+err.Error()+")", map[string]interface{}{})
		return
	}

	if token.Valid {
		claims, ok := token.Claims.(jwt.MapClaims)
		
		if ok {
			UserId, _ := claims["id"].(string)
			
			data.ApplicationName = appName
			data.UserId = UserId

			title := data.Title
			desc  := data.Description

			if appName == "" {
				helper.Logger("error", "In Server: APP_NAME headers is required")
				helper.Response(w, 400, true, "APP_NAME headers is required", map[string]interface{}{})
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

			result, err := service.CreateNews(data)

			if err != nil {
				helper.Response(w, 400, true, err.Error(), map[string]interface{}{})
				return
			}

			helper.Logger("info", "Create News success")
			helper.Response(w, http.StatusOK, false, "Successfully", result)
		} else {
			helper.Logger("error", "In Server: Invalid token claims.")
			helper.Response(w, 400, true, "Internal server error ("+err.Error()+")", map[string]interface{}{})
		}
	}


}
