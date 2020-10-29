package services

import (
	"github.com/Matias-Barrios/GoUnitTestingExample/models"
)

type IUserProvider interface {
	GetUsers() []models.User
}

type UserProvider struct{}

func (up UserProvider) GetUsers() []models.User {
	return []models.User{
		models.User{
			Name: "Pepe",
			Age:  22,
		},
		models.User{
			Name: "Maria",
			Age:  34,
		},
	}
}
