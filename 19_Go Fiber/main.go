package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"time"
)

func main() {
	app := fiber.New(fiber.Config{
		IdleTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
		ReadTimeout:  time.Second * 5,
		Prefork:      true,
	})

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello!")
	})

	if fiber.IsChild() {
		fmt.Println("This is Child Process")
	} else {
		fmt.Println("This is Parent Process")
	}
	err := app.Listen("localhost:3000")
	if err != nil {
		panic(err)
	}
}
