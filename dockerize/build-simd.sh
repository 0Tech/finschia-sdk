#!/bin/sh

set -e

source="$1"
target="$2"
builder=$3

[ -n "$source" -a -n "$target" -a -n "$builder" ]

source=$(realpath "$source")
target=$(realpath "$target")

[ -e "$source" ]
rm -rf "$target"
mkdir "$target"

docker run --rm -t \
	   --env APP=simd \
	   --env VERSION=$(git describe --always | sed 's/^v//') \
	   --env COMMIT=$(git log -1 --format='%H') \
	   --env TARGET_OS=linux/amd64 \
	   --env LEDGER_ENABLED=false \
	   --env DEBUG=false \
	   --mount type=bind,source="$source",target=/sources,readonly=true \
	   --mount type=bind,source="$target",target=/target \
	   --entrypoint=/sources/dockerize/build-simd-native.sh \
	   $builder
sh -c "cd $target; ln -s simd* simd"
cat "$target/build_report"
