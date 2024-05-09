build:
	scripts/build.sh

start: build
	go run src/server.go
