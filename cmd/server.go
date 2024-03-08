package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/usawyer/urlShortener/internal/handlers"
	lg "github.com/usawyer/urlShortener/internal/logger"
	"github.com/usawyer/urlShortener/internal/routes"
	"github.com/usawyer/urlShortener/internal/service"
	"github.com/usawyer/urlShortener/internal/storage"
	ch "github.com/usawyer/urlShortener/internal/storage/cache"
	db "github.com/usawyer/urlShortener/internal/storage/database"
)

func main() {
	zapLogger := lg.InitLogger()
	cache := ch.New(zapLogger)
	database := db.New(zapLogger)
	store := storage.New(cache, database)
	srvc := service.New(store, zapLogger)
	handler := handlers.New(srvc)
	app := fiber.New()
	routes.InitRoutes(app, handler)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Fatal(app.Listen(":8080"))
	}()

	<-stop

	if err := app.Shutdown(); err != nil {
		fmt.Println("Error while shutting down:", err)
	}
	fmt.Println("Server gracefully stopped")

}

//TODO
// README
// tests
