# Build stage
FROM golang:1.22 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o .

# Final stage
FROM scratch
COPY --from=builder /app /app
ENTRYPOINT ["/app/registray"]
