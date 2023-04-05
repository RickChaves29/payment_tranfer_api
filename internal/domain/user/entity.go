package user

type UserEntity struct {
	ID       uint64
	Name     string
	Email    string
	Password string
	Balance  uint64
}

type CreateUserEntity struct {
	Name     string
	Email    string
	Password string
}

type UserTransferPaymentEntity struct {
	Payer   string
	Receive string
	Amount  uint64
}
