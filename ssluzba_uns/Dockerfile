# syntax=docker/dockerfile:1

## Build
FROM golang AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY controllers/ controllers/
COPY models/ models/

RUN go build -o /ssluzba

## Deploy
FROM ubuntu

WORKDIR /

COPY --from=build /ssluzba /ssluzba

EXPOSE 8080

ENTRYPOINT ["/ssluzba"]