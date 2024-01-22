
build.linux:
	@echo "Building for Linux"
	@GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o mqttconsumertomongo.exe main.go

build.windows:
	@echo "Building for Windows"
	@GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o mqttconsumertomongo.exe main.go

build.darwin:
	@echo "Building for Darwin"
	@GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o mqttconsumertomongo.exe main.go

run.test:
	@echo "Running tests"
	@go test -v -cover ./...