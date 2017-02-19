FROM golang:latest

ADD . /go/src/github.com/grameh/GoDevEnvironment/
RUN go get -v github.com/gorilla/mux \
              github.com/mattn/go-sqlite3

EXPOSE 8080
