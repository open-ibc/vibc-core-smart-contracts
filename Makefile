build:
	forge build

test: build
	forge test -vvv

clean:
	rm -rf dist

.DEFAULT_GOAL := build

.PHONY: test
