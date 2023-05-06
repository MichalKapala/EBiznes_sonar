package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type StoreApp struct {
    Db *gorm.DB
    Echo *echo.Echo
}

func (app *StoreApp) Init( e *echo.Echo, db *gorm.DB){

	app.Db = db
	app.Echo = e

	app.InitRoutings()


}

func (app *StoreApp) InitRoutings(){

	productController := &ProductController{app.Db}
	bucketController := &BucketController{app.Db}

	app.Echo.GET("/products", productController.GetAllProducts)
	app.Echo.GET("/products/:id", productController.GetProductById)
    app.Echo.POST("/products", productController.CreateProduct)
    app.Echo.PUT("/products", productController.UpdateProduct)
    app.Echo.DELETE("/products/:id", productController.DeleteProduct)

	app.Echo.GET("/bucket", bucketController.GetAllItems)
	app.Echo.GET("/bucket/:id", bucketController.GetItemById)
	app.Echo.POST("/bucket", bucketController.AddItem)
	app.Echo.PUT("/bucket", bucketController.UpdateItem)
    app.Echo.DELETE("/bucket/:id", bucketController.DeleteItem)

	app.Echo.POST("/bucket/confirmation", bucketController.AddConfirmation)

}

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	
	db, err := gorm.Open(sqlite.Open("store_data.db"), &gorm.Config{})

	if err != nil {
        panic("Bład przy podłączaniu bazy danych!")
    }

	db.AutoMigrate(&Product{}, &Category{}, &Bucket{})

	app := &StoreApp{}
	app.Init(e, db)

	e.Logger.Fatal(e.Start(":8080"))
}