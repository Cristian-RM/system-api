package testing

import (
	"testing"

	"github.com/Cristian-RM/system-api/pkg/models"
)

func TestInsertUser(t *testing.T) {
	models.Init()

	emanta_user := models.EmantaUser{Name: "Cristian", UserName: "crismors@gmail.com", Password: "12345", LastLogin: "12/12/1997"}

	insertedUser, _ := models.GetEmantaUserByUserName("crismors@gmail.com")
	if insertedUser != nil {
		models.HardDeleteEmantaUser(emanta_user.UserName)

	}
	emanta_user.CreateEmantaUser()
	insertedUser, _ = models.GetEmantaUserByUserName("crismors@gmail.com")

	if insertedUser.Name != emanta_user.Name {
		t.Error("User cannot be inserted and readed from db.")
	}

}
