FROM golang:1.18
RUN ["mkdir", "e-tournaments"]
WORKDIR /e-tournaments
COPY utils utils
COPY go.mod .
COPY go.sum .
COPY vendor vendor
RUN [ "mkdir", "tournament" ]
COPY tournament/domain tournament/domain
COPY tournament/infrastruct tournament/infrastruct
COPY tournament/interfaces tournament/interfaces
COPY tournament/usecases tournament/usecases
COPY tournament/worker.go tournament
RUN go build tournament/worker.go
ENTRYPOINT ["./worker"]
EXPOSE 9090
EXPOSE 8080-8082
