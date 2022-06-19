## Getting Started

### Start Docker Compose
`docker-compose up -d`

### Running Tests
`cd {serviceName} && go test ./...`

#### DB and Kafka Error
Sometimes when we start the DB and Kafka, we get an error and stuck in infinite restart like `[error] failed to initialize database, got error dial tcp 192.168.176.5:3306: connect: connection refused`
that's because the mysql or kafka not starting yet. Since we always retry by `always: restart` inside docker compose, the container will restart until the mysql and kafka completely start.

## Services List

### Product Services

Documentation and the API is listed in GraphQL: http://127.0.0.1:8080/

### Order Services

Order service only create the record to db and pass the 

#### Order API

#### Request

``` json
curl --location --request POST 'http://127.0.0.1:3000/api/v1/orders' \
--header 'Content-Type: application/json' \
--data-raw '{
    "product_id": 1,
    "total_price": 10000,
    "total_quantity": 1
}'
```

#### Response

```text
Order created
```

### Stock Services

Stock services only listen to Kafka and expose the rpc server


### Stock Services Listener