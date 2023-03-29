package user

type UserUsecase struct {
	userRepo IUserRepository
}

func NewUserUsecase(ur IUserRepository) *UserUsecase {
	return &UserUsecase{
		userRepo: ur,
	}
}

func (uc *UserUsecase) CreateNewUser(data CreateUserEntity) error {
	err := uc.userRepo.Save(data)
	if err != nil {
		return err
	}
	return nil
}

func (uc *UserUsecase) GetUserByEmail(email string) (*UserEntity, error) {
	user, err := uc.userRepo.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}
