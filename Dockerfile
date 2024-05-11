# base go image
FROM golang:1.22-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o webServer ./

RUN chmod +x /app/webServer

# build a tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/webServer /app

CMD [ "/app/webServer" ]