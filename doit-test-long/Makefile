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
	docker-compose up --build

# Clean up
clean:
	go clean
