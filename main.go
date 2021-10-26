package main

import (
	database "employee-app/config"
	"log"

	controller "employee-app/app/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	database.CreateDatabaseConnection()

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	employeeRoute := router.Group("/employee")
	{
		employeeRoute.POST("/", controller.CreateEmployeeController)
		employeeRoute.GET("/:id", controller.GetEmployeeController)
		employeeRoute.PUT("/:id", controller.UpdateEmployeeController)
		employeeRoute.DELETE("/:id", controller.DeleteEmployeeController)
	}

	if err := router.Run(":5000"); err != nil {
		log.Fatal("Gin run error", err)
	}
}
