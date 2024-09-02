package main

import (
	"golang-rabbitmq-server/controllers"
	"golang-rabbitmq-server/services/rabbitmq"

	"github.com/gin-gonic/gin"
)

func init() {

	// RABBITMQ
	rabbitmq.NewRabbitMQConnection()
}

func main() {

	// DEFER CLOSE CONNECTION TO RABBITMQ
	defer rabbitmq.RabbitMQClient.CloseConnection()

	router := gin.Default()

	// CREATE ORDER ROUTE
	router.POST("api/v1/orders", controllers.CreateOrder)

	// RUN SERVER
	router.Run(":" + "8080")
}
