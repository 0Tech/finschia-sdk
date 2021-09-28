cmake_minimum_required(VERSION 3.16)

project(dockerize)

set(SIMD_BUILDER_NAME cosmossdk/rbuilder:latest CACHE STRING "Image name of the builder")
mark_as_advanced(SIMD_BUILDER_NAME)

add_custom_target(build_simd ALL
  DEPENDS build-simd.sh
  COMMAND ${CMAKE_CURRENT_SOURCE_DIR}/build-simd.sh ${CMAKE_CURRENT_SOURCE_DIR}/.. simd ${SIMD_BUILDER_NAME}
  )

set(SIMD_WORKDIR /home/simd CACHE STRING "the working directory of simd")
mark_as_advanced(SIMD_WORKDIR)
set(SIMD_SYNCDIR ${SIMD_WORKDIR}/sync CACHE STRING "the sync directory of simd")
mark_as_advanced(SIMD_SYNCDIR)

add_subdirectory(runtime)
add_subdirectory(deploy)
