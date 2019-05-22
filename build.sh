#!/bin/bash
echo 'Starting build.sh script'
# build for linux/amd64 and move to bin folder
echo 'Building for linux/amd64...'
GOOS=linux GOARCH=amd64 mage build
mv ./rubrikbeat ./bin/linux/amd64/
# build for linux/amd64 and move to bin folder
echo 'Building for darwin/amd64...'
GOOS=darwin GOARCH=amd64 mage build
mv ./rubrikbeat ./bin/darwin/amd64/
echo 'Build complete'