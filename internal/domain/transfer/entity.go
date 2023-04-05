package transfer

type TransferEntity struct {
	ID           uint
	PayerName    string
	ReceiveName  string
	Amount       uint64
	TransferDate string
}

type CreateTransfer struct {
	PayerName   string
	ReceiveName string
	Amount      uint64
}
