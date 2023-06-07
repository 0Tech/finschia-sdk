#!/bin/sh

set -e

# every variable must not be empty
assert_variables() {
	for name in $@
	do
		if [ -z "$(eval echo '$'$name)" ]
		then
			echo no variable: $name
			false
		fi
	done
}

assert_variables PROJECT_NAME

docker-compose -p $PROJECT_NAME -f $PROJECT_NAME.yml down --volumes --remove-orphans
