# Dockerize

## TL;DR

``` shell
mkdir build
cd build
cmake ..
ctest
```

## Introduction

You can build the binary in a Docker instance, build & install a Docker image
of it, and finally launch the cluster using the image. It helps you maintain
clear environment.

## Usage

### Configure

At first, you need to configure the project by:

``` shell
build_dir=build  # the folder which you want to output the artifacts into
cmake -S . -B $build_dir
```

After that, you may change some variables by:

``` shell
variable_name=FIXTURE_NUM_REGIONS  # the name of variable
variable_value=7                   # set the value of variable to
cmake -S . -B $build_dir -D$variable_name=$variable_value
```

You can provide as many changes as you want, and the changes would be
accumulated. The changes are cached so it persists forever until you override
the change or remove your build directory.

You can see the list of variables which you may want to set:

``` shell
cmake -S . -B $build_dir -L
```

You may also want to read the descriptions:

``` shell
cmake -S . -B $build_dir -LH
```

### Build

Now you may trigger some target. The core targets would be:

* `build`: build the project and make the binaries (e.g. daemon and cosmovisor)

You can trigger a target by:

``` shell
target_name=build
cmake --build $build_dir --target $target_name
```

### Test

You can trigger all the tests by:

``` shell
ctest --test-dir $build_dir
```

Or, a certain test by:

``` shell
test_name=upgrade_chain_auto
ctest --test-dir $build_dir -R $test_name
```

There are several variables relevant to the tests:

* `FIXTURE_NUM_REGIONS`: number of regions consisting chain cluster
                         (each region has one validator)
* `FIXTURE_DAEMON_VERSION`: version of binary in almost all scenarios
* `FIXTURE_DAEMON_VERSION_NEW`: version of the new binary in upgrade scenarios
