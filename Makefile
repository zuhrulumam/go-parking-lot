APP_NAME = parking-lot
IMAGE_NAME = $(APP_NAME):latest

# Run seeding: floor=5, row=20, col=20
seed:
	go run main.go seed 5 20 20

# Run app
start:
	go run main.go start

all: seed start

test:
	go test -v ./...

coverage:
	go test $$(go list ./... | grep -v /mocks) -coverprofile=coverage.out
	go tool cover -func=coverage.out

# up
up:
	docker-compose up --build

# Clean up
clean:
	go run main.go clean-db

run-load-test:
	k6 run test.js --summary-export=summary.json > output.log 2>&1

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