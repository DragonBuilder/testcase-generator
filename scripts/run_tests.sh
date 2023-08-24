#!/bin/bash
export $(xargs < ./scripts/env/devtest.env)

## run all tests
go test -v ./...

## specific functions within a package
# go test -v testcase-generator/internal/openai -run TestStreamingChat 
# go test -v testcase-generator/internal -run TestStreamingGenerateTestcases 