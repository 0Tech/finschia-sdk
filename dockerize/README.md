# Dockerize

## Introduction

You can build the binary in a Docker instance, build & install a Docker image of it, and finally launch the cluster using the image. It helps you maintain clear environment.

## Usage

The procedure consists of three parts: build, install, and deploy. You can invoke the procedure step-by-step, or invoke only one subprocedure you want which implicitly executes the dependencies for you.

At first, you need to configure the project by:

``` shell
build_dir=build    # the folder which you want to output the artifacts into
cmake -S . -B $build_dir
```

After that, you may change some variables by:

``` shell
variable_name=SIMD_NUM_INSTANCES    # the name of variable
variable_value=7                    # set the value of variable to
cmake -S . -B $build_dir -D$variable_name=$variable_value
```

You can provide as many changes as you want, and the changes are accumulated. The changes are cached so it persists forever until you make a change or remove your build directory.

You can see the list of variables you can change:

``` shell
cmake -S . -B $build_dir -L
```

You may also want to read the descriptions:

``` shell
cmake -S . -B $build_dir -LH
```

Now you can trigger some target. The targets which you may want to look at are:

* `build_simd`: build the project and make the binary
* `install_simd`: build a Docker image using the built binary
* `deploy_simd`: deploy a cluster using the Docker image

You can trigger a target by:

``` shell
target_name=deploy_simd
cmake --build $build_dir --target $target_name
```

You can launch the cluster after you have triggered `deploy_simd` target. You can learn the details at [here](deploy/README.md).
