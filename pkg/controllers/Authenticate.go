package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Cristian-RM/system-api/pkg/models"
	"github.com/Cristian-RM/system-api/pkg/utils"
	"github.com/gorilla/mux"
)

var newEmantaUser models.EmantaUser

func GetBookById(w http.ResponseWriter, r *http.Request) {
	//todo validar que solo la misma persona o un admin pueda traearse esta info, excepto el password y esas cosas.
	newEmantaUser = models.EmantaUser{}
	vars := mux.Vars(r)
	username := vars["UserName"]
	//Por ahora retornaremos con el username
	// ID, err := strconv.ParseInt(username, 0, 0)
	// if err != nil {
	// 	fmt.Println("Error while parsing")
	// }
	bookDetails, _ := models.GetEmantaUserByUserName(username)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Conten-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
func CreateBook(w http.ResponseWriter, r *http.Request) {
	createBook := &models.EmantaUser{}
	utils.ParseBody(r, createBook)
	b := createBook.CreateEmantaUser()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsin")
	}
	book := models.DeleteEmantaUser(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateEmantaUser = &models.EmantaUser{}
	utils.ParseBody(r, updateEmantaUser)
	vars := mux.Vars(r)
	usernameID := vars["UserNameID"]
	ID, err := strconv.ParseInt(usernameID, 0, 0)
	if err != nil {
		fmt.Println("error while parsin")
	}
	bookDetails, db := models.GetEmantaUserById(ID)
	if updateEmantaUser.Name != "" {
		bookDetails.Name = updateEmantaUser.Name
	}
	// if updateEmantaUser.password != "" {
	// 	bookDetails.password = updateEmantaUser.Author
	// }

	// if updateEmantaUser.Publication != "" {
	// 	bookDetails.Publication = updateEmantaUser.Publication
	// }

	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
