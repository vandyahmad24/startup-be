build-rest:
	@echo " >> building rest binaries"
	@go mod tidy
	@go build -v -o build main.go



run-rest: build-rest
	@echo "--- Running System ---"
	@./build
