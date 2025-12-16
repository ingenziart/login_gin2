# Run the server
run:
	go run main.go

# Fix dependencies
tidy:
	go mod tidy

# Build binary
build:
	go build -o bin/user-service main.go

# Clean binary
clean:
	rm -rf bin


