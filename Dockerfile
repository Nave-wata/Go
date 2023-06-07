FROM golang:latest

RUN useradd -ms /bin/bash dev

WORKDIR /go/src/app

USER dev