# Golang RabbitMQ Tutorial

In this project, we will learn how to use RabbitMQ with Golang.

Check out the article associated with this project [here](https://medium.com/@georgelopez7/golang-rabbitmq-tutorial-part-1-introduction-to-rabbitmq-and-golang-e9f3f6f4f5a9).

## System Architecture

Below you can see a diagram of the system architecture for this project:

## Getting Started

Follow the steps below to get started with the project:

### Prerequisites

The only real requirement for this project is Docker Desktop.

- [Docker Desktop](https://www.docker.com/products/docker-desktop/)

### Running the project locally

A step by step series of examples that tell you how to get a development
environment running

Clone the repository

    git clone https://github.com/georgelopez7/golang-rabbitmq-tutorial.git

Go to the project directory

    cd golang-rabbitmq-tutorial

Run the docker compose file

    docker compose up --build

### Testing the project

Now that we have the project running, we can test it by sending a POST request to `http://localhost:8080/api/v1/orders` with the following body:

```json
{
  "coffee_type": "Espresso",
  "price": 2.5
}
```

And see it appear printed out on the console of the consumer!
