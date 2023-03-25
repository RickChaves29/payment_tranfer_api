package usecases

type User struct {
	ID      string
	Balance int64
}

type Transfer struct {
	Payer   string
	Receive string
	Amount  int64
}
