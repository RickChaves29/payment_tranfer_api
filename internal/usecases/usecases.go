package usecases

import "errors"

type User struct {
	ID      string
	Balance int64
}
type IUserRepository interface {
	FindUserByID(id string) (*User, error)
	UpdateBalance(id string, amount int64) error
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

func (uc *UserUsecase) TransferPayment(payer, receive string, amount int64) error {
	userPayer, err := uc.repo.FindUserByID(payer)
	if err != nil {
		return err
	}
	if userPayer.Balance < amount {
		return errors.New("insufficient funds")
	}

	userReceive, err := uc.repo.FindUserByID(receive)
	if err != nil {
		return err
	}
	err = uc.repo.UpdateBalance(userPayer.ID, userPayer.Balance-amount)
	if err != nil {
		return err
	}
	err = uc.repo.UpdateBalance(userReceive.ID, userReceive.Balance+amount)
	if err != nil {
		return err
	}
	return nil
}
