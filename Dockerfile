# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./cmd/* ./
COPY ./internal ./internal
COPY ./resources ../resources
COPY ./static ../static
COPY ./protofiles ./protofiles
RUN go build -o /new-backend-challenge

ENV BLACK_FRIDAY_DAY=12-30
ENV GIN_MODE=release
ENV PORT=:8080
CMD [ "/new-backend-challenge" ]

EXPOSE 8080/tcp
