package elasticsearchhandler

import (
	"github.com/gofiber/fiber/v2"
)

type ElasticsearchHandler interface {
	IndexList(c *fiber.Ctx) error
	IndexSearch(c *fiber.Ctx) error
}
