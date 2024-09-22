build:
	go build -o todos .

test:
	go test ./...

clean:
	rm ./todos