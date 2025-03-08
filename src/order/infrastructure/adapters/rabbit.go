package adapters

import (
    "log"
    amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
    connection *amqp.Connection
    channel    *amqp.Channel
    queue      amqp.Queue
}

func NewRabbitMQ(url, queueName string) (*RabbitMQ, error) {
    conn, err := amqp.Dial(url)
    if err != nil {
        return nil, err
    }

    ch, err := conn.Channel()
    if err != nil {
        return nil, err
    }

    q, err := ch.QueueDeclare(
        queueName, // name
        true,      // durable
        false,     // delete when unused
        false,     // exclusive
        false,     // no-wait
        nil,       // arguments
    )
    if err != nil {
        return nil, err
    }

    return &RabbitMQ{
        connection: conn,
        channel:    ch,
        queue:      q,
    }, nil
}

func (r *RabbitMQ) Publish(message string) error {
    err := r.channel.Publish(
        "",         // exchange
        r.queue.Name, // routing key
        false,      // mandatory
        false,      // immediate
        amqp.Publishing{
            ContentType: "text/plain",
            Body:        []byte(message),
        })
    if err != nil {
        return err
    }
    log.Printf(" [x] Sent %s", message)
    return nil
}

func (r *RabbitMQ) Close() {
    r.channel.Close()
    r.connection.Close()
}