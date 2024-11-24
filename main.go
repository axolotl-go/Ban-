package main

import (
	"github.com/axolotl-go/Bank/db"
	"github.com/axolotl-go/Bank/models"
	"github.com/axolotl-go/Bank/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db.Dbconnection()
	db.DB.AutoMigrate(models.User{})

	app := fiber.New()

	app.Post("/user", routes.CreateUser)
	app.Get("/user/:id", routes.GetUser)
	app.Get("/users", routes.GetUsers)

	// Movements
	app.Post("/deposit/:id", routes.Deposit)
	app.Post("/withdrawal/:id", routes.Withdrawal)

	app.Listen(":8080")

}
