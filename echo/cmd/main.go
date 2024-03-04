package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/shubham8patni/echo/pkg/db"
	"github.com/shubham8patni/echo/pkg/handler"
	"github.com/shubham8patni/echo/pkg/models"
	"net/http"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	err := db.Initialize()
	if err != nil {
		log.Error(err)
		return
	}
	result := db.DB1.AutoMigrate(&models.Brands{}, &models.Products{})
	if result.Error != nil {
		log.Error(result)
		// Handle error here, e.g., log the error or return an error message
	}
	e.GET("/brands", handler.GetAllBrands)
	e.GET("/products", handler.GetAllProducts)
	e.GET("/brands/:param", handler.GetSingleBrand) // param can be numeric string in case of brand ID or string in case of brand name
	e.Logger.Fatal(e.Start(":1323"))
}

func greetings(c echo.Context) error {
	return c.String(http.StatusOK, "Hello There!")
}
