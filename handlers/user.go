package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Matias-Barrios/GoUnitTestingExample/models"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := _userprovider.GetUsers()
	if err != nil {
		ErrorHandler(w, r, err, 500)
		return
	}
	bs, _ := json.MarshalIndent(users, "", "  ")
	w.Write(bs)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user models.User
	decoder.Decode(&user)
	err := _userprovider.CreateUser(user)
	if err != nil {
		ErrorHandler(w, r, err, 500)
		return
	}

	w.Write([]byte("ok"))

}
