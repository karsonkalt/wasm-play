#!/bin/bash

if [ -d "wabt" ]; then
    read -p "Directory 'wabt' already exists. Do you want to remove it? [y/N] " answer
    if [[ "$answer" =~ ^[Yy]$ ]]; then
        rm -rf wabt
        echo "wabt directory removed"
        git clone --recursive https://github.com/WebAssembly/wabt
        cd wabt
        if git submodule update --init; then
            echo "Submodules updated."
        else
            echo "Failed to update submodules"
            exit 1
        fi
        cd ..
        rm -rf wabt/.git
        echo "removed wabt .git directory"
        echo "wabt installed successfully"
    fi
fi
