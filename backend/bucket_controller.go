package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type BucketController struct {
    Db *gorm.DB
}

func (controller *BucketController) GetAllItems(c echo.Context) error {
    var products []Bucket
    controller.Db.Find(&products)
    return c.JSON(http.StatusOK, products)
}

func (controller *BucketController) GetItemById(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    var item Bucket
    controller.Db.First(&item, id)
    if item.ProductId == 0 {
        return echo.ErrNotFound
    }
    return c.JSON(http.StatusOK, item)
}

func (controller *BucketController) AddItem(c echo.Context) error {
    item := new(Bucket)
    product := new(Product)
	
    if err := c.Bind(item); err != nil {
        return err
    }

    controller.Db.First(&product, item.ProductId)

    if product.Id == 0 {
        return echo.ErrNotFound
    } 

	controller.Db.FirstOrCreate(&item)
    return c.JSON(http.StatusCreated, item)
}


// CAN BE IMPROVED
func (controller *BucketController) UpdateItem(c echo.Context) error {
    var item Bucket
	var foundItem Bucket

	if err := c.Bind(&item); err != nil {
        return err
    }

	foundItem = item
    controller.Db.First(&foundItem, item.ProductId)

    if item.ProductId == 0 {
        return echo.ErrNotFound
    }

    controller.Db.Save(&item)
	return c.JSON(http.StatusOK, item)
}

func (controller *BucketController) DeleteItem(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var item Bucket
	controller.Db.First(&item, id)
	if item.ProductId == 0 {
		return echo.ErrNotFound
	}
	controller.Db.Delete(&item)
	return c.NoContent(http.StatusNoContent)
}

func (controller *BucketController) AddConfirmation(c echo.Context) error {
    confirmation := new(Confirmation)

    if err := c.Bind(confirmation); err != nil {
        println("ERROR BINDING")
        return err
    }

    // Here coud add some confirmation handling but will go easy way to just print details

    println("Recieved order confirmation! ")
    println("Order will be sent on: ")
    println(confirmation.Name)
    println(confirmation.Address)
    println(confirmation.City)
    println(confirmation.State)
    println(confirmation.Zip)

    return c.JSON(http.StatusOK, confirmation)
}