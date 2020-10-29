package handlers

import (
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
	return r
}

// Services

var _userprovider services.IUserProvider = services.UserProvider{}
