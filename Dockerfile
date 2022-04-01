FROM ubuntu:18.04

WORKDIR /app
ADD . /app

RUN ./server