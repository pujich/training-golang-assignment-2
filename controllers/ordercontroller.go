package controllers

import (
	"fmt"
	"golang-second-assigment/config"
	"golang-second-assigment/models"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetAllOrders(ctx echo.Context) error {

	db := config.GetDB()

	Order := []models.Order{}
	err := db.Model(&models.Order{}).Preload("Items").Find(&Order).Error

	if err != nil {
		fmt.Println(err)

		return GenerateErrorResponse(ctx, err.Error())
	}

	return GenerateSuccessResponse(ctx, "Get Data Berhasil", Order)
}

func AddOrderByGorm(ctx echo.Context) error {

	db := config.GetDB()

	Orders := models.Order{}

	if err := ctx.Bind(&Orders); err != nil {
		fmt.Println(err)
		return GenerateErrorResponse(ctx, "Format Data Salah")
	}

	err := db.Create(&Orders).Error
	if err != nil {
		fmt.Println(err)
		return GenerateErrorResponse(ctx, err.Error())
	}

	return GenerateSuccessResponse(ctx, "Add Data Success", nil)
}

func UpdateOrderByGorm(ctx echo.Context) error {

	db := config.GetDB()

	id := ctx.Param("id")
	if id == "" {
		return GenerateErrorResponse(ctx, "Data ID Tidak Ditemukan")
	}

	paramID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		return GenerateErrorResponse(ctx, err.Error())
	}

	Orders := models.Order{}

	if err := ctx.Bind(&Orders); err != nil {
		fmt.Println(err)
		return GenerateErrorResponse(ctx, "Format Data Salah")
	}

	Orders.ID = uint(paramID)

	err = db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&Orders).Error
	if err != nil {
		fmt.Println(err)
		return GenerateErrorResponse(ctx, err.Error())
	}

	return GenerateSuccessResponse(ctx, "Update Data Success", Orders)
}

func DeleteOrderByGorm(ctx echo.Context) error {

	db := config.GetDB()

	id := ctx.Param("id")
	if id == "" {
		return GenerateErrorResponse(ctx, "Data ID Tidak Ditemukan")
	}

	paramID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		return GenerateErrorResponse(ctx, err.Error())
	}

	Orders := models.Order{}

	if err := ctx.Bind(&Orders); err != nil {
		fmt.Println(err)
		return GenerateErrorResponse(ctx, "Format Data Salah")
	}

	Orders.ID = uint(paramID)

	tx := db.Begin()
	err = tx.Where("order_id = ?", uint(paramID)).Unscoped().Delete(&models.Item{}).Error
	if err != nil {
		tx.Rollback()
		fmt.Println(err)
		return GenerateErrorResponse(ctx, err.Error())
	}
	err = tx.Unscoped().Delete(&Orders).Error
	if err != nil {
		tx.Rollback()
		fmt.Println(err)
		return GenerateErrorResponse(ctx, err.Error())
	}

	tx.Commit()

	return GenerateSuccessResponse(ctx, "Deleta Data Success", nil)
}
