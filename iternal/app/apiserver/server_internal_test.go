package apiserver

import (
	"http-rest-api/iternal/app/store/teststore"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser_HandleUsersCreate(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/users", nil)
	s := newServer(teststore.New())

	s.ServeHTTP(rec, req)

	assert.Equal(t, rec.Code, http.StatusOK)
}
