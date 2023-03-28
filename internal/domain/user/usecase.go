package user

type UserUsecase struct {
	userRepo IUserRepository
}

func NewUserUsecase(ur IUserRepository) *UserUsecase {
	return &UserUsecase{
		userRepo: ur,
	}
}
