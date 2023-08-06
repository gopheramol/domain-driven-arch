# Domain driven architecture DDA

## Stack
### Golang
### ORM - SQLC
### DATABASE - PostgreSQL
### Deployment - Docker
### Clean code



# Project stucture

├── Dockerfile: The Dockerfile to build your Go application image.

├── Makefile: Makefile for automating tasks like building, running, and testing the application.

├── app.env: Environment file containing configuration variables for your application.

├── cmd: Directory containing the main Go application entry point.
│   └── main.go: The main entry point of your Go application.

├── config: Directory to store configuration-related code.
│   └── config.go: Configuration file or package for loading and managing application configuration.

├── docker-compose.yaml: Docker Compose file for managing multi-container deployments.

├── go.mod: Go module file containing module dependencies and version information.

├── go.sum: Go module file containing checksums of the dependencies.

├── internal: Directory containing internal application packages.
│   ├── ad: Directory for the ad-related packages.
│   │   ├── handler: HTTP handlers for ad-related endpoints.
│   │   ├── model: Structs representing the ad domain model.
│   │   ├── repository: Database repository implementation for ad entities.
│   │   └── service: Business logic or service layer for ad-related operations.
│   └── user: Directory for the user-related packages.
│       ├── handler: HTTP handlers for user-related endpoints.
│       ├── model: Structs representing the user domain model.
│       ├── repository: Database repository implementation for user entities.
│       └── service: Business logic or service layer for user-related operations.
├── pkg: Directory for shared application packages or libraries.
│   └── db: Directory for database-related code.
│       ├── orm: Directory for database migrations and SQL queries.
│       └── postgres: PostgreSQL database connection package.

├── start.sh: Script for starting your application.

├── wait-for.sh: Script for handling dependencies, such as waiting for the database to be ready.

└── web: Directory for web-related code.
    └── middleware: Contains middleware functions used in your web application, such as authentication and logging.