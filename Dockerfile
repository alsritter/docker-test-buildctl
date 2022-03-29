FROM ubuntu

WORKDIR /app
ADD . /app

RUN ./dctest