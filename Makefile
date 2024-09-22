build:
	go build -o todos .

test:
	go test ./...

fmt:
	go fmt ./...

clean:
	rm ./todos