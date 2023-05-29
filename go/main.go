package main

import (
	"elkstack/handlers/elasticsearchhandler"
	"elkstack/services/indexsrv"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	indexService := indexsrv.NewIndexService()

	elasticsearchHandler := elasticsearchhandler.NewElasticsearchHandler(indexService)

	app := fiber.New()

	app.Get("index/list", elasticsearchHandler.IndexList)

	app.Get("index/:index", elasticsearchHandler.IndexSearch)

	log.Fatal(app.Listen(fmt.Sprintf(":%v", 5000)))

}
