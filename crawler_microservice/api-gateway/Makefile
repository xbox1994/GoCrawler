build:
	protoc --proto_path=proto --proto_path=%GOPATH% --micro_out=proto --go_out=proto api-gateway.proto
	SET GOOS=linux&&SET GOARCH=amd64&&go build
