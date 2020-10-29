package handlers

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthOk(t *testing.T) {
	req := httptest.NewRequest("GET", "localhost:8000/health", nil)
	w := httptest.NewRecorder()
	Health(w, req)

	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)

	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "text/plain; charset=utf-8", resp.Header.Get("Content-Type"))
	assert.Equal(t, "Alive", string(body))
}
