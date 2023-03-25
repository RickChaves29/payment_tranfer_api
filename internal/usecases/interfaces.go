package usecases

type IUserRepository interface {
	FindUserByID(id string) (*User, error)
	UpdateBalance(id string, amount int64) error
}

type ITranferRepository interface {
	FindAllTransfersByUserID(id string) ([]Transfer, error)
}
