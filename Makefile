default: run

clear:
	rm -rf build

run:
	mkdir -p build
	go build -o ./build/todo_server src/main.go
	./build/todo_server
