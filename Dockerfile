FROM golang:1.21.6-alpine

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

# Switch back to the root user to adjust permissions in the entrypoint script
USER root

# Copy the entrypoint script and give execute permission
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

# Switch back to the app user
USER app

# Set the entrypoint script
ENTRYPOINT ["/entrypoint.sh"]

CMD ["air", "-c", ".air.toml"]
