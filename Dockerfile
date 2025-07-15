# Start from the official Go image
FROM golang:1.24.5

# Set the working directory inside the container
WORKDIR /app

#copy go.mod and go.sum
COPY go.mod go.sum /app/

#download dependencies
RUN go mod download

# Copy all files into the container
COPY . .

# Build the Go app
RUN go build -o main ./cmd

# Expose port 8080 (Gin default)
EXPOSE 4000

# Run the binary
CMD ["./main"]
