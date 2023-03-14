package integration_test

import (
	"os"
	"testing"

	"github.com/RickChaves29/payment_tranfer_api/internal/data"
	"github.com/RickChaves29/payment_tranfer_api/internal/data/repositories"
	"github.com/RickChaves29/payment_tranfer_api/internal/usecases"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	file, _ := os.Create("test.db")
	defer file.Close()
	db, _ := data.ConnectionDB("sqlite3", "test.db")
	defer db.Close()
	db.Exec(`CREATE TABLE users (
		id VARCHAR(200) UNIQUE,
		balance UNSIGNED BIG INT
	)`)
	db.Exec(`INSERT INTO users (id, balance) VALUES ('1234', 1000)`)
	db.Exec(`INSERT INTO users (id, balance) VALUES ('4321', 2000)`)
}
func TestFindUserById(t *testing.T) {
	expect := usecases.User{
		ID:      "1234",
		Balance: 1000,
	}
	db, _ := data.ConnectionDB("sqlite3", "test.db")
	r := repositories.NewUserRepository(db)
	user, _ := r.FindUserByID("1234")
	if user.Balance != expect.Balance {
		t.Errorf("have %v want %v", user.Balance, expect.Balance)
	}
	t.Cleanup(func() {
		db.Exec(`DROP TABLE users`)
	})
}
