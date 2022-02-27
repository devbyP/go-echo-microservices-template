package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/devbyP/untitled/models"
	"github.com/labstack/echo/v4"
)

func internalError(c echo.Context, mess string) error {
	return c.JSON(
		http.StatusInternalServerError,
		fmt.Sprintf(`{"message": "internal server error: %s"}`, mess))
}

func fetchError(c echo.Context) error {
	return internalError(c, "cannot fetch user")
}

func jsonParseError(c echo.Context) error {
	return internalError(c, "cannot parse json")
}

func bindDataError(c echo.Context) error {
	return c.JSON(http.StatusBadRequest, `{"message": "error bind data from body"}`)
}

func signInError(c echo.Context) error {
	return c.JSON(http.StatusUnauthorized, `{"message": "incorrect username or password"}`)
}

func FetchAllUsersHandler(c echo.Context) error {
	users, err := models.FetchAllUser()
	if err != nil {
		return fetchError(c)
	}
	j, err := json.Marshal(users)
	if err != nil {
		return jsonParseError(c)
	}
	return c.JSON(http.StatusOK, j)
}

func FetchUserHandler(c echo.Context) error {
	id := c.Param("id")
	user, err := models.FetchUser(id)
	if err != nil {
		return fetchError(c)
	}
	j, err := json.Marshal(user)
	if err != nil {
		return jsonParseError(c)
	}
	return c.JSON(http.StatusOK, j)
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
		return internalError(c, "cannot sign up")
	}
	return c.JSON(http.StatusOK, `{"message": "ok"}`)
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
	return c.JSON(http.StatusOK, `{"message": "sign-in success"}`)
}

func HelloHandler(c echo.Context) error {
	return c.String(http.StatusOK, "hello world")
}
