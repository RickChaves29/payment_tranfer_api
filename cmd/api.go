package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/RickChaves29/payment_tranfer_api/internal/data"
	"github.com/RickChaves29/payment_tranfer_api/internal/data/dtos"
	"github.com/RickChaves29/payment_tranfer_api/internal/data/repositories"
	"github.com/RickChaves29/payment_tranfer_api/internal/usecases"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func init() {
	db, err := data.ConnectionDB("postgres", os.Getenv("PAYMENT_DB"))
	if err != nil {
		log.Printf("LOG - [conn-init-error]: %v", err.Error())
	}
	defer db.Close()
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id VARCHAR(200) PRIMARY KEY,
		balance BIGINT DEFAULT 0
	)`)
	if err != nil {
		log.Printf("LOG - [create-table-error]: %v", err.Error())
	}
	_, err = db.Exec(`INSERT INTO users (id, balance) VALUES ('1234', 1000) ON CONFLICT (id) DO NOTHING`)
	if err != nil {
		log.Printf("LOG - [insert-error]: %v", err.Error())
	}
	_, err = db.Exec(`INSERT INTO users (id, balance) VALUES ('4321', 2000) ON CONFLICT (id) DO NOTHING`)
	if err != nil {
		log.Printf("LOG - [insert-error]: %v", err.Error())
	}
}
func main() {
	conn, err := data.ConnectionDB("postgres", os.Getenv("PAYMENT_DB"))
	if err != nil {
		log.Printf("LOG - [conn-main-error]: %v", err.Error())
	}
	r := repositories.NewUserRepository(conn)
	uc := usecases.NewUserUsecase(r)
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Send([]byte(`{"message": "hello world"}`))
	})
	app.Get("/api/v1/users/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		user, err := uc.GetUser(id)
		if err != nil {
			c.SendStatus(400)
			log.Printf("LOG - [get_user-error]: %v", err.Error())
			return err
		}
		userDTO := dtos.UserDTO{
			ID:      user.ID,
			Balance: user.Balance,
		}

		data, err := json.Marshal(userDTO)
		if err != nil {
			log.Printf("LOG - [json-error]: %v", err.Error())
		}
		return c.Send(data)
	})
	app.Post("/api/v1/transfer", func(c *fiber.Ctx) error {
		body := new(dtos.TransferDataDTO)
		err := c.BodyParser(body)
		if err != nil {
			return err
		}
		err = uc.TransferPayment(body.Payer, body.Receive, body.Amount)
		if err != nil {
			c.SendStatus(401)
			log.Printf("LOG - [transfer_payment-error]: %v", err.Error())
			return err
		}
		return nil
	})
	app.Listen(":5000")
}
