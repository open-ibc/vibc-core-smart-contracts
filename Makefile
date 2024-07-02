.SILENT:

.PHONY: build-contracts bindings-gen-go

build-contracts:
	echo "Building contracts"; \
	rm -frd ./out; \
	forge install; \
	forge build contracts lib/openzeppelin-contracts/contracts/proxy/ERC1967/ERC1967Proxy.sol lib/openzeppelin-contracts/contracts/proxy/ERC1967/ERC1967Upgrade.sol -C contracts --lib-paths lib --skip test script $(wildcard contracts/implementation_templates/**/*) utils draft Rc $(wildcard contracts/mocks/**/*) --extra-output-files abi

# Libraries do not generate the correct ABI code (ChannelState enum causes errors)
# So the Ibc.sol abi generated code from abigen throws errors but is not needed.
# This is because the types exported are included in the ABIs of contracts and
# correctly interpreted (enums -> uint8 etc), the library methods are used inside
# of other contract methods and thus bindings for them do not need to be generated
# as they are not publicly exposed, but rather used within the contract itself.
#
# 	ABIGen issue ref: https://github.com/ethereum/solidity/issues/9278
bindings-gen-go: build-contracts
	echo "Generating Go vIBC bindings..."; \
	rm -rfd ./bindings/go/* ; \
	files=$$(/usr/bin/find $$(pwd)/out -type d | sed 's,^.*/out/,out/,' | awk -F'/' '{if ($$NF ~ /\.sol$$/) print $$0}' | awk '{split($$1,p,"[/]"); if (p[2] !~ /([Rr]c|[Vv])[0-9]+|Init|Std|AddressUpgradeable|[0-9]\$$/ && gsub(/\./,"",p[2])==1) print $$0}' | grep "out/[A-Z]" | xargs -I {} /usr/bin/find {} -maxdepth 1 -name '*.abi.json'); \
	for abi_file in $$files; do \
		abi_base=$$(basename $$(dirname $$abi_file)); \
		if [ "$$abi_base" = "Ibc.sol" ]; then \
			continue; \
		fi; \
		type=$$(basename $$abi_file .abi.json); \
		pkg=$$(basename $$abi_base .sol | tr "[:upper:]" "[:lower:]"); \
		mkdir -p ./bindings/go/$$pkg; \
		abigen --abi $$abi_file --pkg $$pkg --type $$type --out ./bindings/go/$$pkg/$$type.go; \
	done; \
	echo "Done."

