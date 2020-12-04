package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Matias-Barrios/GoUnitTestingExample/configproviders"
	"github.com/Matias-Barrios/GoUnitTestingExample/repositories"
	"github.com/Matias-Barrios/GoUnitTestingExample/services"
	"github.com/gorilla/mux"
)

type IHandlerProvider interface {
	GetRouter() *mux.Router
}

type HandlerProvider struct{}

func (h HandlerProvider) GetRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/health", Health).Methods("GET")
	r.HandleFunc("/users", GetUsers).Methods("GET")
	r.HandleFunc("/users", CreateUser).Methods("POST")
	return r
}

// Services

var _userprovider services.IUserProvider

func ErrorHandler(w http.ResponseWriter, r *http.Request, err error, status int) {
	w.WriteHeader(status)
	response := struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}{
		Status:  status,
		Message: err.Error(),
	}
	bytes, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		w.Header().Add("Content-Type", "application/text")
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(bytes)
}

func init() {
	_userprovider = &services.UserProvider{
		Name:             "Pepe",
		DatabaseProvider: repositories.NewDBProvider(configproviders.EnvironmentVariableProvider{}),
	}
}
