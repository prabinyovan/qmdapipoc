package auth

import (
	"net/http"
	contracts "qmdapipoc/contracts/auth"
	"qmdapipoc/domain/users"
	auth_service "qmdapipoc/services/auth"
	"qmdapipoc/utils/errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

const (
	SecretKey = "58zr6NZdudRmebl3fFYaY584wbub2cuJ"
)

type jwtCustomClaims struct {
	UserGlobalKey string `json:"userglobalkey"`
	jwt.StandardClaims
}

func Login(c echo.Context) (err error) {
	var user users.User

	if err = c.Bind(&user); err != nil {
		error := errors.NewCustomError("PasswordEncryptionFailed", "005")
		c.JSON(http.StatusBadRequest, error)
		return
	}

	result, getRrr := auth_service.GetUser(user)

	if getRrr != nil {
		error := errors.NewCustomError("UserNotFound", getRrr.ErrorCode)
		c.JSON(http.StatusNotFound, error)
		return
	}

	// Set custom claims
	// claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
	// 	Issuer:    strconv.Itoa(int(result.ID)),
	// 	ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
	// })

	// Set custom claims
	claims := &jwtCustomClaims{
		result.UserGlobalKey,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 365).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))

	if err != nil {
		return err
	}

	auth_out_contract := contracts.AuthOutContract{
		Token:     t,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	c.JSON(http.StatusOK, auth_out_contract)
	return
}
