test: main
	./main

main: main.go
	# go build controller.go
	go build main.go controller.go

clean:
	rm -rf main
	rm -rf controller

run: 
	go run main.go controller.go
