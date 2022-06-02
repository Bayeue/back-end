package routes

import (
	handler_users "ppob/users/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware middleware.JWTConfig
	UserHandler   handler_users.UsersHandler
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {

	// product public
	e.POST("/login", cl.UserHandler.Authorization)

}
