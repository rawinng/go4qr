package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrepareServer(t *testing.T) {
	router := prepareServer()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/qr", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.NotEmpty(t, w.Body.String())
}
