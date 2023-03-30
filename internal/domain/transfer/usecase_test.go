package transfer_test

import (
	"testing"
	"time"

	"github.com/RickChaves29/payment_tranfer_api/internal/domain/transfer"
)

type TransferRepoTest struct {
	repo []transfer.TransferEntity
}

func NewRepoTest(transfer []transfer.TransferEntity) *TransferRepoTest {
	return &TransferRepoTest{
		repo: transfer,
	}
}
func (rt *TransferRepoTest) Save(userID uint64, data transfer.CreateTransfer) error {
	tranfer := transfer.TransferEntity{
		ID:           uint(len(rt.repo) + 1),
		PayerName:    data.PayerName,
		ReceiveName:  data.ReceiveName,
		Amount:       data.Amount,
		TransferDate: time.Now().GoString(),
	}
	rt.repo = append(rt.repo, tranfer)
	return nil
}

func (rt *TransferRepoTest) FindAllTransfers(userId uint64) ([]transfer.TransferEntity, error) {
	return rt.repo, nil
}

func TestIfCreateNewTransferIsSuccess(t *testing.T) {
	rt := NewRepoTest([]transfer.TransferEntity{})
	uc := transfer.NewTaskferUsecase(rt)

	newTransfer := transfer.CreateTransfer{
		PayerName:   "João 1",
		ReceiveName: "João 2",
		Amount:      100,
	}
	uc.CreateNewTransfer(1, newTransfer)
	t.Run("should create new transfer success", func(t *testing.T) {
		if len(rt.repo) == 0 {
			t.Errorf("have %v want %v", len(rt.repo), 1)
		}
	})
}

func TestIfListAllTransfersIsSuccess(t *testing.T) {
	transfers := []transfer.TransferEntity{
		{
			ID:           1,
			PayerName:    "any 1",
			ReceiveName:  "any 2",
			Amount:       100,
			TransferDate: time.Now().GoString(),
		},
		{
			ID:           2,
			PayerName:    "any 1",
			ReceiveName:  "any 3",
			Amount:       200,
			TransferDate: time.Now().GoString(),
		},
		{
			ID:           3,
			PayerName:    "any 1",
			ReceiveName:  "any 5",
			Amount:       600,
			TransferDate: time.Now().GoString(),
		},
	}

	rt := NewRepoTest(transfers)
	uc := transfer.NewTaskferUsecase(rt)

	t.Run("should return list of transfers when user request", func(t *testing.T) {
		transfers, _ := uc.ListAllTransfers(2)

		if len(rt.repo) != len(transfers) {
			t.Errorf("have %v want %v", len(rt.repo), len(transfers))
		}
	})
}
