# syntax=docker/dockerfile:1.0

FROM golang:1.16-alpine

RUN mkdir mapbox

WORKDIR /mapbox

COPY . /mapbox/

RUN go mod download



RUN go build  -o main .


EXPOSE 4040

CMD [ "/mapbox/main" ]








