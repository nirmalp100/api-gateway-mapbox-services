# syntax=docker/dockerfile:1.0

FROM golang:1.16-alpine

RUN mkdir gateway

WORKDIR /gateway

COPY . /gateway/


RUN go mod download

RUN go build -o app .

EXPOSE 3000

CMD [ "/gateway/app" ]


