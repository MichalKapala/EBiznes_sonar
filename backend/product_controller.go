package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ProductController struct {
    Db *gorm.DB
}

func (controller *ProductController) GetAllProducts(c echo.Context) error {
    var products []Product
    controller.Db.Find(&products)
    return c.JSON(http.StatusOK, products)
}

func (controller *ProductController) GetProductById(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    var product Product
    controller.Db.First(&product, id)
    if product.Id == 0 {
        return echo.ErrNotFound
    }
    return c.JSON(http.StatusOK, product)
}

func (controller *ProductController) CreateProduct(c echo.Context) error {
    product := new(Product)
	
    if err := c.Bind(product); err != nil {
        return err
    }

	controller.Db.FirstOrCreate(&product)
    return c.JSON(http.StatusCreated, product)
}


// CAN BE IMPROVED
func (controller *ProductController) UpdateProduct(c echo.Context) error {
    var product Product
	var foundProduct Product

	if err := c.Bind(&product); err != nil {
        return err
    }

	foundProduct = product
    controller.Db.First(&foundProduct, product.Id)

    if product.Id == 0 {
        return echo.ErrNotFound
    }

    controller.Db.Save(&product)
	return c.JSON(http.StatusOK, product)
}

func (controller *ProductController) DeleteProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var product Product
	controller.Db.First(&product, id)
	if product.Id == 0 {
		return echo.ErrNotFound
	}
	controller.Db.Delete(&product)
	return c.NoContent(http.StatusNoContent)
}