package routes

import (
	"kasir/controller"
	"kasir/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {

	e := echo.New()

	e.Use(middleware.NotFoundHandler)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to RESTful API Services")
	})

	e.GET("/users", controller.GetUser)
	e.GET("/users/:id", controller.GetUserID)
	e.POST("/users/register", controller.RegisterUser)
	e.POST("/users/login", controller.LoginUser)
	e.PUT("/users/:id", controller.UpdateUser)
	e.DELETE("/users/:id", controller.DeleteUser)

	e.GET("/product", controller.GetProduct)
	e.GET("/product/:id", controller.GetIDProduct)
	e.POST("/product", controller.CreateProduct)
	e.PUT("/product/:id", controller.UpdateProduct)
	e.DELETE("/product/:id", controller.DeleteProduct)

	e.GET("/keranjang", controller.GetKeranjang)
	e.GET("/keranjang/:id", controller.GetIDKeranjang)
	e.POST("/keranjang", controller.CreateKeranjang)
	e.PUT("/keranjang/:id", controller.UpdateKeranjang)
	e.DELETE("/keranjang/:id", controller.DeleteKeranjang)

	return e

}

// get semua produk
// post json array of object produk ID dan jumlah
// perlu 3 API post keranjang, pembayaran
// response get product adl array dari product
// response keranjang detail
