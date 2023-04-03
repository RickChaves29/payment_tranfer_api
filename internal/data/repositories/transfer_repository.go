package repositories

import (
	"database/sql"

	"github.com/RickChaves29/payment_tranfer_api/internal/domain/transfer"
)

type transferRepository struct {
	db *sql.DB
}

func NewTransferRepository(db *sql.DB) *transferRepository {
	return &transferRepository{db: db}
}

func (r *transferRepository) Save(userID uint64, data transfer.CreateTransfer) error {
	_, err := r.db.Exec(`INSERT INTO transfers (payer, receive, amount, user_id) VALUES ($1, $2, $3, $4)`, data.PayerName, data.ReceiveName, data.Amount, userID)
	if err != nil {
		return err
	}
	return nil
}
func (r *transferRepository) FindAllTransfers(userId uint64) ([]transfer.TransferEntity, error) {
	var transfers []transfer.TransferEntity
	rows, err := r.db.Query(`
	SELECT transfers.id, transfers.payer, transfers.receive,transfers.amount, transfers.transfer_date 
	FROM transfers
	INNER JOIN users 
	ON users.id = $1
	`, userId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var transfer transfer.TransferEntity
		rows.Scan(&transfer.ID, &transfer.PayerName, &transfer.ReceiveName, &transfer.Amount, &transfer.TransferDate)

		transfers = append(transfers, transfer)
	}
	return transfers, nil
}
