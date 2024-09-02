package main

import (
	"encoding/json"
	"fmt"
	"log"

	"golang-rabbitmq-consumer/services/rabbitmq"
)

type CoffeeOrder struct {
	CoffeeType string  `json:"coffee_type"`
	Price      float64 `json:"price"`
}

func init() {

	// CONNECT TO RABBITMQ
	rabbitmq.NewRabbitMQConnection()

}

func main() {

	// DEFER CLOSE CONNECTION TO RABBITMQ
	defer rabbitmq.RabbitMQClient.CloseConnection()

	// CONSUME RABBITMQ QUEUE --> "coffee_orders"
	msgs, err := rabbitmq.RabbitMQClient.ConsumeRabbitMQQueue("coffee_orders")

	if err != nil {
		log.Fatalf("Failed to consume RabbitMQ queue: %s", err)
		return
	}

	// CHANNEL TO RECEIVE COFFEE ORDERS
	forever := make(chan bool)

	go func() {
		for d := range msgs {
			var coffeeOrder CoffeeOrder
			err := json.Unmarshal(d.Body, &coffeeOrder)
			if err != nil {
				log.Printf("Error reading coffee order (please check the JSON format): %s", err)
				continue
			}

			// PRINT THE COFFEE ORDER
			fmt.Printf("Received a coffee order: Coffee Type = %s, Price = %f\n", coffeeOrder.CoffeeType, coffeeOrder.Price)
		}
	}()

	log.Printf("[*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
