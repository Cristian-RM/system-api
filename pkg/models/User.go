package models

import (
	"github.com/Cristian-RM/system-api/pkg/config"
	"github.com/jinzhu/gorm"
)

type User struct {
	userName string
	passwork string
}

var db *gorm.DB

type EmantaUser struct {
	gorm.Model
	Name      string `gorm:""json:"name"`
	userName  string `json:"username"`
	password  string `json:"password"`
	lastLogin string `json:"lastlogin"`
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
	db := db.Where("userName=?", UserName).Find(&EmantaUser)
	return &EmantaUser, db
}

func DeleteEmantaUser(Id int64) EmantaUser {
	var EmantaUser EmantaUser
	db.Where("ID=?", Id).Delete(EmantaUser)
	return EmantaUser
}
