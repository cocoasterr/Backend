mux = app/gorillaMux/main.go
gin = app/gin/main.go

build:
	go build -o bin/app $(mux)

run: build
	go mod tidy
	./bin/app
