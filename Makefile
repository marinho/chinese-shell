# command to run build and run the zhell program
run:
	@echo "Building zhell..."
	@go build -o zhell main.go
	@echo "Running zhell..."
	@./zhell

default: run
