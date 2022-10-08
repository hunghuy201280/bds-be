FROM golang:1.18.6-alpine3.16
WORKDIR /app
RUN apk update && apk add bash

ADD go.mod .
RUN go mod download

ADD . .

COPY wait-for-it.sh /wait-for-it.sh
RUN chmod +x ./wait-for-it.sh
RUN go build -o ./service
EXPOSE 3000
CMD ./wait-for-it.sh db:3306 -t 120 && ./service
