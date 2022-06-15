package cmd

import (
	"github.com/segmentio/kafka-go"
	"gitlab.com/jeremylo/microsvc/ordersvc/domain"
	"gitlab.com/jeremylo/microsvc/ordersvc/domain/model"
	"gitlab.com/jeremylo/microsvc/ordersvc/internal"
	"gitlab.com/jeremylo/microsvc/ordersvc/internal/gorm"
	"gitlab.com/jeremylo/microsvc/ordersvc/repository"
	"gitlab.com/jeremylo/microsvc/ordersvc/service"
	"go.opentelemetry.io/otel/sdk/resource"
	"time"

	"log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

// initDatabase Initialize the database repository
func initDatabase() *internal.Database {
	db, err := gorm.NewClient()
	if err != nil {
		panic("Error when creating new database instances")
	}

	if err = db.AutoMigrate(&model.Order{}); err != nil {
		return nil
	}

	return &internal.Database{DB: db}
}

// initRepo Initialize the repository
func initRepo(db *internal.Database) domain.OrderRepository {
	return &repository.OrderRepositoryImpl{
		DB: db,
	}
}

func initService(repo domain.OrderRepository, messageBroker domain.MessageBroker) domain.Services {
	return domain.Services{
		StockService: &service.StockServiceImpl{
			MessageBroker: messageBroker,
			Repository:    repo,
		},
	}
}

// InitEntities Initialize the database entities
func InitEntities(db *internal.Database, messageBroker domain.MessageBroker) domain.Services {
	repo := initRepo(db)

	return initService(repo, messageBroker)
}

func initMessageReader(topic string) *kafka.Reader {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		Topic:     topic,
		Partition: 0,
		GroupID:   "ordersvc-listener",
	})

	return r
}

func initMessageWriter() *kafka.Writer {
	return &kafka.Writer{
		Addr:         kafka.TCP("localhost:9092"),
		Balancer:     &kafka.LeastBytes{},
		BatchTimeout: 10 * time.Millisecond,
	}
}

func initMessageBroker(writer *kafka.Writer, reader *kafka.Reader) domain.MessageBroker {
	return &domain.MessageBrokerImpl{
		Writer: writer,
		Reader: nil,
	}
}

func initTracer() *sdktrace.TracerProvider {
	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint("http://127.0.0.1:14268/api/traces")))
	if err != nil {
		log.Fatal(err)
	}
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(
			resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceNameKey.String("order-svc"),
			)),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return tp
}
