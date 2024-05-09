#!/bin/bash

# copy wasm_exec.js from go to the project
GOROOT=$(go env GOROOT)
DEST_DIR="./assets"
cp "$GOROOT/misc/wasm/wasm_exec.js" "$DEST_DIR"

if [ -e "$DEST_DIR/wasm_exec.js" ]; then
    echo "wasm_exec.js was copied successfully."
else
    echo "Failed to copy wasm_exec.js. Check the paths and permissions."
fi

# build the binary
echo "Building wasm binary..."
GOOS=js GOARCH=wasm go build -o assets/main.wasm src/main.go
echo "Build wasm successful"
