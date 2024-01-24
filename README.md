# Library Management System Backend

Welcome to the backend of our Library Management System (LMS). Designed for efficiency and scalability, this system forms the backbone of a simple digital library management solution. Leveraging cutting-edge technologies and robust architectures, our LMS backend efficiently handles complex library operations, ensuring reliability and high performance.

## Key Features

- **Robust Digital Library Management**: Manages digital assets with advanced backend functionalities.
- **Full CRUD Operations**: Supports Create, Read, Update, and Delete (CRUD) operations for books, loans, reservations, fines, and user accounts.
- **Advanced Admin Capabilities**: Provides administrative tools for efficient and effective library management.
- **Role-Based Access Control (RBAC)**: Implements fine-grained access control for secure and efficient management of library resources.
- **Session-Based Authentication**: Ensures secure user authentication and management.
- **Dockerized Deployment**: Offers a convenient way to deploy the backend with Docker.
- **Scalable Data Storage with Postgres**: Utilizes PostgreSQL for robust and scalable data storage.
- **Redis for Performance**: Leverages Redis for high-speed caching and session storage, ensuring a responsive and efficient system.

## Core Technologies

The LMS backend is built using a range of powerful technologies:

- **Language**: [Go (Golang)](https://go.dev/doc/install) - Renowned for its efficiency and scalability in backend development.
- **ORM**: [GORM](https://gorm.io/index.html) - Offers a developer-friendly ORM library for Go.
- **Framework**: [Fiber](https://docs.gofiber.io/) - An Express-inspired, high-performance web framework for Go.
- **Databases**:
  - [**Postgres**](https://www.postgresql.org/): A versatile and reliable relational database system.
  - [**Redis**](https://redis.io/): A fast key-value store, excellent for caching and session storage.
- **Deployment**: [Docker](https://www.docker.com/) - A containerization platform for easy deployment.

## Accessing a Deployed Version

For those interested in exploring a deployed version of this setup, it is available for viewing and interaction:

- **Deployed on Railway**:
  - Visit the deployed version on Railway at [this link](https://railway.app/project/d296ea6f-2941-4176-8b32-ef7e210cf56a).

This deployed instance provides a convenient way to see the PostgreSQL and Redis services in action without the need for local installation.

## Setup Instructions

### 1. Install Go

Download and install Go from [here](https://go.dev/doc/install).

### 2. Update Environment Variables

- Copy `.env.example` to `.env.development`.
- Modify the variables in `.env.development` to suit your environment.

### 3. Running the Server

Start the server with:

```bash
make dockerup
```

### 4. Setting up Database

Run the docker-compose file to start the database:

```bash
# Keep the container running in a separate terminal
make dockerup
```

Run the database migrations:

```bash
# Execute a Shell Inside the Container
make dockershell
```

- Run the following commands inside the container:
- Create the database: `go run cmd/createdb/main.go`
- Migrate the database: `go run cmd/migratedb/main.go -dir=up`
- Rollback the database (specify the number of steps to roll back): `go run cmd/migratedb/main.go -dir=down -step= #$(step)`
- Seed the database: `go run cmd/seeddb/main.go`
- Drop all tables (if necessary): `go run cmd/flushdb/main.go`
- Drop the database (if necessary): `go run cmd/dropdb/main.go`
- Exit the container: `exit`

### 5. Making Changes to Frontend

The current is a single-page application (SPA) built using React. You can find the source code [here](https://github.com/ForAeons/lms-frontend-v2).

The backend is designed to serve the frontend as static files. If you wish to make changes to the frontend, you will need to set up the frontend locally and build the frontend files by running the following commands:

```bash
# cd into the frontend folder
bun run build
```

Once the frontend files have been built, copy the files from the `dist` folder to the `frontend` folder in the backend.

```bash
mv path_to_frontend/dist/* path_to_backend/frontend/
```

Great! You are now ready to serve the updated frontend from the backend.

### 6. Additional Development Setup

Install necessary Go packages and initialize Git hooks:

```bash
go get -u github.com/swellaby/captain-githook
captain-githook init
```

---

Our Library Management System Backend is designed to meet the needs of simple libraries, offering a perfect blend of performance, security, and ease of maintenance. Whether for academic, public, or private libraries, it provides the essential infrastructure to manage library operations effectively and efficiently.
