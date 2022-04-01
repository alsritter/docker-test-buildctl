FROM ubuntu:18.04

WORKDIR /app
ADD dctest /app

RUN ./dctest