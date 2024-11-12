BINARY_NAME=zenithApp

build:
	@go mod vendor
	@echo "Building Zenith..."
	@go build -o tmp/${BINARY_NAME} .
	@echo "Zenith built!"

run: build
	@echo "Starting Zenith..."
	@./tmp/${BINARY_NAME} &
	@echo "Zenith started!"

clean:
	@echo "Cleaning..."
	@go clean
	@rm tmp/${BINARY_NAME}
	@echo "Cleaned!"

test:
	@echo "Testing..."
	@go test ./...
	@echo "Done!"

start: run

stop:
	@echo "Stopping Zenith..."
	@-pkill -SIGTERM -f "./tmp/${BINARY_NAME}"
	@echo "Stopped Zenith!"

restart: stop start