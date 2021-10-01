cmake_minimum_required(VERSION 3.16)

project(dockerize)

set(GO_SOURCE_DIR ${CMAKE_SOURCE_DIR}/..)
file(RELATIVE_PATH PROJECT_DIR_NAME ${GO_SOURCE_DIR} ${CMAKE_SOURCE_DIR})

set(SIMD_WORKDIR /home/simd CACHE STRING "the working directory of simd")
mark_as_advanced(SIMD_WORKDIR)
set(SIMD_SYNCDIR ${SIMD_WORKDIR}/sync)

add_subdirectory(binary)
add_subdirectory(runtime)
add_subdirectory(deploy)
