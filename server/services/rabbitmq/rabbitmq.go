// rabbitmq.go
package rabbitmq

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
)

var RabbitMQClient *RabbitMQ

// STRUCT TO STORE RABBITMQ CONNECTION AND CHANNEL
type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

// STRUCT TO STORE COFFEE ORDER
type CoffeeOrder struct {
	CoffeeType string  `json:"coffee_type"`
	Price      float64 `json:"price"`
}

// FUNCTION TO CREATE A NEW RABBITMQ CONNECTION AND CHANNEL
func NewRabbitMQConnection() {

	// CONNECT TO RABBITMQ
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_CONNECTION_URL"))
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}

	// OPEN A RABBITMQ CHANNEL
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a RabbitMQ channel: %s", err)
	}

	// STORE RABBITMQ CONNECTION AND CHANNEL
	RabbitMQClient = &RabbitMQ{
		Conn:    conn,
		Channel: ch,
	}
}

// FUNCTION TO SEND A MESSAGE TO THE RABBITMQ QUEUE
func (r *RabbitMQ) SendCoffeeOrder(coffee_order CoffeeOrder, rabbitmq_queue string) error {

	// DECLARE THE QUEUE TO ENSURE IT EXISTS
	q, err := r.Channel.QueueDeclare(
		rabbitmq_queue, // queue name
		true,           // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)

	if err != nil {
		log.Fatalf("Failed to declare a queue: %s", err)
	}

	// CONVERT THE MESSAGE TO JSON FORMAT
	body, err := json.Marshal(coffee_order)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %v", err)
	}

	// PUBLISH THE MESSAGE TO THE QUEUE
	err = r.Channel.Publish(
		"",     // exchange
		q.Name, // routing key (queue name)
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})

	if err != nil {
		return fmt.Errorf("failed to publish message: %v", err)
	}

	log.Printf("Coffee order has been sent to RabbitMQ queue: %s", coffee_order)

	return nil
}

// FUNCTION TO CLOSE THE RABBITMQ CONNECTION AND CHANNEL
func (r *RabbitMQ) CloseConnection() {
	r.Channel.Close()
	r.Conn.Close()
}
