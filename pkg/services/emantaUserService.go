package services

import (
	"errors"

	"github.com/Cristian-RM/system-api/pkg/models"
)

type EmantaUserServices struct {
	User models.EmantaUser
}

func (*EmantaUserServices) CreateNewUser(NewUser *models.EmantaUser) (Error error) {
	//Existe el usuario
	Error = nil
	insertedUser, _ := models.GetEmantaUserByUserName(NewUser.UserName)
	if insertedUser != nil {
		//user exists
		Error = errors.New("user already exists")
	}

	return Error
}

func (*EmantaUserServices) ValidateData(NewUser *models.EmantaUser) {

}
