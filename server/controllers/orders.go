package controllers

import (
	"golang-rabbitmq-server/services/rabbitmq"
	"net/http"

	"github.com/gin-gonic/gin"
)

// FUNCTION TO CREATE A NEW COFFEE ORDER
func CreateOrder(c *gin.Context) {

	// GET REQUEST BODY
	var coffee_order rabbitmq.CoffeeOrder
	err := c.BindJSON(&coffee_order)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
		return
	}

	// SEND COFFEE ORDER TO RABBITMQ
	rabbitmq_queue := "coffee_orders"
	err = rabbitmq.RabbitMQClient.SendCoffeeOrder(coffee_order, rabbitmq_queue)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to send coffee order to RabbitMQ queue"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Coffee order has been sent to RabbitMQ queue."})
}
