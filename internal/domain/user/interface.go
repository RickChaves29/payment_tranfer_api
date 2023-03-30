package user

type IUserRepository interface {
	Save(data CreateUserEntity) error
	FindUserByEmail(email string) (*UserEntity, error)
	UpdateBalance(id uint64, amount uint64) error
}
