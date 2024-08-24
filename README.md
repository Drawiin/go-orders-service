# Orders Service

## Prerequisites

Before running this project, ensure that the following dependencies are installed:

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [Go](https://golang.org/)

## Setup and Running the Project

1. **Environment Configuration:**

   - The project already includes a `.env` file with the default env modify then if they dont work for you

2. **Start Docker Services:**

   - Ensure all required Docker services are up and running by executing:
     ```bash
     docker compose up -d
     ```

3. **Start the Application:**
   - Run the project using the following Go command:
     ```bash
     go run cmd/ordersystem/main.go cmd/ordersystem/wire_gen.go
     ```

> **Tip:** The project includes php my admin access on [http://localhost:8080/](http://localhost:8080/)

## Testing the Project

### REST API Endpoints

- **Base URL:** `http://localhost:8000`
- **Available Endpoints:**
  - Create Order: `POST /orders/create`
  - List Orders: `GET /orders/list`
  - Get Order by ID: `GET /orders/{id}`

> **Tip:** You can find example HTTP requests in the `api/orders.http` file.

### GraphQL

- **Base URL:** `http://localhost:8081`
- Access the GraphQL playground by visiting [localhost:8081](http://localhost:8081).
- For further customization, the GraphQL server port can be changed via the `GRAPHQL_SERVER_PORT` environment variable.

### gRPC

To test the gRPC services, use [Evans](https://github.com/ktr0731/evans) with reflection support enabled. Run the following command:

```bash
evans -r repl
```

> **Note:** Ensure Evans is installed on your system.

## Code Generation

If you make any changes to the dependencies, be sure to regenerate the Wire dependency injection code with:

```bash
wire gen ./cmd/ordersystem
```

---

This version is more polished and organized, making it easier to understand and follow.
