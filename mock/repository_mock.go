package mock_test

import (
	"errors"

	"github.com/RickChaves29/payment_tranfer_api/internal/usecases"
)

type MockRepository struct {
	users []usecases.User
}

func NewMockRepository(u []usecases.User) *MockRepository {
	return &MockRepository{
		users: u,
	}
}
func (mr *MockRepository) FindUserByID(id string) (*usecases.User, error) {
	var user *usecases.User
	if id == "" {
		return nil, errors.New("id is empty")
	}
	for _, data := range mr.users {
		if id == data.ID {
			user = &data
			break
		}
	}

	if user == nil {
		return nil, errors.New("id not exists")
	}
	return user, nil
}

func (mr *MockRepository) UpdateBalance(id string, amount int64) (err error) {
	for i, data := range mr.users {
		if id == data.ID {
			mr.users[i].Balance = amount
			break
		} else {
			err = errors.New("update balance fail")

		}
	}
	return err
}
