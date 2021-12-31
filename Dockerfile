# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./* ./

RUN go build -o /new-backend-challenge

EXPOSE 8080

ENV BLACK_FRIDAY_DAY="12-30"

CMD [ "/new-backend-challenge" ]