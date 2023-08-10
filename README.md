# Library Management System

## Setup

### Install Go

`https://golang.google.cn/doc/install`

### Updating Environment Variables

- Edit the variables in .env.development accordingly

### Setting up Backend

```bash
make setupDB
```

### Additional Setup for development

```bash
go get -u github.com/swellaby/captain-githook
captain-githook init
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
