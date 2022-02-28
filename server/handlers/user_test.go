package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/devbyP/untitled/models"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func init() {
	godotenv.Overload("test.env")
}

func TestSignup(t *testing.T) {
	db, err := models.ConnectDB()
	if err != nil {
		t.Error("error connecting to test db")
	}
	models.SetDB(db)
	if err := models.PingTest(); err != nil {
		t.Error("ping database fail")
	}
	defer models.GetDB().Close()
	defer db.Close()
	testUserJson := `{"username": "testUser1", "password": "1234", "firstName": "test", "lastName": "user"}`
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/user", strings.NewReader(testUserJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err = SignUpHandler(c)
	if err != nil {
		t.Errorf("signup error: %s", err)
	}
	if rec.Code != http.StatusCreated {
		t.Errorf("fail signup expected status code %d, instead %d", http.StatusCreated, rec.Code)
	}
	db.QueryRow("DELETE FROM users WHERE username = 'testUser1'")
}

func TestFetchUser(t *testing.T) {
	db, err := models.ConnectDB()
	if err != nil {
		t.Error("error connection to test db")
	}
	models.SetDB(db)
	defer db.Close()
	defer models.GetDB().Close()
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/user", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err = FetchAllUsersHandler(c)
	if err != nil {
		t.Errorf("error response %s", err)
	}
	if rec.Code != http.StatusOK {
		t.Errorf("error with status %d", rec.Code)
	}
}
