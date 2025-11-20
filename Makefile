# Run the server
run:
	go run server/main.go

# Fix dependencies
tidy:
	go mod tidy

# Build binary
build:
	go build -o bin/user-service cmd/server/main.go

# Clean binary
clean:
	rm -rf bin


