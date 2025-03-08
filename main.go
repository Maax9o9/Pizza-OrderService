package main

import (
	infra "order-service/src/order/infrastructure"
	routes "order-service/src/order/infrastructure/routes"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	createOrderController := infra.Init()
	routes.OrderRoutes(r, createOrderController)

	r.Run(":8080")
}