package cmd

import (
	"github.com/segmentio/kafka-go"
	"gitlab.com/jeremylo/microsvc/ordersvc/domain"
	"gitlab.com/jeremylo/microsvc/ordersvc/domain/model"
	"gitlab.com/jeremylo/microsvc/ordersvc/internal"
	"gitlab.com/jeremylo/microsvc/ordersvc/internal/gorm"
	"gitlab.com/jeremylo/microsvc/ordersvc/repository"
	"gitlab.com/jeremylo/microsvc/ordersvc/service"
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

func initMessageBroker(writer *kafka.Writer, reader *kafka.Reader) domain.MessageBroker {
	return &domain.MessageBrokerImpl{
		Writer: writer,
		Reader: nil,
	}
}
