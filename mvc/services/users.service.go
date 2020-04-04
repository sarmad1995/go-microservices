package services

import (
	"github.com/sarmad1995/mvc/domain"
	utils "github.com/sarmad1995/mvc/utils/error"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
