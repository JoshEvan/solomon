FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .
COPY ./etc /etc/files/solomon-api

ENV CONFIG_PATH=/etc/files/solomon-api/dev.yml

RUN echo $(ls -la /etc/files/solomon-api)
RUN go mod tidy

RUN go build -o consumer ./cmd/consumer

ENTRYPOINT ["/app/consumer"]
