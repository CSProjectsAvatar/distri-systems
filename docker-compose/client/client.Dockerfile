FROM golang:1.18
WORKDIR /docker-compose/client
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY client.go .
RUN go build -o /docker-client-go
CMD [ "/docker-client-go" ]