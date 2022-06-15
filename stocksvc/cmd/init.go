package cmd

import (
	"gitlab.com/jeremylo/microsvc/stocksvc/domain"
	"gitlab.com/jeremylo/microsvc/stocksvc/domain/model"
	"gitlab.com/jeremylo/microsvc/stocksvc/internal"
	"gitlab.com/jeremylo/microsvc/stocksvc/internal/gorm"
	"gitlab.com/jeremylo/microsvc/stocksvc/repository"
	"gitlab.com/jeremylo/microsvc/stocksvc/service"
)

// initDatabase Initialize the database repository
func initDatabase() *internal.Database {
	db, err := gorm.NewClient()
	if err != nil {
		panic("Error when creating new database instances")
	}

	if err = db.AutoMigrate(&model.Stock{}); err != nil {
		return nil
	}

	return &internal.Database{DB: db}
}

// initRepo Initialize the repository
func initRepo(db *internal.Database) domain.StockRepository {
	return &repository.StockRepositoryImpl{
		DB: db,
	}
}

func initService(repo domain.StockRepository) domain.Services {
	return domain.Services{
		StockService: &service.StockServiceImpl{
			Repository: repo,
		},
	}
}

// InitEntities Initialize the database entities
func InitEntities(db *internal.Database) domain.Services {
	repo := initRepo(db)

	return initService(repo)
}
