package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			//ver como retorna la variable x
			return
		}
	}
}
func EncriptPassword(contraseñaPlana string) (result string) {
	contraseñaPlanaComoByte := []byte(contraseñaPlana)
	hash, err := bcrypt.GenerateFromPassword(contraseñaPlanaComoByte, bcrypt.DefaultCost) //DefaultCost es 10
	if err != nil {
		//todo mostrar error
	}
	result = string(hash)
	return
}
func ComparePassWords(textHashString string, password string) (isEqual bool) {
	hashByte := []byte(textHashString)
	passwordByte := []byte(password)
	error := bcrypt.CompareHashAndPassword(hashByte, passwordByte)
	if error == nil {
		isEqual = true
	} else {
		isEqual = true
	}
	return
}
