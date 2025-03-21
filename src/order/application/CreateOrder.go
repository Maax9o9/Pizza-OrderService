package application

import (
    "errors"
    "order-service/src/order/domain/entities"
    "order-service/src/order/application/repositorys"
)

type CreateOrder struct {
    rabbitRepository *repositorys.RabbitRepository
}

func NewCreateOrder(rabbitRepository *repositorys.RabbitRepository) *CreateOrder {
    return &CreateOrder{
        rabbitRepository: rabbitRepository,
    }
}

func (co *CreateOrder) Execute(order entities.Order) (entities.Order, error) {
    // Asegúrate de que el ID de la orden se esté generando correctamente
    if order.ID == 0 {
        return entities.Order{}, errors.New("order ID is not set")
    }

    err := co.rabbitRepository.PublishOrder(order)
    if err != nil {
        return entities.Order{}, err
    }

    return order, nil
}