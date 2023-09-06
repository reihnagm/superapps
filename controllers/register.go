package controllers

import (
	"encoding/json"
	"net/http"
	"superapps/models"
	helper "superapps/helpers"
)

func Register(w http.ResponseWriter, r *http.Request) {

	data := &models.User{}

	err := json.NewDecoder(r.Body).Decode(data)

	if err != nil {
		helper.Logger("error", "In Server: " + err.Error())
		helper.Response(w, 400, true, "Internal server error ("+err.Error()+")", map[string]interface{}{})
		return
	}

	email 	 := data.Email
	password := data.Password
	phone 	 := data.Phone

	if  email == "" {
		helper.Logger("error", "In Server: email field is required")
		helper.Response(w, 400, true, "email field is required", map[string]interface{}{})
		return
	} 

	validateEmail := helper.IsValidEmail(email)

	if validateEmail != true {
		helper.Logger("error", "In Server: E-mail address is invalid")
		helper.Response(w, 400, true, "E-mail address is invalid", map[string]interface{}{})
		return
	} 

	if  phone == "" {
		helper.Logger("error", "In Server: phone field is required")
		helper.Response(w, 400, true, "phone field is required", map[string]interface{}{})
		return
	} 

	if  password == "" {
		helper.Logger("error", "In Server: password field is required")
		helper.Response(w, 400, true, "password field is required", map[string]interface{}{})
		return
	} 

	helper.Logger("info", "Register Success")
	helper.Response(w, http.StatusOK, false, "Successfully", map[string]interface{}{})
}