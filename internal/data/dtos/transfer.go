package dtos

type TransferDTO struct {
	ID           uint   `json:"id"`
	PayerName    string `json:"payer_name"`
	ReceiveName  string `json:"receive_name"`
	Amount       uint64 `json:"amount"`
	TransferDate string `json:"transfer_data"`
}
