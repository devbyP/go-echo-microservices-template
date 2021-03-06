package handlers

import (
	"net/http"
	"strings"

	"github.com/devbyP/untitled/pkg/models"
	"github.com/labstack/echo/v4"
)

func FetchAllUsersHandler(c echo.Context) error {
	users, err := models.FetchAllUser()
	if err != nil {
		return fetchError(c, err)
	}
	return c.JSON(http.StatusOK, users)
}

func FetchUserHandler(c echo.Context) error {
	id := c.Param("id")
	if !models.ValidateUUID(id) {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid id"})
	}
	user, err := models.FetchUser(id)
	if err != nil {
		return fetchError(c, err)
	}
	user.ID = "abc"
	if user.ID == "" {
		return c.JSON(http.StatusNotFound, "no user found")
	}
	return c.JSON(http.StatusOK, user)
}

func generateFakeHash(pass string) string {
	return pass + "fakeHash"
}

func compareFakeHash(pass, hash string) bool {
	return strings.HasPrefix(hash, pass)
}

func SignUpHandler(c echo.Context) error {
	reqBody := struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Username  string `json:"username"`
		Password  string `json:"password"`
	}{}
	if err := c.Bind(&reqBody); err != nil {
		return bindDataError(c)
	}
	newUser := models.NewUser(reqBody.FirstName, reqBody.LastName, reqBody.Username)
	hash := generateFakeHash(reqBody.Password)
	if err := newUser.SignUp(hash); err != nil {
		return internalError(c, "cannot sign up", err)
	}
	return c.JSON(http.StatusCreated, map[string]string{"message": "ok"})
}

func signInCheck(username, password string) bool {
	user, err := models.FetchUserByUsername(username)
	if err != nil {
		return false
	}
	if user.ID == "" {
		return false
	}
	hash, err := models.GetUserHash(user.ID)
	if !compareFakeHash(password, hash) {
		return false
	}
	return err != nil
}

func SignInHandler(c echo.Context) error {
	reqBody := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}
	if err := c.Bind(&reqBody); err != nil {
		return bindDataError(c)
	}
	if !signInCheck(reqBody.Username, reqBody.Password) {
		return signInError(c)
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "sign-in success"})
}

func HelloHandler(c echo.Context) error {
	return c.String(http.StatusOK, "hello world")
}
