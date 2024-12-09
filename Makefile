.SILENT:

# Hardcoded for simplicity
CONTRACT_NAMES = channel \
				 CrossL2Prover \
				 Dispatcher \
				 DummyLightClient \
				 DummyProofVerifier \
				 Earth \
				 ERC1967Proxy \
				 FeeVault \
				 Ibc \
				 IbcDispatcher \
				 IbcUtils \
				 ICrossL2Prover \
				 IDispatcher \
				 IFeeVault \
				 ILightClient \
				 IProofVerifier \
				 IUniversalChannelHandler \
				 Mars \
				 Moon \
				 Venus \
				 OptimisticLightClient \
				 OptimisticProofVerifier \
				 SequencerSoloClient \
				 SequencerSignatureVerifier \
				 UniversalChannelHandler

# Create the pattern for each contract
CONTRACT_ABI_PATTERNS = $(addsuffix .sol/*.abi.json,$(addprefix ./out/,$(CONTRACT_NAMES)))
CONTRACT_JSON_PATTERNS := $(addsuffix .sol/*.json,$(addprefix ./out/,$(CONTRACT_NAMES)))

# Use wildcard to expand each pattern
CONTRACT_ABI_FILES = $(foreach pattern,$(CONTRACT_ABI_PATTERNS),$(wildcard $(pattern)))
CONTRACT_BOTH_FILES = $(foreach pattern,$(CONTRACT_JSON_PATTERNS),$(wildcard $(pattern)))
CONTRACT_JSON_FILES = $(filter-out $(CONTRACT_ABI_FILES),$(CONTRACT_BOTH_FILES))

.PHONY: build-contracts bindings-gen-go bindings-gen-ts

build-contracts:
	forge --version || exit 1; \
	echo "Building contracts"; \
	rm -frd ./out; \
	forge install; \
	forge build --skip test script -C contracts \
		--lib-paths lib \
		--extra-output-files abi --force

# Libraries do not generate the correct ABI code (ChannelState enum causes errors)
# So the Ibc.sol abi generated code from abigen throws errors but is not needed.
# This is because the types exported are included in the ABIs of contracts and
# correctly interpreted (enums -> uint8 etc), the library methods are used inside
# of other contract methods and thus bindings for them do not need to be generated
# as they are not publicly exposed, but rather used within the contract itself.
#
# 	ABIGen issue ref: https://github.com/ethereum/solidity/issues/9278
bindings-gen-go: build-contracts 
	go version || exit 1; \
	abigen --version || exit 1; \
	echo "Generating Go vIBC bindings..."; \
	rm -rfd ./bindings/go/* ; \
	for abi_file in $(CONTRACT_ABI_FILES); do \
		abi_base=$$(basename $$(dirname $$abi_file)); \
		if [ "$$abi_base" = "Ibc.sol" ]; then \
			continue; \
		fi; \
		type=$$(basename $$abi_file .abi.json); \
		pkg=$$(basename $$type .sol | tr "[:upper:]" "[:lower:]"); \
		mkdir -p ./bindings/go/$$pkg; \
		abigen --abi $$abi_file --pkg $$pkg --type $$type --out ./bindings/go/$$pkg/$$type.go || exit 1; \
	done; \
	echo Running sanity check on go bindings ; \
	go build ./... || exit 1; \
	echo "Successfully generated go bindings."

bindings-gen-ts: build-contracts
	echo "Generating TypeScript bindings..."; \
	rm -rfd ./src/evm/contracts/*; \
	typechain --target ethers-v6 --out-dir ./src/evm/contracts $(CONTRACT_JSON_FILES); \
	echo "Successfully generated ts bindings."
