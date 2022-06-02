package mysql_users

import (
	"fmt"
	domain_users "ppob/users/domain"

	"gorm.io/gorm"
)

type UsersRepo struct {
	DB *gorm.DB
}

func NewUsersRepo(db *gorm.DB) domain_users.Repository {
	return UsersRepo{
		DB: db,
	}
}

// CheckEmailPassword implements domain_users.Repository
func (ur UsersRepo) CheckEmailPassword(email string, password string) (domain_users.Users, error) {
	var rec Users

	err := ur.DB.Where("email = ? && password = ?", email, password).First(&rec).Error
	if err != nil {
		return domain_users.Users{}, err
	}
	fmt.Println("get user : ", rec)
	return toDomain(rec), nil
}
