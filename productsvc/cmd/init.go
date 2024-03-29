package cmd

import (
	"gitlab.com/jeremylo/microsvc/grpc/model/stock"
	"gitlab.com/jeremylo/microsvc/productsvc/domain"
	"gitlab.com/jeremylo/microsvc/productsvc/domain/model"
	"gitlab.com/jeremylo/microsvc/productsvc/internal"
	"gitlab.com/jeremylo/microsvc/productsvc/internal/gorm"
	"gitlab.com/jeremylo/microsvc/productsvc/repository"
	"gitlab.com/jeremylo/microsvc/productsvc/service"
)

// initDatabase Initialize the database repository
func initDatabase() *internal.Database {
	db, err := gorm.NewClient()
	if err != nil {
		panic("Error when creating new database instances")
	}

	if err = db.AutoMigrate(&model.Product{}); err != nil {
		return nil
	}

	return &internal.Database{DB: db}
}

// initRepo Initialize the repository
func initRepo(db *internal.Database) domain.ProductRepository {
	return &repository.ProductRepositoryImpl{
		DB: db,
	}
}

func initService(repo domain.ProductRepository, stockServiceClient stock.StockServiceClient) domain.Services {
	return domain.Services{
		ProductService: &service.ProductServiceImpl{
			Repository:         repo,
			StockServiceClient: stockServiceClient,
		},
	}
}

// InitEntities Initialize the database entities
func InitEntities(db *internal.Database, stockServiceClient stock.StockServiceClient) domain.Services {
	repo := initRepo(db)

	return initService(repo, stockServiceClient)
}
