#!/bin/bash
# build for linux/amd64 and move to bin folder
GOOS=linux GOARCH=amd64 mage build
mv ./rubrikbeat ./bin/linux/amd64/
# build for linux/amd64 and move to bin folder
GOOS=darwin GOARCH=amd64 mage build
mv ./rubrikbeat ./bin/darwin/amd64/