.PHONY = all clean
all: run

build: main.go
	@echo "Building binary..."
	go build -o restaurantAPI.o
	clear

run:
	make build
	docker-compose up -d --remove-orphans
	./restaurantAPI.o

clean:
	@echo "Cleaning up..."
	docker-compose down
	rm restaurantAPI.o
	go clean
	clear