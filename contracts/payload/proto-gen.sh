#!/bin/bash

# protoc-gen-sol encoder doesn't work
# https://github.com/celestiaorg/protobuf3-solidity#parameters
protoc --plugin "$(which protoc-gen-sol)" --sol_out license=NONE,generate=all:./ proto/packet.proto
