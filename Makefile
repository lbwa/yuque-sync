FUNCTION_DIR=serverless

build: clear
	cd $(FUNCTION_DIR); \
	GOOS=linux GOARCH=amd64 go build -o main main.go

clear:
	cd $(FUNCTION_DIR); \
	rm -rf main.zip; \
	go clean