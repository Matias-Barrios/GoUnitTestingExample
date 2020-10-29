package handlers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRouterOk(t *testing.T) {
	var subject IHandlerProvider = HandlerProvider{}
	router := subject.GetRouter()
	assert.NotNil(t, router)
}
