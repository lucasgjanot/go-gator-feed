build:
	go build -o app main.go
run:
	./app
test:
	go test ./...
clean:
	rm -f app
