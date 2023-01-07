FROM golang:alpine as builder
# Set the current working directory
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
# Copy the source code
COPY . /app

RUN go mod tidy

# Download the dependencies
RUN go mod download

# Build the Go binary
RUN go build -o docker-datavi-api .

# Expose the port that the service listens on
# EXPOSE 8109

# Run the Go binary when the container starts
CMD ["./docker-datavi-api"]