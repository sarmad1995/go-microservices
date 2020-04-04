package domain

import (
	"net/http"

	utils "github.com/sarmad1995/mvc/utils/error"
)

var (
	users = map[int64]*User{
		123: {Id: 123, Email: "123@gmail.com", FirstName: "123", LastName: "456"},
	}
)

func GetUser(userId int64) (*User, *utils.ApplicationError) {
	user := users[userId]
	if user == nil {
		return nil, &utils.ApplicationError{
			Message:    "User not found",
			StatusCode: http.StatusNotFound,
			Code:       "not_found",
		}
	}
	return user, nil
}
