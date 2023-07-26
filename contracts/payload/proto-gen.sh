#!/bin/bash

protoc --plugin "$(which protoc-gen-sol)" --sol_out license=NONE,generate=all:./ proto/packet.proto
