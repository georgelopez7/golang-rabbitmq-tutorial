<div align="center">

  <img src="assets/logo.png" alt="logo" width="200" height="auto" />
  <h1>Golang RabbitMQ Tutorial</h1>
  
  <p>
    A project t learn how to use RabbitMQ with Golang!
  </p>

<br />

<!-- Table of Contents -->

# :notebook_with_decorative_cover: Table of Contents

- [About the Project](#star2-about-the-project)
  - [System Architecture](#system-architecture)
- [Getting Started](#toolbox-getting-started)
  - [Prerequisites](#bangbang-prerequisites)
  - [Installation](#gear-installation)
  - [Run Locally](#running-run-locally)

<!-- About the Project -->

## :star2: About the Project

<!-- Screenshots -->

### :camera: System Architecture

<div align="center"> 
  <img src="https://placehold.co/600x400?text=Your+Screenshot+here" alt="screenshot" />
</div>

<!-- Getting Started -->

## :toolbox: Getting Started

<!-- Prerequisites -->

### :bangbang: Prerequisites

This project requires the following:

- Docker Desktop
- P

<!-- Installation -->

### :gear: Installation

Install my-project with npm

```bash
  yarn install my-project
  cd my-project
```

<!-- Run Locally -->

### :running: Run Locally

Clone the project

```bash
  git clone https://github.com/georgelopez7/golang-rabbitmq-tutorial.git
```

Go to the project directory

```bash
  cd golang-rabbitmq-tutorial
```

Run the docker-compose file

```bash
  docker compose up --build
```

<!-- Usage -->

## :eyes: Usage

You can access the server at http://localhost:8080

Send a POST request to http://localhost:8080/api/v1/orders with the following body:

```json
{
  "coffee_type": "Espresso",
  "price": 2.5
}
```

And see it appear printed out on the console of the consumer!
