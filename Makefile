# command to run build and run the zhell program
run:
	@echo "Building zhell..."
	@mkdir -p ./bin
	@go build -o bin/zhell main.go
	@echo "Running zhell..."
	@./bin/zhell

default: run
