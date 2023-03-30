package user_test

import (
	"errors"
	"testing"

	"github.com/RickChaves29/payment_tranfer_api/internal/domain/user"
)

type UserRepoTest struct {
	repo []user.UserEntity
}

func NewRepoTest(users []user.UserEntity) *UserRepoTest {
	return &UserRepoTest{
		repo: users,
	}
}
func (rt *UserRepoTest) Save(data user.CreateUserEntity) error {
	user := user.UserEntity{
		ID:       uint64(len(rt.repo) - 1),
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		Balance:  1000,
	}
	rt.repo = append(rt.repo, user)
	return nil
}
func (rt *UserRepoTest) FindUserByEmail(email string) (*user.UserEntity, error) {
	var result *user.UserEntity
	for _, user := range rt.repo {
		if user.Email == email {
			result = &user
			break
		}
	}
	return result, nil
}

func (rt *UserRepoTest) UpdateBalance(id uint64, amount uint64) error {
	for i, user := range rt.repo {
		if id == user.ID {
			rt.repo[i].Balance = amount
			break
		} else {
			return errors.New("update balance fail")
		}
	}
	return nil
}
func TestIfCreateIsCorrect(t *testing.T) {
	r := NewRepoTest([]user.UserEntity{})
	uc := user.NewUserUsecase(r)
	data := user.CreateUserEntity{
		Name:     "any name 1",
		Password: "1234",
		Email:    "any@gmail.com",
	}
	uc.CreateNewUser(data)
	t.Run("should create user success", func(t *testing.T) {

		if len(r.repo) == 0 {
			t.Errorf("want 1 have %d", 0)
		}
	})
}

func TestIfReturnUser(t *testing.T) {
	r := NewRepoTest([]user.UserEntity{
		{
			ID:       1,
			Name:     "any",
			Email:    "any@gmail.com",
			Password: "1234",
			Balance:  4000,
		},
		{
			ID:       2,
			Name:     "any 1",
			Email:    "any1@gmail.com",
			Password: "4321",
			Balance:  5000,
		},
	})
	uc := user.NewUserUsecase(r)
	t.Run("should return user 1 success", func(t *testing.T) {
		result, _ := uc.GetUserByEmail("any@gmail.com")

		if result.ID != 1 {
			t.Errorf("have %v want %v", result.ID, 1)
		}
	})
	t.Run("should return user 2 success", func(t *testing.T) {
		result, _ := uc.GetUserByEmail("any1@gmail.com")

		if result.ID != 2 {
			t.Errorf("have %v want %v", result.ID, 2)
		}
	})
}

func TestTransferPayment(t *testing.T) {
	usersTest := []user.UserEntity{
		{
			ID:       1,
			Name:     "any",
			Email:    "any@gmail.com",
			Password: "1234",
			Balance:  4000,
		},
		{
			ID:       2,
			Name:     "any 1",
			Email:    "any1@gmail.com",
			Password: "4321",
			Balance:  5000,
		},
	}
	r := NewRepoTest(usersTest)
	uc := user.NewUserUsecase(r)
	data := user.UserTransferPaymentEntity{
		Payer:   "any@gmail.com",
		Receive: "any1@gmail.com",
		Amount:  1000,
	}
	uc.TransferPayment(data)

	t.Run("should update balance is success", func(t *testing.T) {
		if r.repo[0].Balance != 3000 {
			t.Errorf("have %v want %v", r.repo[0].Balance, 3000)
		}
	})
}
