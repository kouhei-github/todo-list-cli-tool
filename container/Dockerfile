FROM golang:1.20 as build

WORKDIR /app

COPY ./src/go.mod /app
COPY ./src/go.sum /app
COPY ./src /app
RUN go mod download

