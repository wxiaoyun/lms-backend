# Library Management System

## Setup

### Install Go

`https://golang.google.cn/doc/install`

### Updating Environment Variables

- Make a copy of .env.example and rename to .env.development
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

### Running the Server

```bash
make run
```
