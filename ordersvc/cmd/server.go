package cmd

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/segmentio/kafka-go"
	"log"
	"ordersvc/handler"
	"os"
	"os/signal"
	"syscall"
)

func initMessageWriter() *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    "order.created",
		Balancer: &kafka.LeastBytes{},
	}
}

func Serve() {
	app := fiber.New()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c

		_ = <-c
		_ = app.Shutdown()

		os.Exit(0)
	}()

	kw := initMessageWriter()
	defer kw.Close()

	db := initDatabase()
	entities := InitEntities(db, kw)

	app.Use(recover.New())
	app.Use(logger.New())

	orderHandler := handler.Handler{Services: entities}
	app.Post("/api/v1/orders", orderHandler.Handle)

	if err := app.Listen(":3000"); err != nil {
		log.Println("Error when listen from the HTTP Framework ", err)
	}
}
