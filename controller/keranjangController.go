package controller

import (
	"kasir/config"
	"kasir/model"
	"kasir/utils"
	"kasir/utils/req"
	"kasir/utils/res"
	"strconv"

	//"kasir/utils/res"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetKeranjang(c echo.Context) error {
	var keranjang []model.Keranjang

	var response []res.KeranjangRes
	if keranjangError := config.DB.Preload("Kasir").Preload("Product").Find(&keranjang); keranjangError.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, keranjangError.Error)
	}

	response = res.GetKeranjangAll(keranjang)
	return c.JSON(http.StatusOK, utils.SuccessResponse("Success Get All Keranjang", response))
}

func GetIDKeranjang(c echo.Context) error {
	id := c.Param("id")

	var response res.KeranjangRes
	var keranjang model.Keranjang // Inisialisasi keranjang

	if keranjangError := config.DB.Preload("Kasir").Preload("Product").First(&keranjang, "id = ?", id); keranjangError.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, keranjangError.Error)
	}

	response.GetKeranjangRes(keranjang)

	return c.JSON(http.StatusOK, utils.SuccessResponse("Create Successfully", response))
}

func CreateKeranjang(c echo.Context) error {
	keranjang := model.Keranjang{}
	var KeranjangRequest req.KeranjangRequest
	c.Bind(&KeranjangRequest)

	var product model.Product
	productError := config.DB.First(&product, "id = ?", KeranjangRequest.ProductID)
	if productError.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, productError.Error)
	}

	var kasir model.User
	kasirError := config.DB.First(&kasir, "id = ?", KeranjangRequest.KasirID)
	if kasirError.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, kasirError.Error)
	}

	keranjang.KasirID = kasir.ID
	keranjang.ProductID = product.ID
	keranjang.JumlahBarang = KeranjangRequest.JumlahBarang
	keranjang.Status = "payment pending"

	harga, _ := strconv.Atoi(product.Harga)

	keranjang.TotalHarga = uint(harga) * KeranjangRequest.JumlahBarang

	if err := config.DB.Save(&keranjang).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// mengambil data keranjang yang telah dibuat beserta data relasinya
	// mengambil id kerangjang yang dibuat dari DB.Save
	var response res.KeranjangRes
	idKeranjang := keranjang.ID
	if keranjangError := config.DB.Preload("Kasir").Preload("Product").First(&keranjang, "id = ?", idKeranjang); keranjangError.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, keranjangError.Error)
	}
	response.GetKeranjangRes(keranjang)

	return c.JSON(http.StatusOK, utils.SuccessResponse("Create Successfully", response))
}

func UpdateKeranjang(c echo.Context) error {
	var keranjang model.Keranjang
	c.Bind(&keranjang)

	result := config.DB.Save(&keranjang)
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorResponse("Error Updating Keranjang"), result)
	}

	var response res.KeranjangRes
	return c.JSON(http.StatusOK, utils.SuccessResponse("Success Update Id Keranjang", response))
}

func DeleteKeranjang(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var keranjang []model.Keranjang
	if err := config.DB.First(&keranjang, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Error keranjang")
	}

	if err := config.DB.Delete(&keranjang).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":   "success delete id books",
		"Keranjang": keranjang,
	})
}
