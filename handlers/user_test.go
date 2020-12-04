package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/Matias-Barrios/GoUnitTestingExample/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type _userProviderMock struct {
	mock.Mock
}

func (m *_userProviderMock) GetUsers() ([]models.User, error) {
	args := m.Called()
	return args.Get(0).([]models.User), args.Error(1)
}

func (m *_userProviderMock) CreateUser(u models.User) error {
	args := m.Called()
	return args.Get(0).(error)
}

func TestGetUsersOk(t *testing.T) {
	_userProviderMock := new(_userProviderMock)

	_userProviderMock.On("GetUsers").Return([]models.User{
		models.User{
			Name: "test",
			Age:  22,
		},
	}, nil)

	_userprovider = _userProviderMock

	req := httptest.NewRequest("GET", "localhost:8000/users", nil)
	w := httptest.NewRecorder()
	GetUsers(w, req)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var users []models.User
	err := json.Unmarshal(body, &users)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "text/plain; charset=utf-8", resp.Header.Get("Content-Type"))
	assert.Nil(t, err)
	assert.Equal(t, 1, len(users))
	assert.Equal(t, "test", users[0].Name)

	_userProviderMock.AssertExpectations(t)
}
