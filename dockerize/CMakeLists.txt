cmake_minimum_required(VERSION 3.16)

project(dockerize)

set(BUILDER_NAME cosmossdk/rbuilder:latest CACHE STRING "Image name of the builder")
add_custom_target(build_simd ALL
  DEPENDS build-simd.sh
  COMMAND ${CMAKE_CURRENT_SOURCE_DIR}/build-simd.sh ${CMAKE_CURRENT_SOURCE_DIR}/.. simd ${BUILDER_NAME}
  )
