#!/bin/sh

set -e

curr=$1
[ -n "$curr" ]
prev=$curr~

output=$2
[ -n "$output" ]

if ! diff -q $prev $curr 1>/dev/null 2>&1
then
	cp $curr $prev
	touch $output
fi
