# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY ./src/go.mod ./
COPY ./src/go.sum ./
RUN go mod download

COPY ./src/cmd/* ./
COPY ./src/internal ./internal
COPY ./src/resources ../resources
COPY ./src/static ../static
COPY ./src/protofiles ./protofiles
RUN go build -o /new-backend-challenge

ENV BLACK_FRIDAY_DAY=12-30
ENV GIN_MODE=release
ENV PORT=:8080
CMD [ "/new-backend-challenge" ]

EXPOSE 8080/tcp
