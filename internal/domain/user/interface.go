package user

type IUserRepository interface {
	FindUserByID(id string) (*User, error)
	UpdateBalance(id string, amount uint64) error
}
