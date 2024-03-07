package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/usawyer/urlShortener/internal/handlers"
	"github.com/usawyer/urlShortener/internal/routes"
	"github.com/usawyer/urlShortener/internal/service"
	"github.com/usawyer/urlShortener/internal/storage"
	db "github.com/usawyer/urlShortener/internal/storage/database"
	"go.uber.org/zap"

	"log"
)

func main() {
	config := zap.NewDevelopmentConfig()
	logs, err := config.Build()
	if err != nil {
		log.Fatal(err)
	}

	database := db.New(logs)
	store := storage.New(database)
	srvc := service.New(store)
	handler := handlers.New(srvc)

	app := fiber.New()
	app.Use(logger.New())
	routes.InitRoutes(app, handler)
	log.Fatal(app.Listen(":8080"))
}

//TODO context
// logger
// tests
// redis initialization
// database initialization
// README
// dockerfile
