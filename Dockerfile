FROM golang:alpine as builder
RUN apk --no-cache add build-base tzdata ca-certificates
# WORKDIR /home/ec2-user/backend
WORKDIR /app