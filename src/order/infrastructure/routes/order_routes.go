package routes

import (
	controller "order-service/src/order/infrastructure/controller"
	"github.com/gin-gonic/gin"
)

func OrderRoutes(router *gin.Engine, controller *controller.CreateOrderController) {
	router.POST("/order", controller.Handle)
}