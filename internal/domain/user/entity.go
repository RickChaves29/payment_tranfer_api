package user

type UserEntity struct {
	ID       uint64
	Name     string
	Email    string
	Password string
	Balance  uint64
}

type CreateUserEntity struct {
	Name     string
	Email    string
	Password string
}
