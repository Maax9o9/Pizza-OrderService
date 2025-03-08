package controller

import (
	"order-service/src/order/application"
	"order-service/src/order/domain/entities"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateOrderController struct {
	useCase *application.CreateOrder
}

func NewCreateOrderController(uc *application.CreateOrder) *CreateOrderController {
	return &CreateOrderController{useCase: uc}
}

func (c *CreateOrderController) Handle(ctx *gin.Context) {
	var newOrder entities.Order

	if err := ctx.ShouldBindJSON(&newOrder); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	log.Printf("Order received: %+v", newOrder)

	order, err := c.useCase.Execute(newOrder)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, order)
}
