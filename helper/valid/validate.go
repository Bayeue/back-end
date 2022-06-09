package valid

import (
	"fmt"
	"net/http"
	"ppob/app/middlewares"
	handler_users "ppob/users/handler"

	"github.com/labstack/echo/v4"
)

func RoleValidation(role string, userHandler handler_users.UsersHandler) echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := middlewares.GetUser(c)
			fmt.Println("claim : ", claims)
			userRole, status := userHandler.UserRole(claims.Phone)
			fmt.Println("userRole : ", userRole)

			if userRole == role && status == true {
				return hf(c)
			} else {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{
					"message": "account not active, please contact admin",
				})
			}
		}
	}
}