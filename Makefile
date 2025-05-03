BINARY_NAME = vault-cli
build:
	go build -o $(BINARY_NAME)
install: build
	sudo mv $(BINARY_NAME) /usr/local/bin/
	sudo chmod +x /usr/local/bin/$(BINARY_NAME)
clean:
	rm -f $(BINARY_NAME)