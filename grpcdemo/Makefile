server-tls:
	cd server; export GDD_PORT=8080 && go run cmd/main.go; cd ..

client-tls:
	cd client; go run cmd/main.go; cd ..

cert:
	./cert.sh

build-server-tls:
	docker build -f server/Dockerfile -t grpcdemo-server:latest .

build-client-tls:
	docker build -f client/Dockerfile -t grpcdemo-client:latest .

up:
	docker-compose -f docker-compose.yml up -d --remove-orphans

down:
	docker-compose -f docker-compose.yml down --remove-orphans

.PHONY: cert server-tls client-tls
