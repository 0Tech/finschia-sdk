#!/bin/sh

set -e

binary=$1
[ -n "$binary" ]

cp --dereference "$binary" simd
docker build -t simd:latest .
