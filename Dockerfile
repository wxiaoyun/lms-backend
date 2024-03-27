FROM golang:1.22.1-alpine

ARG RAILWAY_ENVIRONMENT

RUN addgroup app && adduser -S -G app app

USER app

WORKDIR /app

COPY go.mod go.sum ./

USER root
RUN chown -R app:app .

USER app

RUN go mod download
RUN go install github.com/cosmtrek/air@latest

COPY . .

CMD ["air", "-c", ".air.toml"]
