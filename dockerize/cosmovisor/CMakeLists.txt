set(COSMOVISOR_VERSION 1.4.0 CACHE STRING "cosmovisor version")

configure_file(Dockerfile.in Dockerfile)
configure_file(build.sh.in build.sh
  @ONLY)

add_custom_command(
  OUTPUT cosmovisor
  DEPENDS build.sh Dockerfile
  COMMAND ./build.sh Dockerfile
  COMMENT "Building cosmovisor"
  VERBATIM)
add_custom_target(build_cosmovisor
  DEPENDS cosmovisor)
