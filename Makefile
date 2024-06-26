.SILENT:

CONTRACT_ABI_FILES = \
	./out/Dispatcher.sol/Dispatcher.json \
	./out/Mars.sol/Mars.json \
	./out/Earth.sol/Earth.json

TMP_ABI_DIR = .tmp_abi_vibc

.PHONY: build-contracts abigen clean extract_artifacts

build-contracts:
	echo "Building contracts"; \
	forge install; \
	forge build

extract_artifacts: build-contracts
	echo "Extracting ABI artifacts..."; \
	rm -rf $(TMP_ABI_DIR); \
	mkdir -p $(TMP_ABI_DIR); \
	for abi_file in $(CONTRACT_ABI_FILES); do \
		cat $$abi_file | jq .abi > $(TMP_ABI_DIR)/$$(basename $$abi_file); \
	done

abigen: extract_artifacts
	echo "Generating Go vIBC bindings..."; \
	for abi_file in $(wildcard $(TMP_ABI_DIR)/*.json); do \
		abi_base=$$(basename $$abi_file); \
		type=$$(basename $$abi_file .json | tr "[:upper:]" "[:lower:]"); \
		pkg=$$(basename $$abi_base .json | tr "[:upper:]" "[:lower:]"); \
		mkdir -p ./bindings/$$pkg; \
		abigen --abi $$abi_file --pkg $$pkg --type $$type --out ./bindings/$$pkg/$$type.go; \
	done

clean:
	echo "Cleaning up all artifacts..."; \
	rm -rf $(TMP_ABI_DIR); \
	forge clean
