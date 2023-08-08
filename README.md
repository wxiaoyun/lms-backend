# Library Management System

## Setup

### Install Go

`https://golang.google.cn/doc/install`

### Updating Environment Variables

- Edit the variables in .env.development accordingly

### Setting up DB

```bash
make setupDB
```

### Running Tests

```bash
make test
```

### Running the Server

```bash
make run
```

### Documentation

- Run the server with `make run` and go to the following url:
- `http://localhost:<port_defined_in.env>/swagger/`
