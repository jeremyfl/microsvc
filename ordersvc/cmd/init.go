package cmd

import (
	"github.com/segmentio/kafka-go"
	"ordersvc/domain"
	"ordersvc/domain/model"
	"ordersvc/internal"
	"ordersvc/internal/gorm"
	"ordersvc/repository"
	"ordersvc/service"
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

func initService(repo domain.OrderRepository, publisher *kafka.Writer) domain.Services {
	return domain.Services{
		StockService: &service.StockServiceImpl{
			Publisher: publisher,
			Repository: repo,
		},
	}
}

// InitEntities Initialize the database entities
func InitEntities(db *internal.Database, publisher *kafka.Writer) domain.Services {
	repo := initRepo(db)

	return initService(repo, publisher)
}
