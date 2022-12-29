FROM golang:alpine as builder
WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /docker-datavi-api

# EXPOSE 8080

# CMD [ "/docker-gs-ping" ]