package services

import (
    "log"
    "order-service/src/order/infrastructure/adapters"
)

type EventService struct {
    rabbitMQ *adapters.RabbitMQ
}

func NewEventService(rabbitMQ *adapters.RabbitMQ) *EventService {
    return &EventService{
        rabbitMQ: rabbitMQ,
    }
}

func (s *EventService) PublishOrderCreatedEvent(orderID string) error {
    message := "Order Created: " + orderID
    err := s.rabbitMQ.Publish(message)
    if err != nil {
        log.Printf("Failed to publish order created event: %v", err)
        return err
    }
    log.Printf("Order created event published: %s", orderID)
    return nil
}

func (s *EventService) Close() {
    s.rabbitMQ.Close()
}