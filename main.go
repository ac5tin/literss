package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	var addr = 8000
	if port, err := strconv.Atoi(os.Getenv("PORT")); err == nil {
		addr = port
	}

	prod := flag.Bool("prod", false, "Enable prefork in Production")

	// fiber app
	app := fiber.New(fiber.Config{
		Prefork: *prod,
	})
	// middleware
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(compress.New())
	app.Use(cors.New())

	// ==== API ROUTES =====
	app.Get("/ping", func(c *fiber.Ctx) error { c.Status(200).Send([]byte("pong")); return nil })

	// start server
	log.Println(fmt.Sprintf("Listening on PORT %d", addr))
	if err := app.Listen(fmt.Sprintf(":%d", addr)); err != nil {
		log.Fatal(err.Error())
	}
}
