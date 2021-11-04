.PHONY: build
build:
	$(MAKE) clean
	go build -o bin/git-extension main.go

.PHONY: clean
clean:
	rm -f ./bin/git-extension

.PHONY: install
install:
	$(MAKE) build
	cp bin/git-extension /usr/local/bin/git-extension
