FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod main.go ./

RUN go mod verify \
  && go mod download \
  && go build -o /app/go/bin/nginx_check

FROM alpine:3.21

WORKDIR /app

COPY --from=builder /app/go/bin/nginx_check /app/nginx_check

CMD [ "/app/nginx_check" ]
