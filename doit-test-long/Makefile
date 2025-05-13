APP_NAME = parking-lot
IMAGE_NAME = $(APP_NAME):latest

# Run seeding: floor=3, row=10, col=10
seed:
	go run main.go seed 3 10 10

# Run app
start:
	go run main.go start

all: seed start

# up
up:
	docker-compose up traefik
	docker-compose up --build

# Clean up
clean:
	go clean

# this about traefik
start-traefik:
	docker-compose up traefik

scale:
	@if [ -z "$(n)" ]; then \
		echo "Usage: make scale n=3"; \
		exit 1; \
	fi; \
	docker-compose up -d --scale app=$(n)

stop:
	docker-compose down