package controller

import (
	"employee-app/app/employee"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateEmployeeController(ctx *gin.Context) {
	var request employee.CreateEmployeeRequest
	requestErr := ctx.ShouldBind(&request)
	if requestErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": requestErr.Error(),
		})

		return
	}

	employeeModel, createErr := employee.Create(request)
	if createErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": createErr.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, employee.CreateEmployeeResponse{
		ID:   employeeModel.ID,
		Name: employeeModel.Name,
	})
}

func GetEmployeeController(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	employeeModel, getErr := employee.GetById(id)
	if getErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": getErr.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, employee.GetEmployeeResponse{
		ID:           employeeModel.ID,
		Name:         employeeModel.Name,
		Age:          employeeModel.Age,
		Gender:       employeeModel.Gender,
		Address:      employeeModel.Address,
		Department:   employeeModel.Department,
		MobileNumber: employeeModel.MobileNumber,
		CreatedAt:    employeeModel.CreatedAt,
		UpdatedAt:    employeeModel.UpdatedAt,
	})
}

func UpdateEmployeeController(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	var request employee.UpdateEmployeeRequest
	requestErr := ctx.ShouldBind(&request)
	if requestErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": requestErr.Error(),
		})

		return
	}

	employeeModel, updateErr := employee.Update(id, request)
	if updateErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": updateErr.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, employee.UpdateEmployeeResponse{
		ID: employeeModel.ID,
	})
}

func DeleteEmployeeController(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	deleteErr := employee.Delete(id)
	if deleteErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": deleteErr.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, employee.UpdateEmployeeResponse{
		ID: id,
	})
}
