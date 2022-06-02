package mysql_users

import (
	domain_users "ppob/users/domain"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID         int
	Name       string
	Slug       string
	Email      string
	Password   string
	Phone      string
	Status     bool
	Address_Id int
	Role       string
}

func toDomain(rec Users) domain_users.Users {
	return domain_users.Users{
		Name:       rec.Name,
		Slug:       rec.Slug,
		Email:      rec.Email,
		Password:   rec.Password,
		Phone:      rec.Phone,
		Status:     rec.Status,
		Address_Id: rec.Address_Id,
		Role:       rec.Role,
	}
}
