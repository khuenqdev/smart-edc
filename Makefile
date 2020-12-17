install:
	go get .

build: install
	go build -o ./bin/smartedc-connector .

build-windows: install
	env GOOS=windows GOARCH=amd64 go build -o ./bin/smartedc-connector.exe .

build-mac: install
	env GOOS=darwin GOARCH=amd64 go build -o ./bin/smartedc-connector .

serve: build
	sudo ./bin/smartedc-connector

clean:
	rm -f ./bin/smartedc-connector