FROM golang:1.18.6-alpine3.16 AS builder
WORKDIR /app

ADD go.mod .
RUN go mod download

ADD . .

RUN go build -o ./service

FROM alpine:3.16
WORKDIR /app
RUN apk update && apk add bash

ARG CACHEBUST=1
COPY --from=builder /app/service .
COPY --from=builder /app/wait-for-it.sh .
ARG CACHEBUST=1
COPY --from=builder /app/configs /app/configs

RUN chmod +x ./wait-for-it.sh
EXPOSE 3000
CMD ./wait-for-it.sh mysql-service:3306 -t 120 && ./service
