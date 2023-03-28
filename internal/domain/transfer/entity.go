package transfer

type Transfer struct {
	ID           uint
	PayerName    string
	ReceiveName  string
	Amount       uint64
	TransferDate string
}
