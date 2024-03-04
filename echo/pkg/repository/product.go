package repository

import (
	"github.com/shubham8patni/echo/pkg/db"
	"github.com/shubham8patni/echo/pkg/models"
)

func GetAllProductsDB() ([]models.Products, error) {
	var products []models.Products
	result := db.DB1.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}
