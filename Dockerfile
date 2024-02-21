FROM golang:1.21-alpine
LABEL authors="oliverkirk-sudo"

RUN apk add --update redis
# 暴露端口
EXPOSE 5000
WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o warpgpt

CMD redis-server & sleep 3 & ./warpgpt
