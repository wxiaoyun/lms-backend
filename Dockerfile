FROM golang:1.21.6-alpine

ARG RAILWAY_ENVIRONMENT

RUN addgroup app && adduser -S -G app app

USER app

WORKDIR /app

COPY go.mod go.sum ./

USER root
RUN chown -R app:app . && \
    chmod -R 766 /app

USER app

# Set umask for the app user
RUN echo "umask 002" >> ~/.profile

RUN go mod download
RUN go install github.com/cosmtrek/air@latest

COPY . .

CMD ["air", "-c", ".air.toml"]
