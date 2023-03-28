package user_test

import (
	"testing"

	"github.com/RickChaves29/payment_tranfer_api/internal/domain/user"
	mock_test "github.com/RickChaves29/payment_tranfer_api/mock"
)

func TestFunctionGetUser(t *testing.T) {
	usersTest := []user.User{
		{
			ID:      "1234",
			Balance: 0,
		},
	}
	mr := mock_test.NewMockRepository(usersTest)
	uc := user.NewUserUsecase(mr)

	t.Run("if user id is correct", func(t *testing.T) {
		user, _ := uc.GetUser("1234")

		idExpect := "1234"
		if user.ID != idExpect {
			t.Errorf("have %v want %v", user.ID, idExpect)
		}
	})
	t.Run("if balance is correct", func(t *testing.T) {
		user, _ := uc.GetUser("1234")
		var balanceExpect uint64 = 0
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

func TestTransferPayment(t *testing.T) {
	usersTest := []user.User{
		{
			ID:      "1234",
			Balance: 1000,
		},
		{
			ID:      "4321",
			Balance: 50,
		},
		{
			ID:      "0000",
			Balance: 0,
		},
	}
	mr := mock_test.NewMockRepository(usersTest)
	uc := user.NewUserUsecase(mr)
	t.Run("if update balance is ok", func(t *testing.T) {
		uc.TransferPayment("1234", "4321", 200)
		if usersTest[0].Balance != 900 && usersTest[1].Balance != 250 {
			t.Errorf("payer have %v want %v", usersTest[0].Balance, 900)
			t.Errorf("receive have %v want %v", usersTest[1].Balance, 250)
		}
	})
	t.Run("if id empty transfer fail", func(t *testing.T) {
		err := uc.TransferPayment("", "4321", 50)
		if err.Error() != "id is empty" {
			t.Errorf("error have %v want %v", err.Error(), "id is empty")
		}
	})
	t.Run("if id not exists transfer fail", func(t *testing.T) {
		err := uc.TransferPayment("2020", "4321", 50)
		if err.Error() != "id not exists" {
			t.Errorf("error have %v want %v", err.Error(), "id not exists")
		}
	})
	t.Run("if the amount is greater than the balance", func(t *testing.T) {
		err := uc.TransferPayment("0000", "4321", 50)
		if err.Error() != "insufficient funds" {
			t.Errorf("error have %v want %v", err.Error(), "insufficient funds")
		}
	})
}
