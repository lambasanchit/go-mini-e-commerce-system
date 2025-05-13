# Use official Golang image to build the app
FROM golang:1.20 AS build

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download the dependencies (this will cache dependencies if go.mod & go.sum are not changed)
RUN go mod download

# Copy the rest of the source code into the container
COPY . .

# Build the Go app
RUN go build -o main .

# Start a new stage from scratch
FROM golang:1.20

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Pre-built binary file from the build stage
COPY --from=build /app/main .

# Expose port (for product-service)
EXPOSE 8002

# Run the binary
CMD ["./main"]
