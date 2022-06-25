package domain_users_test

import (
	"errors"
	"os"
	"ppob/app/middlewares"
	domain_users "ppob/users/domain"
	usersMocks "ppob/users/domain/mocks"
	service_users "ppob/users/service"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	userService       domain_users.Service
	userDomainUser    domain_users.Users
	userDomainAccount domain_users.Account
	userRepo          usersMocks.Repository
)

func TestMain(m *testing.M) {
	userService = service_users.NewUsersService(&userRepo, &middlewares.ConfigJwt{})
	userDomainUser = domain_users.Users{
		ID:        1,
		Name:      "Muhammad Fadillah Abdul Aziz",
		Email:     "aziz@gmail.com",
		Password:  "12345",
		Phone:     "0895631948686",
		Image:     "aziz.jpg",
		Status:    true,
		Role:      "customer",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
	userDomainAccount = domain_users.Account{
		ID:    1,
		Phone: "0895631948686",
		Saldo: 0,
		Pin:   "12345",
	}
	os.Exit(m.Run())
}

func TestGetUsers(t *testing.T) {
	t.Run("get users", func(t *testing.T) {
		userRepo.On("GetAllUser").Return([]domain_users.Users{userDomainUser, {ID: 2, Name: "Kiana", Phone: "081320503262"}}, nil).Once()
		res, err := userService.GetUsers()

		assert.NoError(t, err)
		assert.Equal(t, userDomainUser.Name, res[0].Name)
	})
	t.Run("data emtpy", func(t *testing.T) {
		userRepo.On("GetAllUser").Return([]domain_users.Users{}, errors.New("data empty")).Once()
		res, err := userService.GetUsers()

		assert.Error(t, err)
		assert.Equal(t, []domain_users.Users{}, res)
	})
}

func TestGetUserPhone(t *testing.T) {
	t.Run("get user by phone", func(t *testing.T) {
		userRepo.On("GetByPhone", mock.Anything).Return(userDomainUser, nil).Once()
		res, err := userService.GetUserPhone(userDomainUser.Phone)

		assert.NoError(t, err)
		assert.Equal(t, userDomainUser.Phone, res.Phone)
	})
	t.Run("get user by phone", func(t *testing.T) {
		userRepo.On("GetByPhone", mock.Anything).Return(domain_users.Users{}, errors.New("user not found")).Once()
		res, err := userService.GetUserPhone(userDomainUser.Phone)

		assert.Error(t, err)
		assert.Equal(t, "", res.Phone)
	})
}
func TestEditUser(t *testing.T) {
	t.Run("edit user by phone", func(t *testing.T) {
		userRepo.On("Update", mock.Anything, mock.Anything).Return(nil).Once()

		err := userService.EditUser(userDomainUser.Phone, userDomainUser)

		assert.NoError(t, err)
		assert.Equal(t, nil, err)
	})
	t.Run("failed edit user by phone", func(t *testing.T) {
		userRepo.On("Update", mock.Anything, mock.Anything).Return(errors.New("row affacted")).Once()

		err := userService.EditUser(userDomainUser.Phone, userDomainUser)

		assert.Error(t, err)
		assert.Equal(t, err, err)
	})
}

func TestLogin(t *testing.T) {
	t.Run("login success", func(t *testing.T) {
		userRepo.On("CheckEmailPassword", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(userDomainUser, nil).Once()
		userRepo.On("GenerateToken", mock.AnythingOfType("int"), mock.AnythingOfType("bool")).Return("eeyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwicGhvbmUiOiIwODk1NjMxOTQ4Njg2Iiwic3RhdHVzIjp0cnVlfQ._V-6g3RUQ5Z3_Vv2R9adrVeHCfsCjIb8R6b8G6EYs3A", nil).Once()
		res, err := userService.Login("Muhammad Fadillah Abdul Aziz", "12345")

		assert.NoError(t, err)
		assert.Equal(t, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwicGhvbmUiOiIwODk1NjMxOTQ4Njg2Iiwic3RhdHVzIjp0cnVlfQ._V-6g3RUQ5Z3_Vv2R9adrVeHCfsCjIb8R6b8G6EYs3A", res)
	})
	t.Run("login Failed Bad Request", func(t *testing.T) {
		userRepo.On("CheckEmailPassword", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(domain_users.Users{}, errors.New("data not found")).Once()
		userRepo.On("GenerateToken", mock.AnythingOfType("int"), mock.AnythingOfType("bool")).Return("", errors.New("failed to generate token")).Once()
		res, err := userService.Login("", "12345")

		assert.Error(t, err)
		assert.Equal(t, "", res)
	})
}

func TestRegister(t *testing.T) {
	t.Run("register Success", func(t *testing.T) {
		userRepo.On("Store", mock.Anything).Return(userDomainUser.Phone, nil).Once()
		userRepo.On("GetByPhone", mock.Anything).Return(userDomainUser, nil)
		res, err := userService.Register(userDomainUser)

		assert.NoError(t, err)
		assert.Equal(t, userDomainUser.Phone, res.Phone)
	})
	t.Run("register Data Not Found", func(t *testing.T) {
		userRepo.On("Store", mock.Anything).Return("", errors.New("bad request")).Once()
		userRepo.On("GetByPhone", mock.Anything).Return(userDomainUser, nil)
		res, err := userService.Register(domain_users.Users{})

		assert.Error(t, err)
		assert.Equal(t, domain_users.Users{}, res)
	})
}

func TestInsertAccount(t *testing.T) {
	t.Run("make account user", func(t *testing.T) {
		userRepo.On("StoreAccount", mock.Anything).Return(userDomainAccount, nil).Once()
		res, err := userService.InsertAccount(userDomainAccount)
		assert.NoError(t, err)
		assert.Equal(t, userDomainAccount.Phone, res.Phone)
	})
	t.Run("failed make account user", func(t *testing.T) {
		userRepo.On("StoreAccount", mock.Anything).Return(domain_users.Account{}, errors.New("bad request")).Once()
		res, err := userService.InsertAccount(userDomainAccount)
		assert.Error(t, err)
		assert.Equal(t, domain_users.Account{}, res)
	})
}
func TestGetUserAccount(t *testing.T) {
	t.Run("get account user", func(t *testing.T) {
		userRepo.On("GetUserAccount", mock.Anything).Return(userDomainAccount, nil).Once()
		res := userService.GetUserAccount(userDomainAccount.Phone)

		assert.Equal(t, userDomainAccount.Phone, res.Phone)
	})
	t.Run("failed get account user", func(t *testing.T) {
		userRepo.On("GetUserAccount", mock.Anything).Return(domain_users.Account{}, errors.New("account not found")).Once()
		res := userService.GetUserAccount(userDomainAccount.Phone)

		assert.Equal(t, domain_users.Account{}, res)
	})
}

// func TestAddUserVerif(t *testing.T) {
// 	panic("")
// }

// func TestVerif(t *testing.T) {
// 	panic("")
// }
