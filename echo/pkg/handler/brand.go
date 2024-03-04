package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/shubham8patni/echo/pkg/dto"
	"github.com/shubham8patni/echo/pkg/models"
	"github.com/shubham8patni/echo/pkg/repository"
	"net/http"
	"strconv"
)

//import (
//	"fmt"
//	"github.com/shubham8patni/echo/pkg/repository"
//)

func GetAllBrands(c echo.Context) error {

	responses, err := repository.GetAllBrandsDB()
	if err != nil {
		return c.String(http.StatusInternalServerError, "err")
	}

	brands := make([]dto.AllBrandsDtoRes, 0)

	for _, response := range responses {
		brand := dto.AllBrandsDtoRes{
			ID:   int64(response.ID),
			Name: response.Name,
		}
		brands = append(brands, brand)
	}

	return c.JSON(http.StatusOK, brands)
}

func CreateNewBrand(c echo.Context, brand *models.Brands) error {

	newBrand := &models.Brands{Name: brand.Name}
	err := repository.InsertOneBrand(c, newBrand)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, "{status: New Entry Created.}")

}

func GetSingleBrand(c echo.Context, brandParam string) error {

	var result models.Brands
	brandID, err := strconv.Atoi(brandParam)
	if err != nil {
		log.Info("Brand Parameter is String")
		result, err = repository.GetBrandByName()
	} else if brandID > 0 {
		log.Info("Brand Parameter is Numeric")
		result, err = repository.GetBrandById(c, uint(brandID))
		if err != nil {
			return err
		}
	}

	//var result interface{}
	//var err error
	if brand.ID == 0 && brand.Name != "" {
		result, err = repository.GetBrandByName(c, brand.Name)
		if err != nil {
			return nil
		}
	}
	if brand.ID > 0 && brand.Name == "" {
		result, err = repository.GetBrandById(c, brand.ID)
		if err != nil {
			return nil
		}
	}
	return c.JSON(http.StatusOK, result)
}
