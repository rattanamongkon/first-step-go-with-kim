FROM golang:latest

# Set the current working directory
WORKDIR /app

# Copy the source code
COPY . .

# Download the dependencies
RUN go mod vendor

# Build the Go binary
RUN go build -o main .

# Expose the port that the service listens on
EXPOSE 8080

# Run the Go binary when the container starts
CMD ["./main"]