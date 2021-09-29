#!/bin/sh

set -e

binary=$1
[ -n "$binary" ]

output=$2
[ -n "$output" ]

cp --dereference "$binary" simd
exec 1>/dev/null 2>&1
docker build -t simd:latest .

touch $output
