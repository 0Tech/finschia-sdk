#!/bin/sh

set -e

rel_output=$1
[ -n "$rel_output" ]
output=$(realpath $rel_output)

source_dir=$2
[ -n "$source_dir" ]

dockerizer=$3
[ -n "$dockerizer" ]
cd $source_dir

tar -c $(ls -A1 | grep -Ev '^('$dockerizer'|\.git)$' | sort) | sha512sum | awk '{print $1}' >$output
