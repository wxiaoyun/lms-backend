FROM golang:1.21.6-alpine

ARG RAILWAY_ENVIRONMENT

RUN addgroup app && adduser -S -G app app

USER app

WORKDIR /app

COPY go.mod go.sum ./

USER root
RUN chown -R app:app 
# Create the file_storage directory
RUN mkdir -p /app/file_storage

# Change the ownership to 'app' user and 'app' group
RUN chown -R app:app /app/file_storage

# Set write permissions for user and group
RUN chmod -R u+w,g+w /app/file_storage

USER app


RUN go mod download
RUN go install github.com/cosmtrek/air@latest

COPY . .

CMD ["air", "-c", ".air.toml"]
