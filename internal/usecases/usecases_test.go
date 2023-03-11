package usecases_test

import (
	"testing"

	"github.com/RickChaves29/payment_tranfer_api/internal/usecases"
	mock_test "github.com/RickChaves29/payment_tranfer_api/mock"
)

func TestFunctionGetUser(t *testing.T) {
	usersTest := []usecases.User{
		{
			ID:      "1234",
			Balance: 0,
		},
	}
	mr := mock_test.NewMockRepository(usersTest)
	uc := usecases.NewUserUsecase(mr)

	t.Run("if user id is correct", func(t *testing.T) {
		user, _ := uc.GetUser("1234")

		idExpect := "1234"
		if user.ID != idExpect {
			t.Errorf("have %v want %v", user.ID, idExpect)
		}
	})
	t.Run("if balance is correct", func(t *testing.T) {
		user, _ := uc.GetUser("1234")
		var balanceExpect int64 = 0
		if user.Balance != balanceExpect {
			t.Errorf("have %v want %v", user.Balance, balanceExpect)
		}
	})
	t.Run("if id not exist return error", func(t *testing.T) {
		_, err := uc.GetUser("1235")
		errExpect := "id not exists"
		if err.Error() != errExpect {
			t.Errorf("have %v want %v", err.Error(), errExpect)
		}
	})
	t.Run("if id is empyt return error", func(t *testing.T) {
		_, err := uc.GetUser("")
		errExpect := "id is empty"
		if err.Error() != errExpect {
			t.Errorf("have %v want %v", err.Error(), errExpect)
		}
	})

}
