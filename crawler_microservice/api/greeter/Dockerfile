FROM alpine:latest

RUN mkdir /app

WORKDIR /app

ADD greeter /app/greeter

CMD ["./greeter", "--registry_address", "consul:8500"]
