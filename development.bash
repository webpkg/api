#!/bin/bash
set -e

make clean

make

bin/webgo-darwin-amd64 config
bin/webgo-darwin-amd64 serve