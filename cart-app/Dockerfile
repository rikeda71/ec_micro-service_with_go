FROM golang:1.15-buster as builder

WORKDIR /work

ENV GO111MODULE="on"
ENV GOPATH=""

COPY go.mod go.sum ./
RUN go mod download
