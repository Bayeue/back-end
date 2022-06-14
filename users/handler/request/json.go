package request

import (
	domain_users "ppob/users/domain"
	"time"
)

type RequestJSONUser struct {
	Name     string    `json:"name" form:"name" validate:"required"`
	Email    string    `json:"email" form:"email" validate:"required,email"`
	Password string    `json:"password" form:"password" validate:"required"`
	Phone    string    `json:"phone" form:"phone" validate:"required"`
	DOB      time.Time `json:"dob" form:"dob"`
	Image    string    `json:"img" form:"img"`
}

type RequestJSONAccount struct {
	Phone string
	Saldo int
	Pin   string `json:"pin" form:"pin"`
}

type RequestJSONLogin struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
}

func ToDomainUser(req RequestJSONUser) domain_users.Users {
	return domain_users.Users{
		Name:     req.Name,
		Slug:     "example-slug",
		DOB:      req.DOB,
		Email:    req.Email,
		Password: req.Password,
		Phone:    req.Phone,
		Image:    req.Image,
		Status:   true,
		Role:     "customer",
	}
}
func ToDomainLogin(req RequestJSONLogin) domain_users.Users {
	return domain_users.Users{
		Email:    req.Email,
		Password: req.Password,
	}
}

func ToDomainAccount(req RequestJSONAccount) domain_users.Account {
	return domain_users.Account{
		Phone: req.Phone,
		Saldo: 0,
		Pin:   req.Pin,
	}
}
