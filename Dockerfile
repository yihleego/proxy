FROM golang:alpine

ENV GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /proxy

COPY main.go .
COPY go.mod .

RUN export GOPATH=/proxy
RUN go build

EXPOSE 8080

ENTRYPOINT ["./proxy"]