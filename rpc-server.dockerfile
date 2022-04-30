FROM golang:1.18
RUN ["mkdir", "rpc-server"]
WORKDIR /rpc-server
COPY utils utils
COPY go.mod .
COPY go.sum .
RUN [ "mkdir", "cps" ]
RUN [ "mkdir", "cps/2" ]
COPY cps/2 cps/2
WORKDIR /rpc-server/cps/2
EXPOSE 1234
CMD [ "go", "run", "server.go" ]