package transfer

type TranferUsecase struct {
	transferRepo ITransferRepository
}

func NewTaskferUsecase(tr ITransferRepository) *TranferUsecase {
	return &TranferUsecase{
		transferRepo: tr,
	}
}

func (uc *TranferUsecase) CreateNewTransfer(userID uint64, data CreateTransfer) error {
	err := uc.transferRepo.Save(userID, data)
	if err != nil {
		return err
	}
	return nil
}
