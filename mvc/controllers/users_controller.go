package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	utils "github.com/sarmad1995/mvc/utils/error"

	"github.com/sarmad1995/mvc/services"
)

// GetUser is used to  get a user
func GetUser(res http.ResponseWriter, req *http.Request) {
	userId, err := (strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64))
	if err != nil {
		// return the error
		apiError := &utils.ApplicationError{
			Message:    "User id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonValue, _ := json.Marshal(apiError)

		res.WriteHeader(apiError.StatusCode)
		res.Write(jsonValue)
		return
	}
	user, apiError := services.GetUser(userId)
	if apiError != nil {
		// return not found user
		res.WriteHeader(apiError.StatusCode)
		jsonValue, _ := json.Marshal(apiError)
		print(jsonValue)
		res.Write(jsonValue)
		return
	}
	jsonValue, _ := json.Marshal(user)
	res.Write(jsonValue)
	// return user
}
