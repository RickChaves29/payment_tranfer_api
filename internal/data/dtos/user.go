package dtos

type UserDTO struct {
	ID      uint64 `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Balance uint64 `json:"balance"`
}

type CreateUserDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserTransferPaymentDTO struct {
	Payer   string `json:"payer"`
	Receive string `json:"receive"`
	Amount  uint64 `json:"amount"`
}
