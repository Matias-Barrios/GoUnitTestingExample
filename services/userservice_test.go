package services

import (
	"database/sql"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Matias-Barrios/GoUnitTestingExample/models"
	"github.com/Matias-Barrios/GoUnitTestingExample/repositories"
	"github.com/stretchr/testify/assert"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestGetUsersOk(t *testing.T) {
	db, mock := NewMock()
	mock.ExpectQuery("SELECT firstname, lastname, email, age FROM Users;").
		WillReturnRows(sqlmock.NewRows([]string{"firstname", "lastname", "email", "age"}).
			AddRow("pepe", "guerra", "pepe@gmail.com", 34))
	subject := UserProvider{
		DatabaseProvider: repositories.NewMockDBProvider(db, nil),
	}
	resp, err := subject.GetUsers()

	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 1, len(resp))
}

func TestCreateUserOk(t *testing.T) {
	db, mock := NewMock()
	mock.ExpectExec(`
	INSERT
	 INTO Users \(firstname, lastname, email, age, created_on\) 
	 VALUES\(\$1,\$2,\$3,\$4,\$5\)`).
		WithArgs(
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	subject := UserProvider{
		DatabaseProvider: repositories.NewMockDBProvider(db, nil),
	}
	err := subject.CreateUser(models.User{
		Age:      34,
		Name:     "Matias",
		Email:    "matias@pepe.com",
		Lastname: "Barrios",
	})

	assert.Nil(t, err)
}
