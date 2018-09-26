build:
	protoc --proto_path=proto --micro_out=proto --go_out=proto greeter.proto
	SET GOOS=linux && SET GOARCH=amd64
	go build
	docker-compose build

run:
	docker-compose up