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
