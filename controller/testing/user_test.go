package main

import (
	"kasir/config"
	"kasir/controller"
	"kasir/model"
	"kasir/utils"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func InsertDataUser() error{
	user := model.User{
		ID:        100,
		Name:      "test",
		Email:     "test@gmail.com",
		Password:  "test12",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: gorm.DeletedAt{},
	}

	var err error
	if err = config.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func TestGetUser_Success(t *testing.T) {
	// Create a new Echo instance and a new request.
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	controller.GetUser(c)
	InsertDataUser()
	
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, rec.Body.String(), "success")
}
