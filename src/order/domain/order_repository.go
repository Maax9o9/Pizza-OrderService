package domain

import (
	"order-service/src/order/domain/entities"
)

type OrderRepository interface {
	CreateOrder(order entities.Order) (entities.Order, error)
}