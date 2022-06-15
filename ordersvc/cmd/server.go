package cmd

import (
	"context"
	"github.com/gofiber/contrib/otelfiber"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gitlab.com/jeremylo/microsvc/lib"
	"gitlab.com/jeremylo/microsvc/ordersvc/handler/http"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Serve() {
	ctx := context.Background()

	tp := lib.InitTracer("order-svc")
	defer func() {
		if err := tp.Shutdown(ctx); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	app := fiber.New()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c

		_ = <-c
		_ = app.Shutdown()

		os.Exit(0)
	}()

	kw := lib.InitMessageWriter()
	defer kw.Close()

	mb := initMessageBroker(kw, nil)
	db := initDatabase()
	entities := InitEntities(db, mb)

	app.Use(otelfiber.Middleware("order-svc"))
	app.Use(recover.New())
	app.Use(logger.New())

	orderHandler := http.Handler{Services: entities}
	app.Post("/api/v1/orders", orderHandler.Handle)

	if err := app.Listen(":3000"); err != nil {
		log.Println("Error when listen from the HTTP Framework ", err)
	}
}
