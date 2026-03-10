package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func setupIntegrationRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	godotenv.Load()
	InitDB()

	r := gin.Default()

	r.GET("/fall", getFall)
	r.GET("/fall/:id", getFallDetail)
	r.GET("/subjekt", getSubjekt)
	r.GET("/subjekt/:id", getSubjektDetail)
	r.GET("/log", getLog)
	r.GET("/log/:id", getLogDetail)

	return r
}

func TestIntegration_GetFall_List(t *testing.T) {
	r := setupIntegrationRouter()

	req, _ := http.NewRequest("GET", "/fall?limit=2&offset=0", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var result []Fall
	err := json.Unmarshal(w.Body.Bytes(), &result)
	assert.NoError(t, err)
	assert.True(t, len(result) <= 2)
}

func TestIntegration_GetFall_Detail(t *testing.T) {
	r := setupIntegrationRouter()

	req, _ := http.NewRequest("GET", "/fall/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Contains(t, []int{200, 404}, w.Code)
}

func TestIntegration_GetSubjekt_List(t *testing.T) {
	r := setupIntegrationRouter()

	req, _ := http.NewRequest("GET", "/subjekt?limit=3&name=A", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestIntegration_GetSubjekt_Detail(t *testing.T) {
	r := setupIntegrationRouter()

	req, _ := http.NewRequest("GET", "/subjekt/100", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Contains(t, []int{200, 404}, w.Code)
}

func TestIntegration_GetLog_List(t *testing.T) {
	r := setupIntegrationRouter()

	req, _ := http.NewRequest("GET", "/log?limit=5", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestIntegration_GetLog_Detail(t *testing.T) {
	r := setupIntegrationRouter()

	req, _ := http.NewRequest("GET", "/log/500", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Contains(t, []int{200, 404}, w.Code)
}
