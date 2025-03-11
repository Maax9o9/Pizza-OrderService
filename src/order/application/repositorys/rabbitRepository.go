package repositorys

import (
    "encoding/json"
    "log"
    "order-service/src/order/domain/entities"
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

func (r *RabbitRepository) PublishOrder(order entities.Order) error {
    orderJSON, err := json.Marshal(order)
    if err != nil {
        log.Printf("Error marshaling order: %v", err)
        return err
    }

    message := string(orderJSON)
    err = r.rabbitMQ.Publish(message)
    if err != nil {
        log.Printf("Failed to publish order: %v", err)
        return err
    }
    log.Printf("Order published: %s", message)
    return nil
}

func (r *RabbitRepository) Close() {
    r.rabbitMQ.Close()
}