package domain_users

type Service interface {
	Login(email, password string) (string, error)
}

type Repository interface {
	CheckEmailPassword(email, password string) (Users, error)
}
