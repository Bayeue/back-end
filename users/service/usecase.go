package service_users

import (
	"errors"
	"fmt"
	"ppob/app/middlewares"
	domain_users "ppob/users/domain"
)

type UsersService struct {
	Repository domain_users.Repository
	jwtauth    *middlewares.ConfigJwt
}

func NewUsersService(repo domain_users.Repository, jwt *middlewares.ConfigJwt) domain_users.Service {
	return UsersService{
		Repository: repo,
		jwtauth:    jwt,
	}
}

// Login implements domain_users.Service
func (us UsersService) Login(email string, password string) (string, error) {
	data, err := us.Repository.CheckEmailPassword(email, password)
	fmt.Println("data login : ", data)
	if err != nil {
		return "", errors.New("failed to generate token, user not found")
	}
	token, err := us.jwtauth.GenerateToken(data.ID, data.Status)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}
