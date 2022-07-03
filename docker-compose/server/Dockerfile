# syntax=docker/dockerfile:1

FROM golang:1.18
WORKDIR /docker-compose/server
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY server.go .
EXPOSE 8001
RUN go build -o /docker-server-go
CMD [ "/docker-server-go" ]