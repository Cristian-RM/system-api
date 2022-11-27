package models

import (
	"github.com/Cristian-RM/system-api/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type EmantaUser struct {
	gorm.Model
	Name      string `gorm:"" json:"name"`
	UserName  string `json:"username"`
	Password  string `json:"password"`
	LastLogin string `json:"lastlogin"`
}

func Init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&EmantaUser{})
}

func (b *EmantaUser) CreateEmantaUser() *EmantaUser {
	db.NewRecord(b)
	db.Create(&b)
	return b
}
func GetAllEmantaUsers() []EmantaUser {
	var EmantaUsers []EmantaUser
	db.Find(&EmantaUsers)
	return EmantaUsers
}
func GetEmantaUserById(ID int64) (*EmantaUser, *gorm.DB) {
	var EmantaUser EmantaUser
	db := db.Where("userName=?", ID).Find(&EmantaUser)
	return &EmantaUser, db
}

func GetEmantaUserByUserName(UserName string) (*EmantaUser, *gorm.DB) {
	var EmantaUser EmantaUser
	db := db.Where("user_name=?", UserName).Find(&EmantaUser)
	return &EmantaUser, db
}

func DeleteEmantaUser(username string) EmantaUser {
	var EmantaUser EmantaUser
	db.Where("user_name=?", username).Delete(EmantaUser)
	return EmantaUser
}

func HardDeleteEmantaUser(username string) EmantaUser {
	var EmantaUser EmantaUser
	db.Where("user_name=?", username).Delete(EmantaUser)
	db.Unscoped().Delete(&EmantaUser, "user_name=?", username)
	return EmantaUser
}
