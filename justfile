# list all cmds
_default:
  @echo This script is for internal testing only.
  @echo Do not publish before audit
  @echo Examples:
  @echo just build test-polyibc
  @echo just --dry-run build-all
  @echo just --evaluate
  @echo just --help for more options
  @echo
  @just --list -f {{justfile()}}

# Run all tests
test: test-contracts

# Test smart contracts
test-contracts:
    forge test