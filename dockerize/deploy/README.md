# Deploy #

## Introduction ##

You can deploy a `simd` cluster in few steps, and manipulate it with provided custom CLI.

## Usage ##

You need to build target `deploy_simd` if you have not done yet. You can get the hints at [here](../README.md).

Next, go to binary folder:
``` shell
cd $build/deploy
```

There would be a binary named `dsimd` (meaning dockerized `simd`), and it's all you have to deal with. You can always get help messages if you don't know what to do:
``` shell
./dsimd -h
```

First things first, you need to build up the cluster. Simply type:
``` shell
./dsimd up
```
It will build the cluster and start all the instances for you.

You can check the status:
``` shell
./dsimd ps
```

If you want to stop some instances, for example, instance 0 to 2:
``` shell
./dsimd stop -i 0,1,2
```

And you can also manipulate the clients inside the instances with `exec`:
``` shell
./dsimd exec -i 3 -- query block
```
You should provide the arguments of `simd` after `--`.

Finally, if you have done your tests, tidy it up:
``` shell
./dsimd down
```
It will stop all the instances and remove corresponding containers.
