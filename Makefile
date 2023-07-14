.PHONY: protos
protos:
	 protoc --go_out=. --go-grpc_out=. proto/dummy.proto


.PHONY: server
server:
	go run server/server.go

.PHONY: heavy
heavy:
	go run client/heavy.go

.PHONY: light
light:
	go run client/light.go

		