FROM golang:latest AS builder

RUN mkdir /app
WORKDIR /app

ADD . /app
RUN make

RUN mkdir /config
WORKDIR /config

CMD [ "/app/paper-tracker" ]