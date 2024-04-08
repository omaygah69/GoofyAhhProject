test: main
	./main

main: main.go
	go build main.go

clean:
	rm -rf main
