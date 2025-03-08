package repositorys

import (
    "log"
    "order-service/src/order/infrastructure/adapters"
)

type RabbitRepository struct {
    rabbitMQ *adapters.RabbitMQ
}

func NewRabbitRepository(rabbitMQ *adapters.RabbitMQ) *RabbitRepository {
    return &RabbitRepository{
        rabbitMQ: rabbitMQ,
    }
}

func (r *RabbitRepository) PublishOrder(orderID string) error {
    message := "Orden Enviada: " + orderID
    err := r.rabbitMQ.Publish(message)
    if err != nil {
        log.Printf("Fallo Haciendo la Orden: %v", err)
        return err
    }
    log.Printf("Orden Hecha: %s", orderID)
    return nil
}

func (r *RabbitRepository) Close() {
    r.rabbitMQ.Close()
}