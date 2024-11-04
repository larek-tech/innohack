FROM golang:1.20-alpine

WORKDIR /app

# Install goose
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

COPY migrations/ /migrations

ENTRYPOINT ["goose"]
CMD ["up", "postgres://user:password@localhost:5432/dbname?sslmode=disable"]