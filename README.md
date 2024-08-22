## What do i nedd to run thin project ?
- Make
- Docker & Docker compose
- Go

## How to run this project 

1. Configure all the variables on a `.env` file in the project root, the `.env.example` should have reasanable default configs you can copy to run the project on development mode.
2. Make sure all necessary docker services are running <br>
    > `docker compose up -d`
3. Execute the migrations by running <br>
    > `make migrate`
4. Execte the project by running<br>
    > `go run cmd/ordersystem/main.go cmd/ordersystem/wire_gen.go`

## Regenrate
In clase you change any of the dependencies make sure to rerun wire
> `wire gen ./cmd/ordersystem`