# rabbitmq example with go workspace

go workspace is techique that allow to use multiple module in golang repository

This feature is imported from golang version 1.18

## start project

```shell
go work init
```

```shell
go work use common orders payment
```

## how to run each module

### 1 orders

```shell
go run ./orders
```

### 2 payment

```shell
go run ./payment
```

## enviroment

| enviroment | description |
|------------|-------------|
| RABBITMQ_URL| for connect rabbitmq |