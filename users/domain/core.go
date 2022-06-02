package domain_users

type Users struct {
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
