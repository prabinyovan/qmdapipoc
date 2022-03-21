package users

import (
	"log"
	"net/http"
	"qmdapipoc/domain/users"
	user_service "qmdapipoc/services/users"

	"github.com/labstack/echo/v4"
)

func Register(c echo.Context) (err error) {
	log.Println("Register")
	var user users.User

	if err = c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	_, saveErr := user_service.CreateUser(user)

	if saveErr != nil {
		c.JSON(http.StatusInternalServerError, saveErr)
		return
	}

	c.JSON(http.StatusOK, "")
	return
}

func GetUsers(c echo.Context) (err error) {

	users, getErr := user_service.GetUserCollection()

	if getErr != nil {
		c.JSON(http.StatusInternalServerError, getErr)
		return
	}

	c.JSON(http.StatusOK, users)
	return
}
