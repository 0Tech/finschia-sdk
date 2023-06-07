cmake_minimum_required(VERSION 3.16)

project(finschia-sdk)
enable_testing()

set(GO_SOURCE_DIR ${CMAKE_SOURCE_DIR}/..)
file(RELATIVE_PATH PROJECT_DIR_NAME ${GO_SOURCE_DIR} ${CMAKE_SOURCE_DIR})

set(DAEMON_NAME simd)
set(DAEMON_HOME /home/${DAEMON_NAME}/.simapp)
set(GITHUB_URL https://github.com/Finschia/${CMAKE_PROJECT_NAME})
set(FIXTURE_SYNCDIR /home/${DAEMON_NAME}/sync)

add_subdirectory(daemon)
add_subdirectory(cosmovisor)
add_custom_target(build ALL
  DEPENDS build_daemon build_cosmovisor)

add_subdirectory(runtime)

add_subdirectory(fixture)
