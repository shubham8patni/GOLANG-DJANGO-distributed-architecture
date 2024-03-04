package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/shubham8patni/echo/pkg/repository"
	"net/http"
)

func GetAllProducts(c echo.Context) error {
	response, err := repository.GetAllProductsDB()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, response)
}
