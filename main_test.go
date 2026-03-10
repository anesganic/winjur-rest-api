package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func setupTestRouter() (*gin.Engine, sqlmock.Sqlmock) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	mockDB, mock, _ := sqlmock.New()

	db = sqlx.NewDb(mockDB, "sqlserver")

	return r, mock
}

func TestGetFall_Success(t *testing.T) {
	r, mock := setupTestRouter()
	r.GET("/fall", getFall)

	mock.ExpectQuery("SELECT \\* FROM dbo.FALL").WillReturnRows(sqlmock.NewRows([]string{"FallNo", "Name"}).AddRow(1, "Testprojekt"))

	req, _ := http.NewRequest("GET", "/fall?status=aktiv", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestGetFallDetail_InvalidID(t *testing.T) {
	r, _ := setupTestRouter()
	r.GET("/fall/:id", getFallDetail)

	req, _ := http.NewRequest("GET", "/fall/abc", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.Contains(t, w.Body.String(), "ID muss eine Zahl sein")
}

func TestGetSubjekt_Success(t *testing.T) {
	r, mock := setupTestRouter()
	r.GET("/subjekt", getSubjekt)

	mock.ExpectQuery("SELECT \\* FROM dbo.SUBJEKT").WillReturnRows(sqlmock.NewRows([]string{"AdrNr", "Name1"}).AddRow(100, "ITST GmbH"))

	req, _ := http.NewRequest("GET", "/subjekt?name=ITST", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestGetSubjektDetail_InvalidID(t *testing.T) {
	r, _ := setupTestRouter()
	r.GET("/subjekt/:id", getSubjektDetail)

	req, _ := http.NewRequest("GET", "/subjekt/xyz", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}

func TestGetLog_Success(t *testing.T) {
	r, mock := setupTestRouter()
	r.GET("/log", getLog)

	mock.ExpectQuery("SELECT \\* FROM dbo.LOG").WillReturnRows(sqlmock.NewRows([]string{"LogNo", "Fall"}).AddRow(500, 1))

	req, _ := http.NewRequest("GET", "/log?type=1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestGetLogDetail_InvalidID(t *testing.T) {
	r, _ := setupTestRouter()
	r.GET("/log/:id", getLogDetail)

	req, _ := http.NewRequest("GET", "/log/ungueltig", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}
