FROM golang:1.19

RUN go version

ENV GOPATH=/

COPY ["project/src", ""]