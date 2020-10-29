package handlers

import (
	"encoding/json"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	bs, _ := json.Marshal(_userprovider.GetUsers())
	w.Write(bs)
}
