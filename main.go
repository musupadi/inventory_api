package main

import (
	"apiinventory/Config"
	"apiinventory/Controller"

	"github.com/gin-gonic/gin"
)

func main() {
	Config.Connect()

	r := gin.Default()

	// r.POST("/users", controller.CreateUser)
	r.POST("/Login", Controller.Login)
	// r.GET("/users", controller.GetUsers)

	// //Product
	// r.GET("/Product/:id_shop", controller.GetProduct)

	// //Transaction
	// r.POST("/Transaction", controller.Transaction)
	// r.POST("/Sumtransaction", controller.SumTransaction)

	// // r.PUT("/users/:username", controller.UpdateUser)
	// r.DELETE("/users/:username", controller.DeleteUser)

	r.Run(":8080")
}
