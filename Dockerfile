FROM golang:1.21.6-alpine

ARG RAILWAY_ENVIRONMENT

RUN addgroup app && adduser -S -G app app

USER app

WORKDIR /app

COPY go.mod go.sum ./

USER root
RUN chown -R app:app . && \
    mkdir -p /app/file_storage && \
    chown -R app:app /app/file_storage

USER app

RUN go mod download
RUN go install github.com/cosmtrek/air@latest

COPY . .

CMD ["air", "-c", ".air.toml"]
