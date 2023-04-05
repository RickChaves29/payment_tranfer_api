package main

import (
	"log"
	"strconv"

	"github.com/RickChaves29/payment_tranfer_api/internal/data"
	"github.com/RickChaves29/payment_tranfer_api/internal/data/dtos"
	"github.com/RickChaves29/payment_tranfer_api/internal/data/repositories"
	"github.com/RickChaves29/payment_tranfer_api/internal/domain/transfer"
	"github.com/RickChaves29/payment_tranfer_api/internal/domain/user"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {
	conn, err := data.ConnectionDB("postgres", "postgres://postgres:12345@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Printf("LOG - [conn-main-error]: %v", err.Error())
	}
	userRepository := repositories.NewUserRepository(conn)
	transferRepository := repositories.NewTransferRepository(conn)
	userUsecase := user.NewUserUsecase(userRepository)
	transferUsecase := transfer.NewTaskferUsecase(transferRepository)
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Send([]byte(`{"message": "hello world"}`))
	})
	v1 := app.Group("/api/v1")
	v1.Post("/users", func(c *fiber.Ctx) error {
		var newUser *dtos.CreateUserDTO
		err := c.BodyParser(&newUser)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return err
		}
		err = userUsecase.CreateNewUser(user.CreateUserEntity{
			Name:     newUser.Name,
			Email:    newUser.Email,
			Password: newUser.Password,
		})
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return err
		}
		c.Status(fiber.StatusCreated)
		return nil
	})
	v1.Get("/users/:user_id", func(c *fiber.Ctx) error {
		userID := c.Params("user_id")
		id, err := strconv.Atoi(userID)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return err
		}
		user, err := userUsecase.GetUserById(uint64(id))
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return err
		}
		userDTO := dtos.UserDTO{
			ID:      user.ID,
			Name:    user.Name,
			Email:   user.Email,
			Balance: user.Balance,
		}
		c.Status(fiber.StatusOK)
		return c.JSON(userDTO)
	})
	v1.Post("/users/:user_id/transfers", func(c *fiber.Ctx) error {
		userID := c.Params("user_id")
		transferPaymentDTO := dtos.UserTransferPaymentDTO{}
		c.BodyParser(&transferPaymentDTO)
		id, err := strconv.Atoi(userID)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return err
		}
		newTransferPayment := user.UserTransferPaymentEntity{
			Payer:   transferPaymentDTO.Payer,
			Receive: transferPaymentDTO.Receive,
			Amount:  transferPaymentDTO.Amount,
		}
		err = userUsecase.TransferPayment(newTransferPayment)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return err
		}
		payer, err := userUsecase.GetUserByEmail(newTransferPayment.Payer)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return err
		}
		receive, err := userUsecase.GetUserByEmail(newTransferPayment.Receive)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return err
		}
		newTransfer := transfer.CreateTransfer{
			PayerName:   payer.Name,
			ReceiveName: receive.Name,
			Amount:      newTransferPayment.Amount,
		}
		err = transferUsecase.CreateNewTransfer(uint64(id), newTransfer)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return err
		}
		c.Status(fiber.StatusCreated)
		return nil
	})
	v1.Get("/users/:user_id/transfers", func(c *fiber.Ctx) error {
		var transfersDTO []dtos.TransferDTO
		userID := c.Params("user_id")
		id, err := strconv.Atoi(userID)
		if err != nil {
			return err
		}
		transfers, err := transferUsecase.ListAllTransfers(uint64(id))
		if err != nil {
			return err
		}
		for _, transfer := range transfers {
			transferDTO := dtos.TransferDTO{
				ID:           transfer.ID,
				PayerName:    transfer.PayerName,
				ReceiveName:  transfer.ReceiveName,
				Amount:       transfer.Amount,
				TransferDate: transfer.TransferDate,
			}

			transfersDTO = append(transfersDTO, transferDTO)
		}
		return c.JSON(transfersDTO)
	})
	app.Listen(":3000")
}
