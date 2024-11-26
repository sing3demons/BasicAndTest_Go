FROM golang:alpine AS builder
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /docker-gin

EXPOSE 8080

CMD [ "/docker-gin" ]