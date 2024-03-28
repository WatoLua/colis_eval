FROM golang:1.20.2

ADD . /app/
WORKDIR /app

EXPOSE 7800
VOLUME /app/logs

CMD go run main.go
