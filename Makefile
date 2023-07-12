npm-install:
	npm install

build: npm-install
	npx hardhat compile --force
	tar -c -z -f - artifacts/contracts | \
		base64 | \
		awk 'BEGIN {print "export const contractsTemplate = `"} {print} END {print "`"}' > \
		generated/index.ts
	npx tsc -p tsconfig.json

test: npm-install
	npx hardhat test

clean:
	npx hardhat clean
	rm -rf dist

.DEFAULT_GOAL := build

.PHONY: test
