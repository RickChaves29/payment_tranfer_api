package user

type IUserRepository interface {
	FindUserByID(id uint64) (*User, error)
	FindUserByEmail(email string) (*User, error)
	UpdateBalance(id uint64, amount uint64) error
}
