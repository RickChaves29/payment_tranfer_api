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

func (uc *TranferUsecase) ListAllTransfers(userId uint64) ([]TransferEntity, error) {
	transfers, err := uc.transferRepo.FindAllTransfers(userId)
	if err != nil {
		return nil, err
	}
	return transfers, nil
}
