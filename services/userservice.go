package services

import (
	"fmt"
	"time"

	"github.com/Matias-Barrios/GoUnitTestingExample/models"
	"github.com/Matias-Barrios/GoUnitTestingExample/repositories"
)

type IUserProvider interface {
	GetUsers() ([]models.User, error)
	CreateUser(u models.User) error
}

type UserProvider struct {
	Name             string
	DatabaseProvider repositories.IDBProvider
}

func (up UserProvider) GetUsers() ([]models.User, error) {
	var users = make([]models.User, 0, 10)
	rows, err := up.DatabaseProvider.Query("SELECT firstname, lastname, email, age FROM Users;")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var u models.User = models.User{}
		err := rows.Scan(&u.Name, &u.Lastname, &u.Email, &u.Age)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (up UserProvider) CreateUser(user models.User) error {
	sqlr, err := up.DatabaseProvider.Exec(`
	INSERT
	 INTO Users (firstname, lastname, email, age, created_on) 
	 VALUES($1,$2,$3,$4,$5)`, user.Name, user.Lastname, user.Email, user.Age, time.Now().Format(time.RFC3339))
	if err != nil {
		return err
	}
	rows, err := sqlr.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("could not insert rows")
	}
	if err != nil {
		return err
	}
	return nil
}
