package controllers

import (
	"encoding/json"
	"net/http"
	"superapps/models"
	service "superapps/services"
	helper "superapps/helpers"
)

func VerifyOtp(w http.ResponseWriter, r *http.Request) {
	
	data := &models.User{}

	err := json.NewDecoder(r.Body).Decode(data)

	if err != nil {
		helper.Logger("error", "In Server: " + err.Error())
		helper.Response(w, 400, true, "Internal server error ("+err.Error()+")", map[string]interface{}{})
		return
	}

	email 	 := data.Email
	otp 	 := data.Otp

	if  email == "" {
		helper.Logger("error", "In Server: email field is required")
		helper.Response(w, 400, true, "email field is required", map[string]interface{}{})
		return
	} 

	if  otp == "" {
		helper.Logger("error", "In Server: otp field is required")
		helper.Response(w, 400, true, "otp field is required", map[string]interface{}{})
		return
	} 

	result, err := service.VerifyOtp(data)

	if err != nil {
		helper.Response(w, 400, true, err.Error(), map[string]interface{}{})
		return
	}

	helper.Logger("info", "Verify account success")
	helper.Response(w, http.StatusOK, false, "Successfully", result)
}

func ResendOtp(w http.ResponseWriter, r *http.Request) {
	
	data := &models.User{}

	err := json.NewDecoder(r.Body).Decode(data)

	if err != nil {
		helper.Logger("error", "In Server: " + err.Error())
		helper.Response(w, 400, true, "Internal server error ("+err.Error()+")", map[string]interface{}{})
		return
	}

	email := data.Email

	if  email == "" {
		helper.Logger("error", "In Server: email field is required")
		helper.Response(w, 400, true, "email field is required", map[string]interface{}{})
		return
	} 

	result, err := service.ResendOtp(data)

	if err != nil {
		helper.Response(w, 400, true, err.Error(), map[string]interface{}{})
		return
	}

	helper.Logger("info", "Resend otp success")
	helper.Response(w, http.StatusOK, false, "Successfully", result)
}

