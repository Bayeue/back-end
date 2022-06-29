package helper_xendit

import (
	"encoding/json"
	"errors"
	domain_transaction "ppob/transaction/domain"

	"github.com/labstack/echo/v4"
)

func GetCallback(domain domain_transaction.Callback_Invoice, ctx echo.Context) (interface{}, error) {
	callback_otp := "BjVVRO8eKgceve38jmqm6twtK9YLjtAfk7CbJLxfiToTilHX"
	if ctx.Request().Header.Get("x-callback-token") == callback_otp {
		ctx.JSON(401, map[string]interface{}{
			"message": "unauthorized",
		})
	}

	decoder := json.NewDecoder(ctx.Request().Body)
	callbackData := domain

	err := decoder.Decode(&callbackData)
	if err != nil {
		return "empty", errors.New("internal status error")
	}

	defer ctx.Request().Body.Close()

	callback, _ := json.Marshal(callbackData)

	ctx.Response().Header().Set("Content-Type", "application/json")

	ctx.Response().WriteHeader(200)
	return callback, nil
}
