package controller

import (
	"Erajaya/internal"
	"Erajaya/internal/entity"
	"Erajaya/pkg"
	"errors"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type ProductController struct {
	productSvc internal.ProductServiceItf
}

func NewProductController(productSvc internal.ProductServiceItf) *ProductController {
	return &ProductController{
		productSvc: productSvc,
	}
}

func (p *ProductController) AddProduct(c echo.Context) error {
	ctx := c.Request().Context()

	request := &entity.Product{}
	if err := c.Bind(request); err != nil {
		return err
	}

	result, err := p.productSvc.AddProduct(ctx, request)
	if err != nil {
		log.Println(err)
		return pkg.ErrorResponse(c, 500, errors.New("Internal Server Error"))
	}

	return c.JSON(http.StatusOK, result)
}

func (p *ProductController) ListProduct(c echo.Context) error {
	ctx := c.Request().Context()

	request := &entity.ListProductRequest{}
	if err := c.Bind(request); err != nil {
		return err
	}

	result, err := p.productSvc.ListProduct(ctx, request)
	if err != nil {
		log.Println(err)
		return pkg.ErrorResponse(c, 500, errors.New("Internal Server Error"))
	}

	return c.JSON(http.StatusOK, result)
}
