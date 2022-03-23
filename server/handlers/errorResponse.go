package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// base function for internal server error code: 500.
func internalError(c echo.Context, mess string, err error) error {
	statusCode := http.StatusInternalServerError
	res := newErrorResponse(statusCode, "internal server error: "+mess, err)
	return c.JSON(
		statusCode,
		res,
	)
}

//      specific for fetching to database error.
func fetchError(c echo.Context, err error) error {
	return internalError(c, "cannot fetch user", err)
}

/* func jsonParseError(c echo.Context) error {
        return internalError(c, "cannot parse json")
} */

// specific for data binding in http req body context.
func bindDataError(c echo.Context) error {
	code := http.StatusBadRequest
	err := fmt.Errorf("error bind data from body")
	res := newErrorResponse(code, "cannot read data from request body", err)
	return c.JSON(code, res)
}

func signInError(c echo.Context) error {
	return c.JSON(http.StatusUnauthorized, map[string]string{"message": "incorrect username or password"})
}
