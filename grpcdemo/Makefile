gen:
	go-doudou svc grpc; protoc --proto_path=. --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative transport/grpc/helloworld.proto

server-tls:
	cd helloworld; go run cmd/main.go; cd ..

client-tls:
	cd enumdemo; go run cmd/main.go; cd ..

cert:
	./cert.sh

.PHONY: cert server-tls client-tls
