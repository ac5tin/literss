package api

import (
	"literss/feed"

	"github.com/gofiber/fiber/v2"
)

func addFeed(c *fiber.Ctx) error {
	type inputReq struct {
		URL      string        `json:"url"`
		FeedType feed.FeedType `json:"feed_type"`
		Name     string        `json:"name"`
	}
	var input inputReq
	if err := c.BodyParser(&input); err != nil {
		c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
		return nil
	}
	switch input.FeedType {
	case feed.FeedTypeRSS:
		if err := feed.FS.AddRSSFeed(input.URL, input.Name); err != nil {
			c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
			return nil
		}
	default:
		c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"success": false,
			"error":   "This feed type is currently not supported",
		})
		return nil
	}

	// success
	c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
	})
	return nil
}

func getFeed(c *fiber.Ctx) error {
	type inputReq struct {
		URL      string        `json:"url"`
		FeedType feed.FeedType `json:"feed_type"`
		Name     string        `json:"name"`
	}
	var input inputReq
	if err := c.BodyParser(&input); err != nil {
		c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
		return nil
	}

	articles := new([]feed.Article)

	switch input.FeedType {
	case feed.FeedTypeRSS:
		fd, err := feed.FS.GetRSSFeed(input.URL, input.Name)
		if err != nil {
			c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
			return nil

		}
		if err := fd.Get(articles); err != nil {
			c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
			return nil

		}
	default:
		c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"success": false,
			"error":   "This feed type is currently not supported",
		})
		return nil
	}
	// success
	c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success":  true,
		"articles": articles,
	})

	return nil
}
