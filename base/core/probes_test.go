package core

import (
	"app/base/database"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLiveness(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	InitRouter(Liveness).ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestReadiness(t *testing.T) {
	SetupTest(t)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	InitRouter(Readiness).ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestReadinessFail(t *testing.T) {
	SetupTest(t)

	sqlDB, _ := database.DB.DB()
	assert.Nil(t, sqlDB.Close())
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	InitRouter(Readiness).ServeHTTP(w, req)
	assert.Equal(t, http.StatusServiceUnavailable, w.Code)
}
