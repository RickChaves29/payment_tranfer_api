package user

import (
	"errors"
)

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

func (uc *UserUsecase) TransferPayment(data UserTransferPaymentEntity) error {
	payer, err := uc.userRepo.FindUserByEmail(data.Payer)
	if err != nil {
		return err
	}
	receive, err := uc.userRepo.FindUserByEmail(data.Receive)
	if err != nil {
		return err
	}
	if payer.Balance < data.Amount {
		return errors.New("insufficient founds")
	}
	err = uc.userRepo.UpdateBalance(payer.ID, payer.Balance-data.Amount)
	if err != nil {
		return err
	}
	err = uc.userRepo.UpdateBalance(receive.ID, receive.Balance+data.Amount)
	if err != nil {
		return err
	}
	return nil
}
