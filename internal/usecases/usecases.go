package usecases

type User struct {
	ID      string
	Balance int64
}
type IUserRepository interface {
	FindUserByID(id string) (*User, error)
}
type UserUsecase struct {
	repo IUserRepository
}

func NewUserUsecase(r IUserRepository) *UserUsecase {
	return &UserUsecase{
		repo: r,
	}
}

func (uc *UserUsecase) GetUser(id string) (*User, error) {
	user, err := uc.repo.FindUserByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
