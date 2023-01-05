package main

import (
	"Erajaya/internal/controller"
	"Erajaya/internal/repository"
	"Erajaya/internal/repository/proxy"
	"Erajaya/internal/service"
	"Erajaya/pkg/cache"
	"Erajaya/pkg/database"
	"github.com/labstack/echo/v4"
	"log"
)

func main() {

	db, err := database.NewPostgreClient()

	cache, err := cache.NewRedigoClient()
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewProductRepository(db)
	repoProxy := proxy.NewProductRepository(cache, *repo)

	svc := service.NewProductService(repoProxy)
	ctrl := controller.NewProductController(svc)

	e := echo.New()

	e.POST("/add-product", ctrl.AddProduct)
	e.POST("/list-product", ctrl.ListProduct)

	e.Logger.Fatal(e.Start(":1323"))
}
