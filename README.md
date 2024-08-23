## What do i do need to run thin project ?
- Docker & Docker compose
- Go

## How to run this project 

1. Configure all the variables on a `.env` file in the project root, the `.env.example` contains example values, the values can be easily used for testing 
2. Make sure all necessary docker services are running <br>
    > `docker compose up -d`
3. Execute the migrations by running <br>
    > `make migrate`
4. Execte the project by running<br>
    > `go run cmd/ordersystem/main.go cmd/ordersystem/wire_gen.go`

## How to test this project
- The `api/orders.http` contains example of how to make http requests
- Accessing `localhost:<GRAPHQL_SERVER_PORT>` for the graphql playground
- Evans can be used with the reflection option to test the grpc services by running `evans -r repl`
    > **Make sure you have evans intalled

## Regenerate
In clase you change any of the dependencies make sure to rerun wire
> `wire gen ./cmd/ordersystem`