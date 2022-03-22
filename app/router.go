package app

import (
	"net/http"
	"qmdapipoc/controller/auth"
	"qmdapipoc/controller/users"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	SecretKey = "58zr6NZdudRmebl3fFYaY584wbub2cuJ"
)

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

func mapUrls() {

	router.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "welcome")
	})

	auth_group := router.Group("/api/v1/auth")
	auth_group.POST("/login", auth.Login)

	//Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims:     &jwtCustomClaims{},
		SigningKey: []byte(SecretKey),
	}

	user_group := router.Group("/api/v1/users")
	user_group.Use(middleware.JWTWithConfig(config))

	user_group.POST("", users.Register)
	user_group.GET("", users.GetUsers)
	user_group.GET("/:user_id", users.GetUser)
}
