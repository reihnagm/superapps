package controllers

import (
	"net/http"
	helper "superapps/helpers"
)

func NewsAll(w http.ResponseWriter, r *http.Request) {

	// resp := helper.MessageSuccess(http.StatusOK, false, "Successfully")
	helper.Response(w, http.StatusOK, false, "Successfully", map[string]interface{}{})
}
