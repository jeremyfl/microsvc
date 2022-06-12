package cmd

import (
	"customer/domain"
	"customer/domain/model"
	"customer/internal"
	"customer/internal/gorm"
	"customer/repository"
	"customer/service"
)

// InitDatabase Initialize the database repository
func InitDatabase() *internal.Database {
	db, err := gorm.NewClient()
	if err != nil {
		panic("Error when creating new database instances")
	}

	if err = db.AutoMigrate(&model.Product{}); err != nil {
		return nil
	}

	return &internal.Database{DB: db}
}

// InitRepo Initialize the repository
func InitRepo(db *internal.Database) domain.ProductRepository {
	return repository.ProductRepositoryImpl{
		DB: db,
	}
}

func InitService(repo domain.ProductRepository) domain.Services {
	return domain.Services{
		ProductService: service.ProductServiceImpl{Repository: repo},
	}
}

// InitEntities Initialize the database entities
func InitEntities(db *internal.Database) domain.Services {
	repo := InitRepo(db)

	return InitService(repo)
}