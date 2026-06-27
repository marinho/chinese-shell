# command to run build and run the zhell program
run:
	@echo "Building zhell..."
	@mkdir -p ./bin
	@go build -o bin/zhell .
	@echo "Running zhell..."
	@./bin/zhell $(ARGS)

build:
	@echo "Building zhell..."
	@mkdir -p ./bin
	@go build -o bin/zhell .
	@echo "Build complete. Executable is located at ./bin/zhell"

default: run
