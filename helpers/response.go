package helper

import (
	"encoding/json"
	"net/http"
	"strings"
	"errors"
	entities "superapps/entities"
)

func MessageSuccess(status int, err bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "error": err, "message": message}
}

func MessageError(status int, err bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "error": err, "message": message}
}

func Response(w http.ResponseWriter, status int, err bool, message string, data any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	jsonEncoder := json.NewEncoder(w)
	errs := jsonEncoder.Encode(&entities.Response{status, err, message, data})
	if errs != nil {
		Logger("error", errs.Error())
	}
}

func ResponseWithPagination(w http.ResponseWriter, status int, err bool, message string, total any, perPage any, prevPage any, nextPage any, currentPage any, nextUrl any, prevUrl, data any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	jsonEncoder := json.NewEncoder(w)
	errs := jsonEncoder.Encode(&entities.ResponseWithPagination{status, err, message, total, perPage, prevPage, nextPage, currentPage, nextUrl, prevUrl, data})
	if errs != nil {
		Logger("error", errs.Error())
	}
}

func FormatError(err string) error {

	if strings.Contains(err, "email") {
		Logger("error", "In Server: Email Already Taken")
		return errors.New("Email Already Taken")
	}
	
	if strings.Contains(err, "phone") {
		Logger("error", "In Server: Phone Already Taken")
		return errors.New("Phone Already Taken")
	}

	if strings.Contains(err, "hashedPassword") {
		Logger("error", "In Server: Incorrect Password")
		return errors.New("Incorrect Password")
	}

	return errors.New(err)
}
