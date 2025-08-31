# MicroserviceProject

A sample microservices architecture implemented in Go, featuring Account, Catalog, Order, and GraphQL services. Each service runs in its own container and communicates via gRPC and REST APIs. The project is designed for local development using Docker Compose.

## Prerequisites

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/)
- Go 1.24.5+ (for local builds outside Docker)


## Project Structure

```
.
├── account/         # Account microservice
├── catalog/         # Catalog microservice
├── order/           # Order microservice
├── graphql/         # GraphQL gateway service
├── docker-compose.yaml
├── go.mod
├── go.sum
└── README.md
```

## Running the Project

1. **Build and start all services:**
   ```sh
   docker compose up --build
   ```

2. **Access the GraphQL Playground:**
   Open the playground in your browser:
   ```sh
   "$BROWSER" http://localhost:8080
   ```

3. **Check service status and logs:**
   - List running containers:
     ```sh
     docker compose ps
     ```
   - View logs:
     ```sh
     docker compose logs
     ```

## Service Endpoints

| Service   | Default Port | Description                |
|-----------|-------------|----------------------------|
| Account   | 8081        | Account management API     |
| Catalog   | 8082        | Product catalog API        |
| Order     | 8083        | Order processing API       |
| GraphQL   | 8080        | Unified GraphQL endpoint   |
| Databases | 5432+       | PostgreSQL for services    |

## Development

- Each service has its own Dockerfile (`app.dockerfile`) for building Go binaries.
- Source code is organized by service.
- You can run individual services locally using Go:
  ```sh
  go run ./<service>/cmd/<service>
  ```
  *(Replace `<service>` with `account`, `catalog`, `order`, or `graphql`.)*

## Troubleshooting

- **Go version errors:**  
  Ensure your Dockerfiles use `golang:1.24.5-alpine` or newer to match the `go.mod` requirement.
- **Port conflicts:**  
  Change exposed ports in `docker-compose.yaml` if needed.
- **Dependency issues:**  
  Run `go mod tidy` in the project root to update dependencies.
