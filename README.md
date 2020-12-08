# go-database-app
Example of a golang application with postgres database

## How to build
```bash
$ make clean build
```

## How to test
```bash
$ make clean test
```

## How to play

### running the app
```bash
$ make up
```

### make some requests

#### create order - 201 created
```bash
$ curl --location --request POST 'http://localhost:8080/orders' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": "$fafc88d1-9f8d-4487-8af7-6d22d61ddf93",
    "amount": 5700,
    "items":[
        {
            "name": "Cerveja Budweiser Longneck",
            "quantity": 12,
            "price": 300
        },
        {
            "name": "Coca-cola 2 litros",
            "quantity": 3,
            "price": 700
        }
    ],
    "customer": {
        "name": "Mary Doo",
        "document": "80221189076",
        "address": "Rua Uber, 100"
    }
}'
```

#### create order - 400 bad request
```bash
$ curl --location --request POST 'http://localhost:8080/orders' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": "$b352efa8-c50b-45fe-8c54-8a4725633732",
    "amount": 5700,
    "items":[
        {
            "name": "",
            "quantity": 12,
            "price": 300
        },
        {
            "name": "Coca-cola 2 litros",
            "quantity": 3,
            "price": 0
        }
    ],
    "customer": {
        "name": "Mary Doo",
        "document": "",
        "address": "Rua Uber, 100"
    }
}'
```

## How to kill
```bash
$ make down
```