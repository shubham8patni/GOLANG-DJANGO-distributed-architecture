package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/shubham8patni/echo/pkg/db"
	"github.com/shubham8patni/echo/pkg/models"
)

func GetAllBrandsDB() ([]models.Brands, error) {
	var brands []models.Brands
	result := db.DB1.Find(&brands)
	if result.Error != nil {
		return nil, result.Error
	}
	return brands, nil
}

func InsertOneBrand(c echo.Context, brand *models.Brands) error {
	result := db.DB1.Create(brand)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetBrandByName(c echo.Context, brandName string) (models.Brands, error) {
	var brands models.Brands
	result := db.DB1.Where("name = ?", brandName).First(&brands)
	if result.Error != nil {
		return brands, result.Error
	}
	return brands, nil
}

func GetBrandById(c echo.Context, brandId uint) (models.Brands, error) {
	var brands models.Brands
	result := db.DB1.Where("id = ?", brandId).First(&brands)
	if result.Error != nil {
		return brands, result.Error
	}
	return brands, nil
}
