package repositories

import (
	"database/sql"

	"github.com/RickChaves29/payment_tranfer_api/internal/domain/user"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Save(data user.CreateUserEntity) error {
	_, err := r.db.Exec(`INSERT INTO users (user_name, email, user_password) VALUES ($1, $2, $3)`, data.Name, data.Email, data.Password)
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) FindUserByEmail(email string) (*user.UserEntity, error) {
	user := user.UserEntity{}
	row := r.db.QueryRow(`SELECT * FROM users WHERE email = $1`, email)
	err := row.Scan(&user.ID, &user.Balance)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *userRepository) UpdateBalance(id uint64, amount uint64) error {
	_, err := r.db.Exec(`UPDATE users SET balance = $1 WHERE id = $2`, amount, id)
	if err != nil {
		return err
	}
	return nil
}
