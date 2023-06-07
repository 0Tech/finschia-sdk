# Dockerize

## Introduction

You can build the binary in a Docker instance, build & install a Docker image
of it, and finally launch the cluster using the image. It helps you maintain
clear environment.

## Usage

The procedure consists of two parts: build and install. You can invoke the
procedure step-by-step, or invoke only one subprocedure you want which
implicitly executes the dependencies for you.

At first, you need to configure the project by:

``` shell
build_dir=build    # the folder which you want to output the artifacts into
cmake -S . -B $build_dir
```

After that, you may change some variables by:

``` shell
variable_name=FIXTURE_NUM_REGIONS   # the name of variable
variable_value=7                    # set the value of variable to
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

Now you may trigger some target. The core targets would be:

* `build`: build the project and make the binaries (e.g. daemon and cosmovisor)
* `install`: install a Docker image using the built binaries

You can trigger a target by:

``` shell
target_name=install
cmake --build $build_dir --target $target_name
```

You MUST trigger `install` prior to running the tests.
