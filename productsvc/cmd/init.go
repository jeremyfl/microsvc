package cmd

import (
	"productsvc/domain"
	"productsvc/domain/model"
	"productsvc/internal"
	"productsvc/internal/gorm"
	"productsvc/repository"
	"productsvc/service"
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
	return repository.ProductRepositoryImpl{
		DB: db,
	}
}

func initService(repo domain.ProductRepository) domain.Services {
	return domain.Services{
		ProductService: service.ProductServiceImpl{Repository: repo},
	}
}

// InitEntities Initialize the database entities
func InitEntities(db *internal.Database) domain.Services {
	repo := initRepo(db)

	return initService(repo)
}