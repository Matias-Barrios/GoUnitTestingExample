package handlers

import (
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetRouterOk(t *testing.T) {
	var subject IHandlerProvider = HandlerProvider{}
	router := subject.GetRouter()
	assert.NotNil(t, router)

	req := httptest.NewRequest("GET", "/health", nil)

	m := &mux.RouteMatch{}
	require.True(t, router.Match(req, m), "no match")

	v1 := reflect.ValueOf(m.Handler)
	require.Equal(t, v1.Pointer(), reflect.ValueOf(Health).Pointer(), "wrong handler")
}
