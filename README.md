# Library Management System Backend

Welcome to the backend of our Library Management System (LMS). Designed for efficiency and scalability, this system forms the backbone of a simple digital library management solution. Leveraging cutting-edge technologies and robust architectures, our LMS backend efficiently handles complex library operations, ensuring reliability and high performance.

## Key Features

- **Robust Digital Library Management**: Manages digital assets with advanced backend functionalities.
- **Full CRUD Operations**: Supports Create, Read, Update, and Delete (CRUD) operations for books, loans, reservations, fines, and user accounts.
- **Advanced Admin Capabilities**: Provides administrative tools for efficient and effective library management.
- **Role-Based Access Control (RBAC)**: Implements fine-grained access control for secure and efficient management of library resources.
- **Session-Based Authentication**: Ensures secure user authentication and management.
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

### 3. Setup Backend

To set up the main Postgres database, run:

```bash
make setupDB
```

### 4. Additional Development Setup

Install necessary Go packages and initialize Git hooks:

```bash
go get -u github.com/swellaby/captain-githook
captain-githook init
```

### 5. Running the Server

Start the server with:

```bash
make run
```

---

Our Library Management System Backend is designed to meet the needs of simple libraries, offering a perfect blend of performance, security, and ease of maintenance. Whether for academic, public, or private libraries, it provides the essential infrastructure to manage library operations effectively and efficiently.
