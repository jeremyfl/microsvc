package cmd

import (
	"customer/api/controller"
	"customer/domain"
	"customer/domain/model"
	"customer/internal"
	"customer/internal/gorm"
	"customer/repository"
	"customer/service"
)

// initDatabase initialize the database repository
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

// initRepo initialize the repository
func initRepo(db *internal.Database) domain.ProductRepository {
	return repository.ProductRepositoryImpl{
		DB:  db,
	}
}

func initService(repo domain.ProductRepository) domain.Services {
	return domain.Services{
		ProductService: service.ProductServiceImpl{Repository: repo},
	}
}

// initEntities initialize the database entities
func initEntities(db *internal.Database) *controller.Controller {
	customerRepo := initRepo(db)
	customerService := initService(customerRepo)

	authController := &controller.AuthController{Service: customerService.AuthService}

	return &controller.Controller{
		AuthController:     authController,
	}
}
