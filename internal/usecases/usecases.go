package usecases

import "errors"

type UserUsecase struct {
	userRepo     IUserRepository
	transferRepo ITranferRepository
}

func NewUserUsecase(ur IUserRepository, tr ITranferRepository) *UserUsecase {
	return &UserUsecase{
		userRepo:     ur,
		transferRepo: tr,
	}
}

func (uc *UserUsecase) GetUser(id string) (*User, error) {
	user, err := uc.userRepo.FindUserByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uc *UserUsecase) TransferPayment(payer, receive string, amount int64) error {
	userPayer, err := uc.userRepo.FindUserByID(payer)
	if err != nil {
		return err
	}
	if userPayer.Balance < amount {
		return errors.New("insufficient funds")
	}

	userReceive, err := uc.userRepo.FindUserByID(receive)
	if err != nil {
		return err
	}
	err = uc.userRepo.UpdateBalance(userPayer.ID, userPayer.Balance-amount)
	if err != nil {
		return err
	}
	err = uc.userRepo.UpdateBalance(userReceive.ID, userReceive.Balance+amount)
	if err != nil {
		return err
	}
	return nil
}

func (uc *UserUsecase) GetAllTransfers(id string) ([]Transfer, error) {
	transfers, err := uc.transferRepo.FindAllTransfersByUserID(id)
	if err != nil {
		return nil, err
	}
	return transfers, nil
}
