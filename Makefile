clean:
	rm -f ./gowc 
build:
	go build -o gowc main.go	
t:
	go test ./...
test: clean build t
