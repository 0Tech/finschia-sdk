#!/bin/sh

set -e

# path for the checksum
rel_output=$1
[ -n "$rel_output" ]
output=$(realpath $rel_output)

source_dir=$2
[ -n "$source_dir" ]

binary_dir=$3
[ -n "$binary_dir" ]
cd $source_dir

tar -c $(ls -A1 | grep -Ev '^('$binary_dir'|\.git)$' | sort) | sha512sum | awk '{print $1}' >$output
