package api

import "github.com/gofiber/fiber/v2"

func Routes(router *fiber.Router) {
	(*router).Post("/add_feed", addFeed)
	(*router).Post("/get_feed", getFeed)
}
