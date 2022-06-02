package handler_users

import (
	"fmt"
	"net/http"
	domain_users "ppob/users/domain"
	request_users "ppob/users/handler/request"

	"github.com/labstack/echo/v4"
)

type UsersHandler struct {
	usecase domain_users.Service
}

func NewUsersHandler(uc domain_users.Service) UsersHandler {
	return UsersHandler{
		usecase: uc,
	}
}

func (uh *UsersHandler) Authorization(ctx echo.Context) error {
	rec := request_users.RequestJSON{}

	if err := ctx.Bind(&rec); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request",
			"rescode": http.StatusBadRequest,
		})
	}
	fmt.Println("bind : ", rec)
	res, err := uh.usecase.Login(rec.Email, rec.Password)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "unauthorized",
			"rescode": http.StatusUnauthorized,
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "user success login",
		"rescode": http.StatusOK,
		"data": map[string]interface{}{
			"token": res,
		},
	})
}
