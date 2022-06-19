# README

## Diagram
![Alt text](diagram.png?raw=true "Diagram")

## Getting Started

### Start Docker Compose

`docker-compose up -d`

### Running Tests

`cd {serviceName} && go test ./...`

### DB and Kafka Error

Sometimes when we start the DB and Kafka, we get an error and are stuck in infinite restart like `[error] failed to initialize database, got error dial tcp 192.168.176.5:3306: connect: connection refused` that’s because the MySQL or Kafka not starting yet. Since we always retry by `always: restart` inside docker-compose, the container will restart until the MySQL and Kafka completely start.

### Distributed Tracing

We are using Open Tracing and Jaeger to trace that can be accesed through: [http://localhost:16686/](http://localhost:16686/)

### Distributed Transactions

We utilize event or service choreography to handle the transactions, it will rollback all the changes if any of the service fails. In our case when the stock is exceeded, then the order service will cancel the order.


## Product Services API (GraphQL)

Product service API can be run by `go run . serve` command. Documentation and the API are listed in GraphQL: http://127.0.0.1:8080

### Responsibilities

- Serve the frontend by given queries
- Check current stock by RPC call to Stock services.

## Order Services API (REST API)

Order service API can be run by `go run . serve` command.

Order service only creates the record to DB and trigger `order.created` event that Stock services will llisten.

### Responsibilities

- Creating order.
- Canceling order if the stock is exceeded by checking the stock to stock services.

Request

```json
curl --location --request POST 'http://127.0.0.1:3000/api/v1/orders' \--header 'Content-Type: application/json' \--data-raw '{    "product_id": 1,    "total_price": 10000,    "total_quantity": 1}'
```

Response

```
Order created
```

## Order Services Listener (Kafka)

Order service listener can be run by `go run . listen` command.

### Responsibilities

- Cancel order
    - This service will listen to all topics related to order, such as `stock.exceeded-amount` it will cancel the order if the selected order is matched and mark the order canceled by adding `is_canceled` to be `true`.

## Stock Services (gRPC)

Stock service listener can be run by `go run . serve` command.

### **Responsibilities**

- Serve RPC server to product services to the current stock of a given product.

## Stock Services Listener (Kafka)

Stock service listener can be run by `go run . listen` command.

Responsibilities

- Listening for all orders created events by Order services to support saga pattern
    - It will decrease a stock when a new order comes up
    - If the order stock is exceeded, then it will publish the `stock.exceeded-amount` events that Order services will listen.
- Stock service listener is listening for only one topic: `order.created` and will send all of the failure to `order.created.dlq` topic.