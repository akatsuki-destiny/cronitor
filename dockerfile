FROM golang:1.20.5-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o cronitor ./cmd/main.go

RUN chmod +x /app/cronitor

FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/cronitor /app
COPY ./dev.env ./dev.env

CMD ["/app/cronitor"]