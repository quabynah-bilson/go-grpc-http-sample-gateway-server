# stage 1: build stage
FROM golang:1.21-alpine3.18 AS builder

# optional authors information
LABEL authors="eganow"

# Install git and ca-certificates (needed to be able to call HTTPS)
RUN apk --update add ca-certificates git

# Move to working directory /app
WORKDIR /app

# Copy the code into the container
COPY . .

# Build the application's binary
RUN go build -o main cmd/main.go


# stage 2: run stage
FROM alpine:latest

# Move to working directory /app
WORKDIR /app

# Copy the code into the container from builder
COPY --from=builder /app/main .
#COPY configs/.env .

# expose ports
EXPOSE 9900 50051

# Command to run the application when starting the container
CMD ["/app/main"]