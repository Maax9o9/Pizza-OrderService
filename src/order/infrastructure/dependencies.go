package infrastructure

import (
    "log"
    "order-service/src/order/application"
    "order-service/src/order/infrastructure/controller"
    "order-service/src/order/infrastructure/adapters"
    "order-service/src/order/application/repositorys"
)

func Init() *controller.CreateOrderController {
    rabbitMQ, err := adapters.NewRabbitMQ("amqp://max:123@44.213.165.25:5672/", "Pizzas")
    if err != nil {
        log.Fatalf("Failed to initialize RabbitMQ: %v", err)
    }

    rabbitRepository := repositorys.NewRabbitRepository(rabbitMQ)

    createOrderUseCase := application.NewCreateOrder(rabbitRepository)
    createOrderController := controller.NewCreateOrderController(createOrderUseCase)

    return createOrderController
}