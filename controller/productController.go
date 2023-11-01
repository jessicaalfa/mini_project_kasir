package controller

import (
	"kasir/config"
	"kasir/model"
	"kasir/utils"
	"kasir/utils/req"
	"kasir/utils/res"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetProduct(c echo.Context) error {
	var product []model.Product
	if err := config.DB.Find(&product).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Could not find product. Internal Server Error"))
	}

	if len(product) == 0 {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse("Empty Product"))
	}

	response := res.ProductRes(product)
	return c.JSON(http.StatusOK, utils.SuccessResponse("Success Get All Products", response))
}

func GetIDProduct(c echo.Context) error {
	id := c.Param("id")
	var product model.Product
	if err := config.DB.First(&product, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse("Product not Found"))
	}

	response := res.ProductIDRes(&product)

	return c.JSON(http.StatusOK, utils.SuccessResponse("Success Get ID Product", response))
}

func CreateProduct(c echo.Context) error {
	var create model.Product
	var request req.ProductRequest
	c.Bind(&request)

	create.Name = request.Name
	create.Harga = request.Harga

	if err := config.DB.Save(&create).Error; err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Could not save product"))
	}

	response := res.ProductIDRes(&create)
	return c.JSON(http.StatusCreated, utils.SuccessResponse("Product Created Successfully", response))
}

func UpdateProduct(c echo.Context) error {
	var product model.Product
	var request req.ProductRequest
	id, _ := strconv.Atoi(c.Param("id"))
	c.Bind(&request)

	result := config.DB.First(&product, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Error Updating Product"))
	}
	config.DB.Model(&product).Updates(request)
	response := res.ProductIDRes(&product)
	return c.JSON(http.StatusOK, utils.SuccessResponse("Success Update Id Product", response))
}

func DeleteProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var product []model.Product
	result := config.DB.Delete(&product, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": result.Error,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete id books",
		"product": product,
	})
}
