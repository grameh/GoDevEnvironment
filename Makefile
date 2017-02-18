build:
	go build -o exampleWeb
clean:
	go clean
make run: build
	./exampleWeb

default:
	clean
	build
