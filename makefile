.DEFAULT_GOAL := all

.PHONY: clean
clean:
	go clean ./...
	rm -rf test/mocks/*

build:
	go build /mnt/c/Users/jkell/Google\ Drive/Go\ Files/VSCode_Projects/AVOXI\ Coding\ Challenge/src/cmd/main.go

run:
	go run src/cmd/main.go

