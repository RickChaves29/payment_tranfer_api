package transfer

type ITransferRepository interface {
	Save(userID uint64, data CreateTransfer) error
	FindAllTransfers(userId uint64) ([]TransferEntity, error)
}
