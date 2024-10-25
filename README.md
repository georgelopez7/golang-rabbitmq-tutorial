# Golang RabbitMQ Tutorial

In this project, we will learn how to use RabbitMQ with Golang.

Check out the article associated with this project [here](https://medium.com/@george.benjamin.lopez/connecting-to-rabbitmq-using-golang-835b18d735a2).

## System Architecture

Below you can see a diagram of the system architecture for this project:

![golang-rabbitmq-tutorial-diagram](https://github.com/user-attachments/assets/a4b0a8ad-064c-4c77-a773-0835652283c3)

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
