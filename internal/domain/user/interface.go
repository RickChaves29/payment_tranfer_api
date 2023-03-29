package user

type IUserRepository interface {
	Save(data CreateUserEntity) error
	FindUserByEmail(email string) (*UserEntity, error)
}
