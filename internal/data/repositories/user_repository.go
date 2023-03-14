package repositories

import (
	"database/sql"

	"github.com/RickChaves29/payment_tranfer_api/internal/usecases"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindUserByID(id string) (*usecases.User, error) {
	user := usecases.User{}
	defer r.db.Close()
	row := r.db.QueryRow(`SELECT * FROM users WHERE id = ?`, id)
	err := row.Scan(&user.ID, &user.Balance)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *userRepository) UpdateBalance(id string, amount int64) error {
	return nil
}
