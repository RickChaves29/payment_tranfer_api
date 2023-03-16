package dtos

type TransferDataDTO struct {
	Payer   string `json:"payer"`
	Receive string `json:"receive"`
	Amount  int64  `json:"amount"`
}
