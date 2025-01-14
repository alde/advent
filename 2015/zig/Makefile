# Define a variable to find all .zig files in the folder and subfolders
ZIG_FILES := $(shell find ./src -name "*.zig")

# Define a target for testing all files
.PHONY: test

test:
	zig build test --summary all
