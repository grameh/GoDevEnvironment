.PHONY: all
all: clean build
.PHONY: build
build:
	#go get
	go build -o exampleWeb.exe

.PHONY: clean
clean:
	go clean

.PHONY: run
run: build
	./exampleWeb.exe

.PHONY: docker-run
docker-run:
	docker run -p 49160:8080 -it godevenv /bin/bash -c "cd /go/src/github.com/grameh/GoDevEnvironment; make run"
