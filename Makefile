db:
	sudo docker compose -f internal/infrastructure/docker-compose.yaml up -d

run:
	go run cmd/main.go

build:
	go build cmd/main.go

