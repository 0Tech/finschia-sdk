#!/bin/sh

set -e

[ -n "$VOLUME" ]
[ -n "$SOURCE" ]
[ -n "$TARGET" ]

image=debian:bullseye

docker run --rm -t \
	   --mount type=bind,source="$(realpath .)",target=/source,readonly \
	   --mount type=volume,source=$VOLUME,target=/target \
	   --user $(id -u):$(id -g) \
	   $image install /source/"$SOURCE" -D /target/"$TARGET"
