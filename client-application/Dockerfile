FROM golang:1.18-alpine

RUN mkdir -p /home/request-maker

COPY ./request-maker /home/request-maker

WORKDIR /home/request-maker

RUN go mod tidy

RUN go install