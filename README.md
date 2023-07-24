# Golang Version of the Rails Backend

Because ruby cringe

## Setup

### Install Go

`https://golang.google.cn/doc/install`

### Setting up DB

- Make sure you have created the corresponding database in your local postgresql with the same name as the one in the .env.development file.
- Run the following command to setup the database:

```bash
make setupDB
```

### Running the Server

```bash
make run
```

### Documentation

- Run the server with `make run` and go to the following url:
- `http://localhost:<port_defined_in.env>/swagger/`
