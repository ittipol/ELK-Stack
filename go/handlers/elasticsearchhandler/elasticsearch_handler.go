package elasticsearchhandler

import (
	"elkstack/services/indexsrv"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type elasticsearchHandler struct {
	indexService indexsrv.IndexService
}

func NewElasticsearchHandler(indexService indexsrv.IndexService) ElasticsearchHandler {
	return &elasticsearchHandler{indexService}
}

func (obj elasticsearchHandler) IndexList(c *fiber.Ctx) error {

	response, err := obj.indexService.List()

	if err != nil {
		return fiber.ErrInternalServerError
	}

	c.Status(http.StatusOK)
	return c.JSON(response)
}

func (obj elasticsearchHandler) IndexSearch(c *fiber.Ctx) error {

	index := c.Params("index")

	obj.indexService.Search(index)

	c.Status(http.StatusOK)
	return c.JSON("ok")
}
