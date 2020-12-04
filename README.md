# go-database-app
Example of a golang application with database

## How to build
```bash
$ make clean build
```

## How to play

### running the app
```bash
$ make up
```

### make some requests
```bash
$ curl --location --request POST 'http://localhost:8080/orders' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": "$7626f246-9b2e-4801-9155-a009a9dac741",
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
    ]
}'
```

## How to kill
```bash
$ make down
```