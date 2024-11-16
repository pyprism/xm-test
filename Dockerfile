FROM golang:1.23.3-alpine

RUN apk add --no-cache inotify-tools tzdata

WORKDIR /app
COPY . /app

RUN go mod download
RUN go install github.com/air-verse/air@latest

ENV TZ=Asia/Dhaka

CMD air
