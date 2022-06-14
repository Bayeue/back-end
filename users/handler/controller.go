package handler_users

import (
	"net/http"
	"ppob/app/middlewares"
	"ppob/helper/encryption"
	domain_users "ppob/users/domain"
	"ppob/users/handler/request"
	"ppob/users/handler/response"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type UsersHandler struct {
	usecase    domain_users.Service
	validation *validator.Validate
}

func NewUsersHandler(uc domain_users.Service) UsersHandler {
	return UsersHandler{
		usecase:    uc,
		validation: validator.New(),
	}
}

func (uh *UsersHandler) Authorization(ctx echo.Context) error {
	req := request.RequestJSONLogin{}
	ctx.Bind(&req)
	if err := uh.validation.Struct(req); err != nil {
		stringerr := []string{}
		for _, errval := range err.(validator.ValidationErrors) {
			stringerr = append(stringerr, errval.Field()+" is not "+errval.Tag())
		}
		return ctx.JSON(http.StatusBadRequest, stringerr)
	}

	res, err := uh.usecase.Login(req.Email, req.Password)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusBadRequest,
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

// implementation register users
func (uh *UsersHandler) Register(ctx echo.Context) error {
	req := request.RequestJSONUser{}
	ctx.Bind(&req)
	if err := uh.validation.Struct(req); err != nil {
		stringerr := []string{}
		for _, errval := range err.(validator.ValidationErrors) {
			stringerr = append(stringerr, errval.Field()+" is not "+errval.Tag())
		}
		return ctx.JSON(http.StatusBadRequest, stringerr)
	}

	encrypt, err := encryption.HashPassword(req.Password)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "internal server error",
			"rescode": http.StatusInternalServerError,
		})
	}

	req.Password = encrypt
	data, err := uh.usecase.Register(request.ToDomainUser(req))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request",
			"rescode": http.StatusBadRequest,
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success register",
		"rescode": http.StatusOK,
		"data": map[string]interface{}{
			"token": data,
		},
	})
}

// implementation store/save pin data users
func (uh *UsersHandler) InsertAccount(ctx echo.Context) error {
	req := request.RequestJSONAccount{}
	ctx.Bind(&req)
	if err := uh.validation.Struct(req); err != nil {
		stringerr := []string{}
		for _, errval := range err.(validator.ValidationErrors) {
			stringerr = append(stringerr, errval.Field()+" is not "+errval.Tag())
		}
		return ctx.JSON(http.StatusBadRequest, stringerr)
	}
	// Encryption
	encrypt, err := encryption.HashPassword(req.Pin)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "internal server error",
			"rescode": http.StatusInternalServerError,
		})
	}
	req.Pin = encrypt
	// get data from jwt
	dataUser := middlewares.GetUser(ctx)
	req.Phone = dataUser.Phone

	res, err := uh.usecase.InsertAccount(request.ToDomainAccount(req))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "internal server error",
			"rescode": http.StatusInternalServerError,
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
		"data":    res,
	})
}

// implementation get all data
func (uh *UsersHandler) GetUsers(ctx echo.Context) error {
	sliceResponse := []interface{}{}
	res, err := uh.usecase.GetUsers()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "internal server error",
			"rescode": http.StatusInternalServerError,
		})
	}
	for _, value := range res {
		sliceResponse = append(sliceResponse, response.FromDomainUsers(value))
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get users",
		"rescode": http.StatusOK,
		"result":  sliceResponse,
	})
}

// Implementation get user by id for admin (web)
func (uh *UsersHandler) GetUserForAdmin(ctx echo.Context) error {
	phone := ctx.Param("phone")

	user, err := uh.usecase.GetUserPhone(phone)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request",
			"rescode": http.StatusBadRequest,
		})
	}
	account, err := uh.usecase.GetUserAccount(phone)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request",
			"rescode": http.StatusBadRequest,
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get users",
		"rescode": http.StatusOK,
		"result": map[string]interface{}{
			"user":    response.FromDomainUsers(user),
			"account": response.FromDomainAccount(account),
		},
	})
}

// Implementation get user by id for admin (web)
func (uh *UsersHandler) GetUserForCustomer(ctx echo.Context) error {
	jwtClaims := middlewares.GetUser(ctx)
	phone := jwtClaims.Phone

	user, err := uh.usecase.GetUserPhone(phone)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request",
			"rescode": http.StatusBadRequest,
		})
	}
	account, err := uh.usecase.GetUserAccount(phone)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request",
			"rescode": http.StatusBadRequest,
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get customer",
		"rescode": http.StatusOK,
		"result": map[string]interface{}{
			"user":    response.FromDomainUsers(user),
			"account": response.FromDomainAccount(account),
		},
	})
}

// implementation update user data
func (uh *UsersHandler) UpdateProfile(ctx echo.Context) error {
	req := request.RequestJSONUser{}
	ctx.Bind(&req)
	if err := uh.validation.Struct(req); err != nil {
		stringerr := []string{}
		for _, errval := range err.(validator.ValidationErrors) {
			stringerr = append(stringerr, errval.Field()+" is not "+errval.Tag())
		}
		return ctx.JSON(http.StatusBadRequest, stringerr)
	}
	encrypt, err := encryption.HashPassword(req.Password)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "internal server error",
			"rescode": http.StatusInternalServerError,
		})
	}

	req.Password = encrypt
	user := middlewares.GetUser(ctx)
	err = uh.usecase.EditUser(user.Phone, request.ToDomainUser(req))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request",
			"rescode": http.StatusBadRequest,
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update customer profile",
		"rescode": http.StatusOK,
	})
}

// implementation Destroy user for Admin
func (uh *UsersHandler) DestroyUserForAdmin(ctx echo.Context) error {
	panic("")
}

func (uh *UsersHandler) UserRole(phone string) (string, bool) {
	var role string
	var status bool
	user, err := uh.usecase.GetUserPhone(phone)
	if err == nil {
		role = user.Role
		status = user.Status
	}
	return role, status
}
