FROM alpine:latest

RUN mkdir /app

WORKDIR /app

ADD api-gateway /app/api-gateway

CMD ["./api-gateway", "--registry_address", "consul:8500"]
