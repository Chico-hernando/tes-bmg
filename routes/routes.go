package routes

import (
	"bmg/constants"
	"bmg/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e:= echo.New()

	eAuth := e.Group("")
	eAuth.Use(middleware.JWT([]byte(constants.JWT_SECRET)))

	e.POST("/register", controllers.CreateUserController)
	e.POST("/login",controllers.LoginUserController)
	eAuth.PUT("/user/:id",controllers.UpdateUserController)
	eAuth.POST("/reff",controllers.ReffCodeController)
	e.GET("/user",controllers.GetUserByName)

	e.GET("/hero",controllers.HeroController)

	return e
}